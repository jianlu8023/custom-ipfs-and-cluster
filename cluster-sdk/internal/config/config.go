package config

import (
	"time"
)

const (
	Timeout5  = 5 * time.Second
	Timeout10 = 10 * time.Second
	Timeout3  = 3 * time.Second
)

type Config struct {
	Host     string `mapstructure:"host" json:"host"`
	UserName string `mapstructure:"username,omitempty" json:"userName"`
	PassWord string `mapstructure:"password,omitempty" json:"passWord"`
	Port     int    `mapstructure:"port" json:"port"`
	Protocol string `mapstructure:"protocol" json:"protocol"`
}
