// Code generated by protoc-gen-cobra. DO NOT EDIT.

package pb

import (
	proto "github.com/golang/protobuf/proto"
	client "github.com/infraprime-tech/protoc-gen-cobra/client"
	flag "github.com/infraprime-tech/protoc-gen-cobra/flag"
	iocodec "github.com/infraprime-tech/protoc-gen-cobra/iocodec"
	cobra "github.com/spf13/cobra"
	pflag "github.com/spf13/pflag"
	grpc "google.golang.org/grpc"
	strconv "strconv"
	strings "strings"
)

func TypesClientCommand(options ...client.Option) *cobra.Command {
	cfg := client.NewConfig(options...)
	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("Types"),
		Short: "Types service client",
		Long:  "",
	}
	cfg.BindFlags(cmd.PersistentFlags())
	cmd.AddCommand(
		_TypesEchoCommand(cfg),
	)
	return cmd
}

func _TypesEchoCommand(cfg *client.Config) *cobra.Command {
	req := &Sound{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("Echo"),
		Short: "Echo RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Types"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Types", "Echo"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewTypesClient(cc)
				v := &Sound{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.Echo(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().Float64Var(&req.Double, cfg.FlagNamer("Double"), 0, "")
	cmd.PersistentFlags().Float32Var(&req.Float, cfg.FlagNamer("Float"), 0, "")
	cmd.PersistentFlags().Int32Var(&req.Int32, cfg.FlagNamer("Int32"), 0, "")
	cmd.PersistentFlags().Int64Var(&req.Int64, cfg.FlagNamer("Int64"), 0, "")
	cmd.PersistentFlags().Uint32Var(&req.Uint32, cfg.FlagNamer("Uint32"), 0, "")
	cmd.PersistentFlags().Uint64Var(&req.Uint64, cfg.FlagNamer("Uint64"), 0, "")
	cmd.PersistentFlags().Int32Var(&req.Sint32, cfg.FlagNamer("Sint32"), 0, "")
	cmd.PersistentFlags().Int64Var(&req.Sint64, cfg.FlagNamer("Sint64"), 0, "")
	cmd.PersistentFlags().Uint32Var(&req.Fixed32, cfg.FlagNamer("Fixed32"), 0, "")
	cmd.PersistentFlags().Uint64Var(&req.Fixed64, cfg.FlagNamer("Fixed64"), 0, "")
	cmd.PersistentFlags().Int32Var(&req.Sfixed32, cfg.FlagNamer("Sfixed32"), 0, "")
	cmd.PersistentFlags().Int64Var(&req.Sfixed64, cfg.FlagNamer("Sfixed64"), 0, "")
	cmd.PersistentFlags().BoolVar(&req.Bool, cfg.FlagNamer("Bool"), false, "")
	cmd.PersistentFlags().StringVar(&req.String_, cfg.FlagNamer("String_"), "", "")
	flag.BytesBase64Var(cmd.PersistentFlags(), &req.Bytes, cfg.FlagNamer("Bytes"), "")
	_Sound_NestedEnumVar(cmd.PersistentFlags(), &req.NestedEnum, cfg.FlagNamer("NestedEnum"), "")
	_GlobalEnumVar(cmd.PersistentFlags(), &req.GlobalEnum, cfg.FlagNamer("GlobalEnum"), "")
	cmd.PersistentFlags().Float64SliceVar(&req.ListDouble, cfg.FlagNamer("ListDouble"), nil, "")
	cmd.PersistentFlags().Float32SliceVar(&req.ListFloat, cfg.FlagNamer("ListFloat"), nil, "")
	cmd.PersistentFlags().Int32SliceVar(&req.ListInt32, cfg.FlagNamer("ListInt32"), nil, "")
	cmd.PersistentFlags().Int64SliceVar(&req.ListInt64, cfg.FlagNamer("ListInt64"), nil, "")
	flag.Uint32SliceVar(cmd.PersistentFlags(), &req.ListUint32, cfg.FlagNamer("ListUint32"), "")
	flag.Uint64SliceVar(cmd.PersistentFlags(), &req.ListUint64, cfg.FlagNamer("ListUint64"), "")
	cmd.PersistentFlags().Int32SliceVar(&req.ListSint32, cfg.FlagNamer("ListSint32"), nil, "")
	cmd.PersistentFlags().Int64SliceVar(&req.ListSint64, cfg.FlagNamer("ListSint64"), nil, "")
	flag.Uint32SliceVar(cmd.PersistentFlags(), &req.ListFixed32, cfg.FlagNamer("ListFixed32"), "")
	flag.Uint64SliceVar(cmd.PersistentFlags(), &req.ListFixed64, cfg.FlagNamer("ListFixed64"), "")
	cmd.PersistentFlags().Int32SliceVar(&req.ListSfixed32, cfg.FlagNamer("ListSfixed32"), nil, "")
	cmd.PersistentFlags().Int64SliceVar(&req.ListSfixed64, cfg.FlagNamer("ListSfixed64"), nil, "")
	cmd.PersistentFlags().BoolSliceVar(&req.ListBool, cfg.FlagNamer("ListBool"), nil, "")
	cmd.PersistentFlags().StringSliceVar(&req.ListString, cfg.FlagNamer("ListString"), nil, "")
	flag.BytesBase64SliceVar(cmd.PersistentFlags(), &req.ListBytes, cfg.FlagNamer("ListBytes"), "")
	_Sound_NestedEnumSliceVar(cmd.PersistentFlags(), &req.ListNestedEnum, cfg.FlagNamer("ListNestedEnum"), "")
	_GlobalEnumSliceVar(cmd.PersistentFlags(), &req.ListGlobalEnum, cfg.FlagNamer("ListGlobalEnum"), "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseFloat64, "string=float64", &req.MapStringDouble, cfg.FlagNamer("MapStringDouble"), "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseFloat32, "string=float32", &req.MapStringFloat, cfg.FlagNamer("MapStringFloat"), "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseInt32, "string=int32", &req.MapStringInt32, cfg.FlagNamer("MapStringInt32"), "")
	cmd.PersistentFlags().StringToInt64Var(&req.MapStringInt64, cfg.FlagNamer("MapStringInt64"), nil, "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseUint32, "string=uint32", &req.MapStringUint32, cfg.FlagNamer("MapStringUint32"), "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseUint64, "string=uint64", &req.MapStringUint64, cfg.FlagNamer("MapStringUint64"), "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseInt32, "string=int32", &req.MapStringSint32, cfg.FlagNamer("MapStringSint32"), "")
	cmd.PersistentFlags().StringToInt64Var(&req.MapStringSint64, cfg.FlagNamer("MapStringSint64"), nil, "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseUint32, "string=uint32", &req.MapStringFixed32, cfg.FlagNamer("MapStringFixed32"), "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseUint64, "string=uint64", &req.MapStringFixed64, cfg.FlagNamer("MapStringFixed64"), "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseInt32, "string=int32", &req.MapStringSfixed32, cfg.FlagNamer("MapStringSfixed32"), "")
	cmd.PersistentFlags().StringToInt64Var(&req.MapStringSfixed64, cfg.FlagNamer("MapStringSfixed64"), nil, "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseBool, "string=bool", &req.MapStringBool, cfg.FlagNamer("MapStringBool"), "")
	cmd.PersistentFlags().StringToStringVar(&req.MapStringString, cfg.FlagNamer("MapStringString"), nil, "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseBytesBase64, "string=bytesBase64", &req.MapStringBytes, cfg.FlagNamer("MapStringBytes"), "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, _Sound_NestedEnumParse, "string=Sound_NestedEnum", &req.MapStringNestedEnum, cfg.FlagNamer("MapStringNestedEnum"), "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, _GlobalEnumParse, "string=GlobalEnum", &req.MapStringGlobalEnum, cfg.FlagNamer("MapStringGlobalEnum"), "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseInt32, flag.ParseString, "int32=string", &req.MapInt32String, cfg.FlagNamer("MapInt32String"), "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseInt64, flag.ParseString, "int64=string", &req.MapInt64String, cfg.FlagNamer("MapInt64String"), "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseUint32, flag.ParseString, "uint32=string", &req.MapUint32String, cfg.FlagNamer("MapUint32String"), "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseUint64, flag.ParseString, "uint64=string", &req.MapUint64String, cfg.FlagNamer("MapUint64String"), "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseInt32, flag.ParseString, "int32=string", &req.MapSint32String, cfg.FlagNamer("MapSint32String"), "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseInt64, flag.ParseString, "int64=string", &req.MapSint64String, cfg.FlagNamer("MapSint64String"), "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseUint32, flag.ParseString, "uint32=string", &req.MapFixed32String, cfg.FlagNamer("MapFixed32String"), "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseUint64, flag.ParseString, "uint64=string", &req.MapFixed64String, cfg.FlagNamer("MapFixed64String"), "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseInt32, flag.ParseString, "int32=string", &req.MapSfixed32String, cfg.FlagNamer("MapSfixed32String"), "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseInt64, flag.ParseString, "int64=string", &req.MapSfixed64String, cfg.FlagNamer("MapSfixed64String"), "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseBool, flag.ParseString, "bool=string", &req.MapBoolString, cfg.FlagNamer("MapBoolString"), "")
	flag.TimestampVar(cmd.PersistentFlags(), &req.Timestamp, cfg.FlagNamer("Timestamp"), "")
	flag.DurationVar(cmd.PersistentFlags(), &req.Duration, cfg.FlagNamer("Duration"), "")
	flag.BoolWrapperVar(cmd.PersistentFlags(), &req.WrapperBool, cfg.FlagNamer("WrapperBool"), "")
	flag.BytesBase64WrapperVar(cmd.PersistentFlags(), &req.WrapperBytes, cfg.FlagNamer("WrapperBytes"), "")
	flag.DoubleWrapperVar(cmd.PersistentFlags(), &req.WrapperDouble, cfg.FlagNamer("WrapperDouble"), "")
	flag.FloatWrapperVar(cmd.PersistentFlags(), &req.WrapperFloat, cfg.FlagNamer("WrapperFloat"), "")
	flag.Int32WrapperVar(cmd.PersistentFlags(), &req.WrapperInt32, cfg.FlagNamer("WrapperInt32"), "")
	flag.Int64WrapperVar(cmd.PersistentFlags(), &req.WrapperInt64, cfg.FlagNamer("WrapperInt64"), "")
	flag.StringWrapperVar(cmd.PersistentFlags(), &req.WrapperString, cfg.FlagNamer("WrapperString"), "")
	flag.UInt32WrapperVar(cmd.PersistentFlags(), &req.WrapperUint32, cfg.FlagNamer("WrapperUint32"), "")
	flag.UInt64WrapperVar(cmd.PersistentFlags(), &req.WrapperUint64, cfg.FlagNamer("WrapperUint64"), "")
	flag.TimestampSliceVar(cmd.PersistentFlags(), &req.ListTimestamp, cfg.FlagNamer("ListTimestamp"), "")
	flag.DurationSliceVar(cmd.PersistentFlags(), &req.ListDuration, cfg.FlagNamer("ListDuration"), "")
	flag.BoolWrapperSliceVar(cmd.PersistentFlags(), &req.ListWrapperBool, cfg.FlagNamer("ListWrapperBool"), "")
	flag.BytesBase64WrapperSliceVar(cmd.PersistentFlags(), &req.ListWrapperBytes, cfg.FlagNamer("ListWrapperBytes"), "")
	flag.DoubleWrapperSliceVar(cmd.PersistentFlags(), &req.ListWrapperDouble, cfg.FlagNamer("ListWrapperDouble"), "")
	flag.FloatWrapperSliceVar(cmd.PersistentFlags(), &req.ListWrapperFloat, cfg.FlagNamer("ListWrapperFloat"), "")
	flag.Int32WrapperSliceVar(cmd.PersistentFlags(), &req.ListWrapperInt32, cfg.FlagNamer("ListWrapperInt32"), "")
	flag.Int64WrapperSliceVar(cmd.PersistentFlags(), &req.ListWrapperInt64, cfg.FlagNamer("ListWrapperInt64"), "")
	flag.StringWrapperSliceVar(cmd.PersistentFlags(), &req.ListWrapperString, cfg.FlagNamer("ListWrapperString"), "")
	flag.UInt32WrapperSliceVar(cmd.PersistentFlags(), &req.ListWrapperUint32, cfg.FlagNamer("ListWrapperUint32"), "")
	flag.UInt64WrapperSliceVar(cmd.PersistentFlags(), &req.ListWrapperUint64, cfg.FlagNamer("ListWrapperUint64"), "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseTimestamp, "string=timestamp", &req.MapStringTimestamp, cfg.FlagNamer("MapStringTimestamp"), "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseDuration, "string=duration", &req.MapStringDuration, cfg.FlagNamer("MapStringDuration"), "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseBoolWrapper, "string=bool", &req.MapStringWrapperBool, cfg.FlagNamer("MapStringWrapperBool"), "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseBytesBase64Wrapper, "string=bytesBase64", &req.MapStringWrapperBytes, cfg.FlagNamer("MapStringWrapperBytes"), "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseDoubleWrapper, "string=float64", &req.MapStringWrapperDouble, cfg.FlagNamer("MapStringWrapperDouble"), "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseFloatWrapper, "string=float32", &req.MapStringWrapperFloat, cfg.FlagNamer("MapStringWrapperFloat"), "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseInt32Wrapper, "string=int32", &req.MapStringWrapperInt32, cfg.FlagNamer("MapStringWrapperInt32"), "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseInt64Wrapper, "string=int64", &req.MapStringWrapperInt64, cfg.FlagNamer("MapStringWrapperInt64"), "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseStringWrapper, "string=string", &req.MapStringWrapperString, cfg.FlagNamer("MapStringWrapperString"), "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseUInt32Wrapper, "string=uint32", &req.MapStringWrapperUint32, cfg.FlagNamer("MapStringWrapperUint32"), "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseUInt64Wrapper, "string=uint64", &req.MapStringWrapperUint64, cfg.FlagNamer("MapStringWrapperUint64"), "")

	return cmd
}

type _GlobalEnumValue GlobalEnum

func _GlobalEnumVar(fs *pflag.FlagSet, p *GlobalEnum, name, usage string) {
	fs.Var((*_GlobalEnumValue)(p), name, usage)
}

func (v *_GlobalEnumValue) Set(val string) error {
	if e, err := parseGlobalEnum(val); err != nil {
		return err
	} else {
		*v = _GlobalEnumValue(e)
		return nil
	}
}

func (*_GlobalEnumValue) Type() string { return "GlobalEnum" }

func (v *_GlobalEnumValue) String() string { return (GlobalEnum)(*v).String() }

type _GlobalEnumSliceValue struct {
	value   *[]GlobalEnum
	changed bool
}

func _GlobalEnumSliceVar(fs *pflag.FlagSet, p *[]GlobalEnum, name, usage string) {
	fs.Var(&_GlobalEnumSliceValue{value: p}, name, usage)
}

func (s *_GlobalEnumSliceValue) Set(val string) error {
	ss := strings.Split(val, ",")
	out := make([]GlobalEnum, len(ss))
	for i, s := range ss {
		var err error
		if out[i], err = parseGlobalEnum(s); err != nil {
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

func (*_GlobalEnumSliceValue) Type() string { return "GlobalEnumSlice" }

func (*_GlobalEnumSliceValue) String() string { return "[]" }

func _GlobalEnumParse(val string) (interface{}, error) {
	return parseGlobalEnum(val)
}

func parseGlobalEnum(s string) (GlobalEnum, error) {
	if i, ok := GlobalEnum_value[s]; ok {
		return GlobalEnum(i), nil
	} else if i, err := strconv.ParseInt(s, 0, 32); err == nil {
		return GlobalEnum(i), nil
	} else {
		return 0, err
	}
}

type _Sound_NestedEnumValue Sound_NestedEnum

func _Sound_NestedEnumVar(fs *pflag.FlagSet, p *Sound_NestedEnum, name, usage string) {
	fs.Var((*_Sound_NestedEnumValue)(p), name, usage)
}

func (v *_Sound_NestedEnumValue) Set(val string) error {
	if e, err := parseSound_NestedEnum(val); err != nil {
		return err
	} else {
		*v = _Sound_NestedEnumValue(e)
		return nil
	}
}

func (*_Sound_NestedEnumValue) Type() string { return "Sound_NestedEnum" }

func (v *_Sound_NestedEnumValue) String() string { return (Sound_NestedEnum)(*v).String() }

type _Sound_NestedEnumSliceValue struct {
	value   *[]Sound_NestedEnum
	changed bool
}

func _Sound_NestedEnumSliceVar(fs *pflag.FlagSet, p *[]Sound_NestedEnum, name, usage string) {
	fs.Var(&_Sound_NestedEnumSliceValue{value: p}, name, usage)
}

func (s *_Sound_NestedEnumSliceValue) Set(val string) error {
	ss := strings.Split(val, ",")
	out := make([]Sound_NestedEnum, len(ss))
	for i, s := range ss {
		var err error
		if out[i], err = parseSound_NestedEnum(s); err != nil {
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

func (*_Sound_NestedEnumSliceValue) Type() string { return "Sound_NestedEnumSlice" }

func (*_Sound_NestedEnumSliceValue) String() string { return "[]" }

func _Sound_NestedEnumParse(val string) (interface{}, error) {
	return parseSound_NestedEnum(val)
}

func parseSound_NestedEnum(s string) (Sound_NestedEnum, error) {
	if i, ok := Sound_NestedEnum_value[s]; ok {
		return Sound_NestedEnum(i), nil
	} else if i, err := strconv.ParseInt(s, 0, 32); err == nil {
		return Sound_NestedEnum(i), nil
	} else {
		return 0, err
	}
}
