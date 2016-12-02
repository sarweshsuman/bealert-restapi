package handlers

import (
	"github.com/robfig/config"
)

type Config struct {
	cfg *config.Config
}

func New(configfilename string) *Config {
	v_cfg , _ := config.ReadDefault(configfilename)
	return &Config{cfg:v_cfg}
}