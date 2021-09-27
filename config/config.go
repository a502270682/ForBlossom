package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	Env        string     `mapstructure:"env"`
	AppName    string     `mapstructure:"app_name"`
	HTTPPort   string     `mapstructure:"http_port"`
	Mysql      Mysql      `mapstructure:"mysql"`
	Redis      Redis      `mapstructure:"redis"`
	WechatConf WechatConf `mapstructure:"wechat_conf"`
}

type WechatConf struct {
	AppID          string `mapstructure:"app_id"`           // appid
	AppSecret      string `mapstructure:"app_secret"`       // appsecret
	Token          string `mapstructure:"token"`            // token
	EncodingAESKey string `mapstructure:"encoding_aes_key"` // EncodingAESKey
}

type Redis struct {
	Host        string `mapstructure:"host" json:"host"`
	Password    string `mapstructure:"password" json:"password"`
	Database    int    `mapstructure:"database" json:"database"`
	MaxIdle     int    `mapstructure:"max_idle" json:"max_idle"`
	MaxActive   int    `mapstructure:"max_active" json:"max_active"`
	IdleTimeout int    `mapstructure:"idle_timeout" json:"idle_timeout"` // second
}

type Mysql struct {
	Master ConnectionConfig `mapstructure:"master"`
	Slave  ConnectionConfig `mapstructure:"slave"`
}

type ConnectionConfig struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Db       string `mapstructure:"db"`
	MaxIdle  int    `mapstructure:"max_idle"`
	MaxOpen  int    `mapstructure:"max_open"`
	Debug    bool   `mapstructure:"debug"`
}

var gConfig Config

func Load(path string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "failed to load conf")
	}

	if err := v.Unmarshal(&gConfig); err != nil {
		return nil, errors.Wrap(err, "fail to Unmarshal conf")
	}

	return &gConfig, nil
}
func GetConfig() *Config {
	return &gConfig
}
