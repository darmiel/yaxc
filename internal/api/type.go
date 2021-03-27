package api

import "github.com/spf13/viper"

type api struct {
	ServerURL string
}

func API() *api {
	server := viper.GetString("server")
	return &api{
		ServerURL: server,
	}
}
