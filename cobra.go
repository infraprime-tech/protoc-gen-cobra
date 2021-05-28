package main

import (
	"fmt"
	"sort"
	"strings"
	"text/template"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

func genFile(gen *protogen.Plugin, file *protogen.File) error {
	if len(file.Services) == 0 {
		return nil
	}

	g := gen.NewGeneratedFile(file.GeneratedFilenamePrefix+".cobra.pb.go", file.GoImportPath)
	g.P("// Code generated by protoc-gen-cobra. DO NOT EDIT.")
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()

	for _, srv := range file.Services {
		if err := genService(g, srv); err != nil {
			return err
		}
	}

	return nil
}

var (
	serviceTemplateCode = `
func {{.GoName}}ClientCommand(options ...client.Option) *cobra.Command {
	cfg := client.NewConfig(options...)
	cmd := &cobra.Command{
		Use: cfg.CommandNamer("{{.GoName}}"),
		Short: "{{.GoName}} service client",
		Long: {{.Comments.Leading | cleanComments | printf "%q"}},{{if .Desc.Options.GetDeprecated}}
		Deprecated: "deprecated",{{end}}
	}
	cfg.BindFlags(cmd.PersistentFlags())
	cmd.AddCommand({{range .Methods}}
		_{{$.GoName}}{{.GoName}}Command(cfg),{{end}}
	)
	return cmd
}
`
	serviceTemplate = template.Must(template.New("service").
		Funcs(template.FuncMap{"cleanComments": cleanComments}).
		Parse(serviceTemplateCode))
	serviceImports = []protogen.GoImportPath{
		"github.com/infraprime-tech/protoc-gen-cobra/client",
		"github.com/spf13/cobra",
	}
)

func genService(g *protogen.GeneratedFile, service *protogen.Service) error {
	for _, imp := range serviceImports {
		g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: imp})
	}
	if err := serviceTemplate.Execute(g, service); err != nil {
		return err
	}

	enums := make(map[string]*enum)

	for _, mth := range service.Methods {
		if err := genMethod(g, mth, enums); err != nil {
			return err
		}
	}

	if len(enums) > 0 {
		names := make([]string, len(enums))
		i := 0
		for name := range enums {
			names[i] = name
			i++
		}
		sort.Strings(names)
		for _, name := range names {
			if err := genEnum(g, enums[name]); err != nil {
				return err
			}
		}
	}

	return nil
}

var (
	methodTemplateCode = `
func _{{.Parent.GoName}}{{.GoName}}Command(cfg *client.Config) *cobra.Command {
	req := {{.InputInitializerCode}}

	cmd := &cobra.Command{
		Use: cfg.CommandNamer("{{.GoName}}"),
		Short: "{{.GoName}} RPC client",
		Long: {{.Comments.Leading | cleanComments | printf "%q"}},{{if .Desc.Options.GetDeprecated}}
		Deprecated: "deprecated",{{end}}
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "{{.Parent.GoName}}"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "{{.Parent.GoName}}", "{{.GoName}}"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := New{{.Parent.GoName}}Client(cc)
				v := &{{.Input.GoIdent.GoName}}{}
	{{if .Desc.IsStreamingClient}}
				stm, err := cli.{{.GoName}}(cmd.Context())
				if err != nil {
					return err
				}
				for {
					if err := in(v); err != nil {
						if err == io.EOF {
							_ = stm.CloseSend()
							break
						}
						return err
					}
					proto.Merge(v, req)
					if err = stm.Send(v); err != nil {
						return err
					}
				}
	{{else}}
				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)
		{{if .Desc.IsStreamingServer}}
				stm, err := cli.{{.GoName}}(cmd.Context(), v)
		{{else}}
				res, err := cli.{{.GoName}}(cmd.Context(), v)
		{{end}}
				if err != nil {
					return err
				}
	{{end}}
	{{if .Desc.IsStreamingServer}}
				for {
					res, err := stm.Recv()
					if err != nil {
						if err == io.EOF {
							break
						}
						return err
					}
					if err = out(res); err != nil {
						return err
					}
				}
				return nil
	{{else}}
		{{if .Desc.IsStreamingClient}}
				res, err := stm.CloseAndRecv()
				if err != nil {
					return err
				}
		{{end}}
				return out(res)
	{{end}}
			})
		},
	}

	{{.InputFieldFlagCode}}

	return cmd
}
`
	methodTemplate = template.Must(template.New("method").
		Funcs(template.FuncMap{"cleanComments": cleanComments}).
		Parse(methodTemplateCode))
	methodImports = []protogen.GoImportPath{
		"github.com/golang/protobuf/proto",
		"github.com/infraprime-tech/protoc-gen-cobra/client",
		"github.com/infraprime-tech/protoc-gen-cobra/flag",
		"github.com/infraprime-tech/protoc-gen-cobra/iocodec",
		"github.com/spf13/cobra",
		"google.golang.org/grpc",
	}
)

func genMethod(g *protogen.GeneratedFile, method *protogen.Method, enums map[string]*enum) error {
	for _, imp := range methodImports {
		g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: imp})
	}
	if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
		g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "io"})
	}

	initCode, flagCode := walkFields(g, method.Input, nil, enums, false, make(map[protogen.GoIdent]bool), 0, "")
	data := struct {
		*protogen.Method
		InputInitializerCode string
		InputFieldFlagCode   string
	}{method, initCode, flagCode}
	return methodTemplate.Execute(g, data)
}

var (
	basicTypes = map[protoreflect.Kind]struct{ Type, Parse, Value, Slice, Pointer, Default string }{
		protoreflect.BoolKind:   {"bool", "ParseBool", "BoolVar", "BoolSliceVar", "BoolPointerVar", "false"},
		protoreflect.Int32Kind:  {"int32", "ParseInt32", "Int32Var", "Int32SliceVar", "Int32PointerVar", "0"},
		protoreflect.Uint32Kind: {"uint32", "ParseUint32", "Uint32Var", "Uint32SliceVar", "Uint32PointerVar", "0"},
		protoreflect.Int64Kind:  {"int64", "ParseInt64", "Int64Var", "Int64SliceVar", "Int64PointerVar", "0"},
		protoreflect.Uint64Kind: {"uint64", "ParseUint64", "Uint64Var", "Uint64SliceVar", "Uint64PointerVar", "0"},
		protoreflect.FloatKind:  {"float32", "ParseFloat32", "Float32Var", "Float32SliceVar", "Float32PointerVar", "0"},
		protoreflect.DoubleKind: {"float64", "ParseFloat64", "Float64Var", "Float64SliceVar", "Float64PointerVar", "0"},
		protoreflect.StringKind: {"string", "ParseString", "StringVar", "StringSliceVar", "StringPointerVar", `""`},
		protoreflect.BytesKind:  {"bytesBase64", "ParseBytesBase64", "BytesBase64Var", "BytesBase64SliceVar", "", "nil"},
	}
	wrappersPkg  = protogen.GoImportPath("google.golang.org/protobuf/types/known/wrapperspb")
	timestampPkg = protogen.GoImportPath("google.golang.org/protobuf/types/known/timestamppb")
	durationPkg  = protogen.GoImportPath("google.golang.org/protobuf/types/known/durationpb")
	knownTypes   = map[protogen.GoIdent]struct{ Type, Parse, Value, Slice string }{
		timestampPkg.Ident("Timestamp"):  {"timestamp", "ParseTimestamp", "TimestampVar", "TimestampSliceVar"},
		durationPkg.Ident("Duration"):    {"duration", "ParseDuration", "DurationVar", "DurationSliceVar"},
		wrappersPkg.Ident("DoubleValue"): {"float64", "ParseDoubleWrapper", "DoubleWrapperVar", "DoubleWrapperSliceVar"},
		wrappersPkg.Ident("FloatValue"):  {"float32", "ParseFloatWrapper", "FloatWrapperVar", "FloatWrapperSliceVar"},
		wrappersPkg.Ident("Int64Value"):  {"int64", "ParseInt64Wrapper", "Int64WrapperVar", "Int64WrapperSliceVar"},
		wrappersPkg.Ident("UInt64Value"): {"uint64", "ParseUInt64Wrapper", "UInt64WrapperVar", "UInt64WrapperSliceVar"},
		wrappersPkg.Ident("Int32Value"):  {"int32", "ParseInt32Wrapper", "Int32WrapperVar", "Int32WrapperSliceVar"},
		wrappersPkg.Ident("UInt32Value"): {"uint32", "ParseUInt32Wrapper", "UInt32WrapperVar", "UInt32WrapperSliceVar"},
		wrappersPkg.Ident("BoolValue"):   {"bool", "ParseBoolWrapper", "BoolWrapperVar", "BoolWrapperSliceVar"},
		wrappersPkg.Ident("StringValue"): {"string", "ParseStringWrapper", "StringWrapperVar", "StringWrapperSliceVar"},
		wrappersPkg.Ident("BytesValue"):  {"bytesBase64", "ParseBytesBase64Wrapper", "BytesBase64WrapperVar", "BytesBase64WrapperSliceVar"},
	}
)

func walkFields(g *protogen.GeneratedFile, message *protogen.Message, path []string, enums map[string]*enum, deprecated bool, visited map[protogen.GoIdent]bool, level int, postSetCode string) (string, string) {
	initLines := make(map[int]string)
	flagLines := make(map[int]string, len(message.Fields))

	target := "req"
	if level > 0 {
		target = strings.Join(path[:level], "")
	}

	for _, fld := range message.Fields {
		path := append(path, fld.GoName)

		if f := flagFormat(g, fld, enums); f != "" {
			flagName := fmt.Sprintf("cfg.FlagNamer(%q)", strings.Join(path, " "))
			comment := cleanComments(fld.Comments.Leading)
			var flagLine string
			if fld.Oneof != nil {
				varName := strings.Join(path, "")
				goPath := fmt.Sprintf("&%s.%s", varName, fld.GoName)
				flagLine = fmt.Sprintf("%s := &%s{}\n", varName, fld.GoIdent.GoName)
				flagLine += fmt.Sprintf(f, goPath, flagName, comment)
				target := strings.Join(append([]string{target}, path[level:len(path)-1]...), ".")
				postSetCode := fmt.Sprintf("%s.%s = %s", target, fld.Oneof.GoName, varName)
				flagLine += fmt.Sprintf("\nflag.WithPostSetHook(cmd.PersistentFlags(), %s, func() { %s })", flagName, postSetCode)
			} else {
				goPath := fmt.Sprintf("&%s.%s", target, strings.Join(path[level:], "."))
				flagLine = fmt.Sprintf(f, goPath, flagName, comment)
			}
			if postSetCode != "" {
				flagLine += fmt.Sprintf("\nflag.WithPostSetHook(cmd.PersistentFlags(), %s, func() { %s })", flagName, postSetCode)
			}
			if deprecated || fld.Desc.Options().(*descriptorpb.FieldOptions).GetDeprecated() {
				flagLine += fmt.Sprintf("\n_ = cmd.PersistentFlags().MarkDeprecated(%s, \"deprecated\")", flagName)
			}
			flagLines[fld.Desc.Index()] = flagLine
		} else if normalizeKind(fld.Desc.Kind()) == protoreflect.MessageKind {
			if fld.Desc.IsList() {
				// message list not supported
			} else if fld.Desc.IsMap() {
				// limited map support
			} else if visited[message.GoIdent] = true; visited[fld.Message.GoIdent] {
				// cycle detected
			} else {
				m := make(map[protogen.GoIdent]bool, len(visited))
				for k, v := range visited {
					m[k] = v
				}

				level := level
				postSetCode := postSetCode
				if fld.Oneof != nil {
					if postSetCode != "" {
						postSetCode += ";"
					}
					target := strings.Join(append([]string{target}, path[level:len(path)-1]...), ".")
					postSetCode += fmt.Sprintf("%s.%s = &%s{%s: %s}", target, fld.Oneof.GoName, fld.GoIdent.GoName, fld.GoName, strings.Join(path, ""))
					level = len(path)
				}
				initCode, flagCode := walkFields(g, fld.Message, path, enums, deprecated, m, level, postSetCode)
				if initCode != "" && fld.Oneof == nil {
					initLines[fld.Desc.Index()] = fmt.Sprintf("%s: %s,", fld.GoName, initCode)
				}
				if flagCode != "" {
					if fld.Oneof != nil {
						flagName := fmt.Sprintf("cfg.FlagNamer(%q)", strings.Join(path, " "))
						flagLine := fmt.Sprintf("%s := %s\n", strings.Join(path, ""), initCode)
						flagLine += fmt.Sprintf("cmd.PersistentFlags().Bool(%s, false, \"\")\n", flagName)
						flagLine += fmt.Sprintf("flag.WithPostSetHook(cmd.PersistentFlags(), %s, func() { %s })\n", flagName, postSetCode)
						flagCode = flagLine + flagCode
					}
					flagLines[fld.Desc.Index()] = flagCode
				}
			}
		}
	}

	initCode := ""
	if len(initLines) > 0 {
		initCode = fmt.Sprintf("\n%s\n", strings.Join(sortedLines(initLines), "\n"))
	}
	return fmt.Sprintf("&%s{%s}", g.QualifiedGoIdent(message.GoIdent), initCode), strings.Join(sortedLines(flagLines), "\n")
}

func flagFormat(g *protogen.GeneratedFile, fld *protogen.Field, enums map[string]*enum) string {
	k := normalizeKind(fld.Desc.Kind())

	if bt, ok := basicTypes[k]; ok {
		if fld.Desc.IsList() {
			switch k {
			case protoreflect.Uint32Kind, protoreflect.Uint64Kind, protoreflect.BytesKind:
				return fmt.Sprintf("flag.%s(cmd.PersistentFlags(), %%s, %%s, %%q)", bt.Slice)
			default:
				return fmt.Sprintf("cmd.PersistentFlags().%s(%%s, %%s, nil, %%q)", bt.Slice)
			}
		} else if k == protoreflect.BytesKind {
			return fmt.Sprintf("flag.%s(cmd.PersistentFlags(), %%s, %%s, %%q)", bt.Value)
		} else if fld.Desc.HasPresence() && fld.Oneof == nil {
			return fmt.Sprintf("flag.%s(cmd.PersistentFlags(), %%s, %%s, %%q)", bt.Pointer)
		} else {
			return fmt.Sprintf("cmd.PersistentFlags().%s(%%s, %%s, %s, %%q)", bt.Value, bt.Default)
		}
	}

	switch k {
	case protoreflect.EnumKind:
		id := g.QualifiedGoIdent(fld.Enum.GoIdent)
		e, ok := enums[id]
		if !ok {
			e = &enum{Enum: fld.Enum}
			enums[id] = e
		}
		if fld.Desc.IsList() {
			e.List = true
			return fmt.Sprintf("_%sSliceVar(cmd.PersistentFlags(), %%s, %%s, %%q)", id)
		} else if fld.Desc.HasPresence() {
			e.Pointer = true
			return fmt.Sprintf("_%sPointerVar(cmd.PersistentFlags(), %%s, %%s, %%q)", id)
		} else {
			e.Value = true
			return fmt.Sprintf("_%sVar(cmd.PersistentFlags(), %%s, %%s, %%q)", id)
		}
	case protoreflect.MessageKind:
		if kt, ok := knownTypes[fld.Message.GoIdent]; ok {
			if fld.Desc.IsList() {
				return fmt.Sprintf("flag.%s(cmd.PersistentFlags(), %%s, %%s, %%q)", kt.Slice)
			} else {
				return fmt.Sprintf("flag.%s(cmd.PersistentFlags(), %%s, %%s, %%q)", kt.Value)
			}
		}
		if fld.Desc.IsMap() {
			kk := normalizeKind(fld.Desc.MapKey().Kind())
			vk := normalizeKind(fld.Desc.MapValue().Kind())
			if kk == protoreflect.StringKind {
				switch vk {
				case protoreflect.StringKind:
					return "cmd.PersistentFlags().StringToStringVar(%s, %s, nil, %q)"
				case protoreflect.Int64Kind:
					return "cmd.PersistentFlags().StringToInt64Var(%s, %s, nil, %q)"
				}
			}

			if bt, ok := basicTypes[kk]; ok {
				keyParser := "flag." + bt.Parse
				keyType := bt.Type
				valParser := ""
				valType := ""
				switch vk {
				case protoreflect.EnumKind:
					id := g.QualifiedGoIdent(fld.Message.Fields[1].Enum.GoIdent)
					e, ok := enums[id]
					if !ok {
						e = &enum{Enum: fld.Message.Fields[1].Enum}
						enums[id] = e
					}
					e.Map = true
					valParser = fmt.Sprintf("_%sParse", id)
					valType = id
				case protoreflect.MessageKind:
					id := fld.Message.Fields[1].Message.GoIdent
					if kt, ok := knownTypes[id]; ok {
						valParser = "flag." + kt.Parse
						valType = kt.Type
					}
				default:
					if bt, ok := basicTypes[vk]; ok {
						valParser = "flag." + bt.Parse
						valType = bt.Type
					}
				}
				if valParser != "" {
					typ := fmt.Sprintf("%s=%s", keyType, valType)
					return fmt.Sprintf("flag.ReflectMapVar(cmd.PersistentFlags(), %s, %s, %q, %%s, %%s, %%q)", keyParser, valParser, typ)
				}
			}
		}
	}

	return ""
}

func normalizeKind(kind protoreflect.Kind) protoreflect.Kind {
	switch kind {
	case protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		return protoreflect.Int32Kind
	case protoreflect.Fixed32Kind:
		return protoreflect.Uint32Kind
	case protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return protoreflect.Int64Kind
	case protoreflect.Fixed64Kind:
		return protoreflect.Uint64Kind
	case protoreflect.GroupKind:
		return protoreflect.MessageKind
	default:
		return kind
	}
}

func sortedLines(m map[int]string) []string {
	keys := make([]int, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	sort.Ints(keys)
	vals := make([]string, len(m))
	for i, k := range keys {
		vals[i] = m[k]
	}
	return vals
}

type enum struct {
	*protogen.Enum
	Value   bool
	Pointer bool
	List    bool
	Map     bool
}

var (
	enumTemplateCode = `
{{if .Value}}
type _{{.GoIdent.GoName}}Value {{.GoIdent.GoName}}

func _{{.GoIdent.GoName}}Var(fs *pflag.FlagSet, p *{{.GoIdent.GoName}}, name, usage string) {
	fs.Var((*_{{.GoIdent.GoName}}Value)(p), name, usage)
}

func (v *_{{.GoIdent.GoName}}Value) Set(val string) error {
	if e, err := parse{{.GoIdent.GoName}}(val); err != nil {
		return err
	} else {
		*v = _{{.GoIdent.GoName}}Value(e)
		return nil
	}
}

func (*_{{.GoIdent.GoName}}Value) Type() string { return "{{.GoIdent.GoName}}" }

func (v *_{{.GoIdent.GoName}}Value) String() string { return ({{.GoIdent.GoName}})(*v).String() }
{{end}}
{{if .Pointer }}
func _{{.GoIdent.GoName}}PointerVar(fs *pflag.FlagSet, p **{{.GoIdent.GoName}}, name, usage string) {
	v := fs.String(name, "", usage)
	hook := func() error {
		if e, err := parse{{.GoIdent.GoName}}(*v); err != nil {
			return err
		} else {
			*p = &e
			return nil
		}
	}
	flag.WithPostSetHookE(fs, name, hook)
}
{{end}}
{{if .List}}
type _{{.GoIdent.GoName}}SliceValue struct {
	value   *[]{{.GoIdent.GoName}}
	changed bool
}

func _{{.GoIdent.GoName}}SliceVar(fs *pflag.FlagSet, p *[]{{.GoIdent.GoName}}, name, usage string) {
	fs.Var(&_{{.GoIdent.GoName}}SliceValue{value: p}, name, usage)
}

func (s *_{{.GoIdent.GoName}}SliceValue) Set(val string) error {
	ss := strings.Split(val, ",")
	out := make([]{{.GoIdent.GoName}}, len(ss))
	for i, s := range ss {
		var err error
		if out[i], err = parse{{.GoIdent.GoName}}(s); err != nil {
			return err
		}
	}
	if !s.changed {
		*s.value = out
		s.changed = true
	} else {
		*s.value = append(*s.value, out...)
	}
	return nil
}

func (*_{{.GoIdent.GoName}}SliceValue) Type() string { return "{{.GoIdent.GoName}}Slice" }

func (*_{{.GoIdent.GoName}}SliceValue) String() string { return "[]" }
{{end}}
{{if .Map}}
func _{{.GoIdent.GoName}}Parse(val string) (interface{}, error) {
	return parse{{.GoIdent.GoName}}(val)
}
{{end}}

func parse{{.GoIdent.GoName}}(s string) ({{.GoIdent.GoName}}, error) {
	if i, ok := {{.GoIdent.GoName}}_value[s]; ok {
		return {{.GoIdent.GoName}}(i), nil
	} else if i, err := strconv.ParseInt(s, 0, 32); err == nil {
		return {{.GoIdent.GoName}}(i), nil
	} else {
		return 0, err
	}
}
`
	enumTemplate = template.Must(template.New("enum").Parse(enumTemplateCode))
	enumImports  = []protogen.GoImportPath{
		"strconv",
		"github.com/spf13/pflag",
	}
)

func genEnum(g *protogen.GeneratedFile, enum *enum) error {
	for _, imp := range enumImports {
		g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: imp})
	}
	// if enum.List is set, import strings
	if enum.List {
		g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "strings"})
	}
	return enumTemplate.Execute(g, enum)
}

func cleanComments(comments protogen.Comments) string {
	return strings.TrimSpace(string(comments))
}
