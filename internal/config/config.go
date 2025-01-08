package config

import (
	"github.com/BurntSushi/toml"
	"github.com/osamikoyo/goomu/pkg/loger"
)

type Config struct{
	Port int64 `toml:"port"`
	Upload_Dir string `toml:"upload_dir"`
}

func New() Config {
	var cfg Config

	if _, err := toml.DecodeFile("config.toml", &cfg);err != nil{
		loger.New().Error().Err(err)
	}

	return cfg
}