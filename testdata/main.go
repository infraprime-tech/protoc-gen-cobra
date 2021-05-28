package main

import (
	"context"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
	"log"
	"time"
)

func main() {
	log.Println("hi")
	cfg := &Config{
		ServerAddr:         "this is a nice address",
		RequestFile:        "",
		RequestFormat:      "",
		ResponseFormat:     "",
		Timeout:            0,
		UseEnvVars:         false,
		EnvVarPrefix:       "",
		TLS:                false,
		ServerName:         "",
		InsecureSkipVerify: false,
		CACertFile:         "",
		CertFile:           "",
		KeyFile:            "",
		headers:            nil,
	}


	cfg.headers = map[string]string{
		"Authorization": "hi",
	}

	headers := metadata.New(cfg.headers)


	cmd := &cobra.Command{
		Use:   "herllo",
		Short: "Deposit RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {

			return myFunc(cmd.Context(), cfg, func(lol0 string, lol1 string, lol2 string) error {
				md := metadata.New(cfg.headers)
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()
				ctx = metadata.NewOutgoingContext(ctx, md)

				log.Println(lol0)
				log.Println(lol1)
				log.Println(lol2)
				log.Println(cfg.ServerAddr)
				return nil
			})
		},
	}

	err := cmd.Execute()
	if err != nil {
		return 
	}
	
}

type Config struct {
	ServerAddr     string
	RequestFile    string
	RequestFormat  string
	ResponseFormat string
	Timeout        time.Duration
	UseEnvVars     bool
	EnvVarPrefix   string

	TLS                bool
	ServerName         string
	InsecureSkipVerify bool
	CACertFile         string
	CertFile           string
	KeyFile            string
	headers            map[string]string

}


func myFunc(ctx context.Context, cfg *Config, fn func(lol0 string, lol1 string, lol2 string) error) error {
	return fn("wassup", "hi", "how are you")
}
