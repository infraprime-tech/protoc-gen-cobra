// Code generated by protoc-gen-cobra. DO NOT EDIT.

package pb

import (
	context "context"
	tls "crypto/tls"
	x509 "crypto/x509"
	fmt "fmt"
	flag "github.com/NathanBaulch/protoc-gen-cobra/flag"
	iocodec "github.com/NathanBaulch/protoc-gen-cobra/iocodec"
	proto "github.com/golang/protobuf/proto"
	cobra "github.com/spf13/cobra"
	pflag "github.com/spf13/pflag"
	oauth2 "golang.org/x/oauth2"
	grpc "google.golang.org/grpc"
	credentials "google.golang.org/grpc/credentials"
	oauth "google.golang.org/grpc/credentials/oauth"
	ioutil "io/ioutil"
	net "net"
	os "os"
	filepath "path/filepath"
	strconv "strconv"
	strings "strings"
	time "time"
)

var Proto2ClientDefaultConfig = &_Proto2ClientConfig{
	ServerAddr:     "localhost:8080",
	RequestFormat:  "json",
	ResponseFormat: "json",
	Timeout:        10 * time.Second,
	AuthTokenType:  "Bearer",
}

type _Proto2ClientConfig struct {
	ServerAddr         string
	RequestFile        string
	RequestFormat      string
	ResponseFormat     string
	Timeout            time.Duration
	TLS                bool
	ServerName         string
	InsecureSkipVerify bool
	CACertFile         string
	CertFile           string
	KeyFile            string
	AuthToken          string
	AuthTokenType      string
	JWTKey             string
	JWTKeyFile         string
}

func (o *_Proto2ClientConfig) addFlags(fs *pflag.FlagSet) {
	fs.StringVarP(&o.ServerAddr, "server-addr", "s", o.ServerAddr, "server address in form of host:port")
	fs.StringVarP(&o.RequestFile, "request-file", "f", o.RequestFile, "client request file (must be json, yaml, or xml); use \"-\" for stdin + json")
	fs.StringVarP(&o.RequestFormat, "request-format", "i", o.RequestFormat, "request format (json, yaml, or xml)")
	fs.StringVarP(&o.ResponseFormat, "response-format", "o", o.ResponseFormat, "response format (json, prettyjson, xml, prettyxml, or yaml)")
	fs.DurationVar(&o.Timeout, "timeout", o.Timeout, "client connection timeout")
	fs.BoolVar(&o.TLS, "tls", o.TLS, "enable tls")
	fs.StringVar(&o.ServerName, "tls-server-name", o.ServerName, "tls server name override")
	fs.BoolVar(&o.InsecureSkipVerify, "tls-insecure-skip-verify", o.InsecureSkipVerify, "INSECURE: skip tls checks")
	fs.StringVar(&o.CACertFile, "tls-ca-cert-file", o.CACertFile, "ca certificate file")
	fs.StringVar(&o.CertFile, "tls-cert-file", o.CertFile, "client certificate file")
	fs.StringVar(&o.KeyFile, "tls-key-file", o.KeyFile, "client key file")
	fs.StringVar(&o.AuthToken, "auth-token", o.AuthToken, "authorization token")
	fs.StringVar(&o.AuthTokenType, "auth-token-type", o.AuthTokenType, "authorization token type")
	fs.StringVar(&o.JWTKey, "jwt-key", o.JWTKey, "jwt key")
	fs.StringVar(&o.JWTKeyFile, "jwt-key-file", o.JWTKeyFile, "jwt key file")
}

func Proto2ClientCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proto2",
		Short: "Proto2 service client",
		Long:  "",
	}
	Proto2ClientDefaultConfig.addFlags(cmd.PersistentFlags())
	cmd.AddCommand(
		_Proto2EchoCommand(),
	)
	return cmd
}

func _Proto2Dial(ctx context.Context) (*grpc.ClientConn, Proto2Client, error) {
	cfg := Proto2ClientDefaultConfig
	opts := []grpc.DialOption{grpc.WithBlock()}
	if cfg.TLS {
		tlsConfig := &tls.Config{InsecureSkipVerify: cfg.InsecureSkipVerify}
		if cfg.CACertFile != "" {
			caCert, err := ioutil.ReadFile(cfg.CACertFile)
			if err != nil {
				return nil, nil, fmt.Errorf("ca cert: %v", err)
			}
			certPool := x509.NewCertPool()
			certPool.AppendCertsFromPEM(caCert)
			tlsConfig.RootCAs = certPool
		}
		if cfg.CertFile != "" {
			if cfg.KeyFile == "" {
				return nil, nil, fmt.Errorf("missing key file")
			}
			pair, err := tls.LoadX509KeyPair(cfg.CertFile, cfg.KeyFile)
			if err != nil {
				return nil, nil, fmt.Errorf("cert/key: %v", err)
			}
			tlsConfig.Certificates = []tls.Certificate{pair}
		}
		if cfg.ServerName != "" {
			tlsConfig.ServerName = cfg.ServerName
		} else {
			addr, _, _ := net.SplitHostPort(cfg.ServerAddr)
			tlsConfig.ServerName = addr
		}
		cred := credentials.NewTLS(tlsConfig)
		opts = append(opts, grpc.WithTransportCredentials(cred))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	if cfg.AuthToken != "" {
		cred := oauth.NewOauthAccess(&oauth2.Token{
			AccessToken: cfg.AuthToken,
			TokenType:   cfg.AuthTokenType,
		})
		opts = append(opts, grpc.WithPerRPCCredentials(cred))
	}
	if cfg.JWTKey != "" {
		cred, err := oauth.NewJWTAccessFromKey([]byte(cfg.JWTKey))
		if err != nil {
			return nil, nil, fmt.Errorf("jwt key: %v", err)
		}
		opts = append(opts, grpc.WithPerRPCCredentials(cred))
	}
	if cfg.JWTKeyFile != "" {
		cred, err := oauth.NewJWTAccessFromFile(cfg.JWTKeyFile)
		if err != nil {
			return nil, nil, fmt.Errorf("jwt key file: %v", err)
		}
		opts = append(opts, grpc.WithPerRPCCredentials(cred))
	}
	if cfg.Timeout > 0 {
		var done context.CancelFunc
		ctx, done = context.WithTimeout(ctx, cfg.Timeout)
		defer done()
	}
	conn, err := grpc.DialContext(ctx, cfg.ServerAddr, opts...)
	if err != nil {
		return nil, nil, err
	}
	return conn, NewProto2Client(conn), nil
}

type _Proto2RoundTripFunc func(cli Proto2Client, in iocodec.Decoder, out iocodec.Encoder) error

func _Proto2RoundTrip(ctx context.Context, fn _Proto2RoundTripFunc) error {
	cfg := Proto2ClientDefaultConfig
	if cfg.ResponseFormat == "" {
		cfg.RequestFormat = "json"
	}
	var in iocodec.Decoder
	if stat, _ := os.Stdin.Stat(); (stat.Mode()&os.ModeCharDevice) == 0 || cfg.RequestFile == "-" {
		in = iocodec.MakeDecoder(cfg.RequestFormat, os.Stdin)
	} else if cfg.RequestFile != "" {
		f, err := os.Open(cfg.RequestFile)
		if err != nil {
			return fmt.Errorf("request file: %v", err)
		}
		defer f.Close()
		if ext := strings.TrimLeft(filepath.Ext(cfg.RequestFile), "."); ext != "" && ext != cfg.ResponseFormat {
			in = iocodec.MakeDecoder(ext, f)
		}
		if in == nil {
			in = iocodec.MakeDecoder(cfg.ResponseFormat, f)
		}
		if in == nil {
			return fmt.Errorf("invalid request format: %q", cfg.RequestFormat)
		}
	} else {
		in = iocodec.MakeDecoder("noop", os.Stdin)
	}
	if cfg.ResponseFormat == "" {
		cfg.ResponseFormat = "json"
	}
	out := iocodec.MakeEncoder(cfg.ResponseFormat, os.Stdout)
	if out == nil {
		return fmt.Errorf("invalid response format: %q", cfg.ResponseFormat)
	}
	conn, client, err := _Proto2Dial(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()
	return fn(client, in, out)
}

func _Proto2EchoCommand() *cobra.Command {
	req := &Sound2{}

	cmd := &cobra.Command{
		Use:   "echo",
		Short: "Echo RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return _Proto2RoundTrip(cmd.Context(), func(cli Proto2Client, in iocodec.Decoder, out iocodec.Encoder) error {
				v := &Sound2{}

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

	_Sound2_EnumPointerVar(cmd.PersistentFlags(), &req.Enum, "enum", "")
	_Sound2_EnumSliceVar(cmd.PersistentFlags(), &req.ListEnum, "listenum", "")
	cmd.PersistentFlags().BoolSliceVar(&req.ListBool, "listbool", nil, "")
	cmd.PersistentFlags().BytesBase64Var(&req.Bytes, "bytes", nil, "")
	cmd.PersistentFlags().Float32SliceVar(&req.ListFloat, "listfloat", nil, "")
	cmd.PersistentFlags().Float64SliceVar(&req.ListDouble, "listdouble", nil, "")
	cmd.PersistentFlags().Int32SliceVar(&req.ListInt32, "listint32", nil, "")
	cmd.PersistentFlags().Int32SliceVar(&req.ListSfixed32, "listsfixed32", nil, "")
	cmd.PersistentFlags().Int32SliceVar(&req.ListSint32, "listsint32", nil, "")
	cmd.PersistentFlags().Int64SliceVar(&req.ListInt64, "listint64", nil, "")
	cmd.PersistentFlags().Int64SliceVar(&req.ListSfixed64, "listsfixed64", nil, "")
	cmd.PersistentFlags().Int64SliceVar(&req.ListSint64, "listsint64", nil, "")
	cmd.PersistentFlags().StringSliceVar(&req.ListString, "liststring", nil, "")
	cmd.PersistentFlags().StringToInt64Var(&req.MapStringInt64, "mapstringint64", nil, "")
	cmd.PersistentFlags().StringToStringVar(&req.MapStringString, "mapstringstring", nil, "")
	flag.BoolPointerVar(cmd.PersistentFlags(), &req.Bool, "bool", "")
	flag.BoolWrapperSliceVar(cmd.PersistentFlags(), &req.ListWrapperBool, "listwrapperbool", "")
	flag.BoolWrapperVar(cmd.PersistentFlags(), &req.WrapperBool, "wrapperbool", "")
	flag.BytesBase64SliceVar(cmd.PersistentFlags(), &req.ListBytes, "listbytes", "")
	flag.BytesBase64WrapperSliceVar(cmd.PersistentFlags(), &req.ListWrapperBytes, "listwrapperbytes", "")
	flag.BytesBase64WrapperVar(cmd.PersistentFlags(), &req.WrapperBytes, "wrapperbytes", "")
	flag.DoubleWrapperSliceVar(cmd.PersistentFlags(), &req.ListWrapperDouble, "listwrapperdouble", "")
	flag.DoubleWrapperVar(cmd.PersistentFlags(), &req.WrapperDouble, "wrapperdouble", "")
	flag.DurationSliceVar(cmd.PersistentFlags(), &req.ListDuration, "listduration", "")
	flag.DurationVar(cmd.PersistentFlags(), &req.Duration, "duration", "")
	flag.Float32PointerVar(cmd.PersistentFlags(), &req.Float, "float", "")
	flag.Float64PointerVar(cmd.PersistentFlags(), &req.Double, "double", "")
	flag.FloatWrapperSliceVar(cmd.PersistentFlags(), &req.ListWrapperFloat, "listwrapperfloat", "")
	flag.FloatWrapperVar(cmd.PersistentFlags(), &req.WrapperFloat, "wrapperfloat", "")
	flag.Int32PointerVar(cmd.PersistentFlags(), &req.Int32, "int32", "")
	flag.Int32PointerVar(cmd.PersistentFlags(), &req.Sfixed32, "sfixed32", "")
	flag.Int32PointerVar(cmd.PersistentFlags(), &req.Sint32, "sint32", "")
	flag.Int32WrapperSliceVar(cmd.PersistentFlags(), &req.ListWrapperInt32, "listwrapperint32", "")
	flag.Int32WrapperVar(cmd.PersistentFlags(), &req.WrapperInt32, "wrapperint32", "")
	flag.Int64PointerVar(cmd.PersistentFlags(), &req.Int64, "int64", "")
	flag.Int64PointerVar(cmd.PersistentFlags(), &req.Sfixed64, "sfixed64", "")
	flag.Int64PointerVar(cmd.PersistentFlags(), &req.Sint64, "sint64", "")
	flag.Int64WrapperSliceVar(cmd.PersistentFlags(), &req.ListWrapperInt64, "listwrapperint64", "")
	flag.Int64WrapperVar(cmd.PersistentFlags(), &req.WrapperInt64, "wrapperint64", "")
	flag.StringPointerVar(cmd.PersistentFlags(), &req.String_, "string_", "")
	flag.StringWrapperSliceVar(cmd.PersistentFlags(), &req.ListWrapperString, "listwrapperstring", "")
	flag.StringWrapperVar(cmd.PersistentFlags(), &req.WrapperString, "wrapperstring", "")
	flag.TimestampSliceVar(cmd.PersistentFlags(), &req.ListTimestamp, "listtimestamp", "")
	flag.TimestampVar(cmd.PersistentFlags(), &req.Timestamp, "timestamp", "")
	flag.UInt32WrapperSliceVar(cmd.PersistentFlags(), &req.ListWrapperUint32, "listwrapperuint32", "")
	flag.UInt32WrapperVar(cmd.PersistentFlags(), &req.WrapperUint32, "wrapperuint32", "")
	flag.UInt64WrapperSliceVar(cmd.PersistentFlags(), &req.ListWrapperUint64, "listwrapperuint64", "")
	flag.UInt64WrapperVar(cmd.PersistentFlags(), &req.WrapperUint64, "wrapperuint64", "")
	flag.Uint32PointerVar(cmd.PersistentFlags(), &req.Fixed32, "fixed32", "")
	flag.Uint32PointerVar(cmd.PersistentFlags(), &req.Uint32, "uint32", "")
	flag.Uint32SliceVar(cmd.PersistentFlags(), &req.ListFixed32, "listfixed32", "")
	flag.Uint32SliceVar(cmd.PersistentFlags(), &req.ListUint32, "listuint32", "")
	flag.Uint64PointerVar(cmd.PersistentFlags(), &req.Fixed64, "fixed64", "")
	flag.Uint64PointerVar(cmd.PersistentFlags(), &req.Uint64, "uint64", "")
	flag.Uint64SliceVar(cmd.PersistentFlags(), &req.ListFixed64, "listfixed64", "")
	flag.Uint64SliceVar(cmd.PersistentFlags(), &req.ListUint64, "listuint64", "")

	return cmd
}

type _Sound2_EnumPointerValue struct {
	set func(*Sound2_Enum)
}

func _Sound2_EnumPointerVar(fs *pflag.FlagSet, p **Sound2_Enum, name, usage string) *_Sound2_EnumPointerValue {
	return &_Sound2_EnumPointerValue{func(e *Sound2_Enum) { *p = e }}
}

func (v *_Sound2_EnumPointerValue) Set(val string) error {
	if e, err := parseSound2_Enum(val); err != nil {
		return err
	} else {
		v.set(&e)
		return nil
	}
}

func (v *_Sound2_EnumPointerValue) Type() string { return "Sound2_EnumPointer" }

func (v *_Sound2_EnumPointerValue) String() string { return "<nil>" }

type _Sound2_EnumSliceValue struct {
	value   *[]Sound2_Enum
	changed bool
}

func _Sound2_EnumSliceVar(fs *pflag.FlagSet, p *[]Sound2_Enum, name, usage string) {
	fs.Var(&_Sound2_EnumSliceValue{value: p}, name, usage)
}

func (s *_Sound2_EnumSliceValue) Set(val string) error {
	ss := strings.Split(val, ",")
	out := make([]Sound2_Enum, len(ss))
	for i, s := range ss {
		var err error
		if out[i], err = parseSound2_Enum(s); err != nil {
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

func (s *_Sound2_EnumSliceValue) Type() string { return "Sound2_EnumSlice" }

func (s *_Sound2_EnumSliceValue) String() string { return "[]" }

func parseSound2_Enum(s string) (Sound2_Enum, error) {
	if i, ok := Sound2_Enum_value[s]; ok {
		return Sound2_Enum(i), nil
	} else if i, err := strconv.ParseInt(s, 0, 32); err == nil {
		return Sound2_Enum(i), nil
	} else {
		return 0, err
	}
}
