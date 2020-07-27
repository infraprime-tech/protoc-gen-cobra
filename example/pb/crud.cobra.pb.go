// Code generated by protoc-gen-cobra. DO NOT EDIT.

package pb

import (
	context "context"
	tls "crypto/tls"
	x509 "crypto/x509"
	fmt "fmt"
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
	strings "strings"
	time "time"
)

var CRUDClientDefaultConfig = &_CRUDClientConfig{
	ServerAddr:     "localhost:8080",
	RequestFormat:  "json",
	ResponseFormat: "json",
	Timeout:        10 * time.Second,
	AuthTokenType:  "Bearer",
}

type _CRUDClientConfig struct {
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

func (o *_CRUDClientConfig) addFlags(fs *pflag.FlagSet) {
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

func CRUDClientCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "crud",
		Short: "CRUD service client",
		Long:  "",
	}
	CRUDClientDefaultConfig.addFlags(cmd.PersistentFlags())
	cmd.AddCommand(
		_CRUDCreateCommand(),
		_CRUDGetCommand(),
		_CRUDUpdateCommand(),
		_CRUDDeleteCommand(),
	)
	return cmd
}

func _CRUDDial(ctx context.Context) (*grpc.ClientConn, CRUDClient, error) {
	cfg := CRUDClientDefaultConfig
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
	return conn, NewCRUDClient(conn), nil
}

type _CRUDRoundTripFunc func(cli CRUDClient, in iocodec.Decoder, out iocodec.Encoder) error

func _CRUDRoundTrip(ctx context.Context, fn _CRUDRoundTripFunc) error {
	cfg := CRUDClientDefaultConfig
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
	conn, client, err := _CRUDDial(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()
	return fn(client, in, out)
}

func _CRUDCreateCommand() *cobra.Command {
	req := &CreateCRUD{}

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return _CRUDRoundTrip(cmd.Context(), func(cli CRUDClient, in iocodec.Decoder, out iocodec.Encoder) error {
				v := &CreateCRUD{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.Create(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.Name, "name", "", "")
	cmd.PersistentFlags().StringVar(&req.Value, "value", "", "")

	return cmd
}

func _CRUDGetCommand() *cobra.Command {
	req := &GetCRUD{}

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return _CRUDRoundTrip(cmd.Context(), func(cli CRUDClient, in iocodec.Decoder, out iocodec.Encoder) error {
				v := &GetCRUD{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.Get(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.Name, "name", "", "")

	return cmd
}

func _CRUDUpdateCommand() *cobra.Command {
	req := &CRUDObject{}

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return _CRUDRoundTrip(cmd.Context(), func(cli CRUDClient, in iocodec.Decoder, out iocodec.Encoder) error {
				v := &CRUDObject{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.Update(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.Name, "name", "", "")
	cmd.PersistentFlags().StringVar(&req.Value, "value", "", "")

	return cmd
}

func _CRUDDeleteCommand() *cobra.Command {
	req := &CRUDObject{}

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return _CRUDRoundTrip(cmd.Context(), func(cli CRUDClient, in iocodec.Decoder, out iocodec.Encoder) error {
				v := &CRUDObject{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.Delete(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.Name, "name", "", "")
	cmd.PersistentFlags().StringVar(&req.Value, "value", "", "")

	return cmd
}
