package api

import "github.com/spf13/viper"

type api struct {
	ServerURL string
}

func API() *api {
	server := viper.GetString("server")
	if server == "" {
		server = "http://127.0.0.1:1332"
	}

	return &api{
		ServerURL: server,
	}
}
