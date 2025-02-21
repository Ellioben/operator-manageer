package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	AuthToken string `json:",optional"`
	Redis     struct {
		Host     string
		Password string
		DB       int
	}
	MySQL struct {
		DSN string
	}
	ClientSets map[string]string
}
