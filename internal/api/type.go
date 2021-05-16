// +build client

package api

import "github.com/spf13/viper"

type Api struct {
	ServerURL string
}

func API() *Api {
	server := viper.GetString("server")
	return &Api{
		ServerURL: server,
	}
}
