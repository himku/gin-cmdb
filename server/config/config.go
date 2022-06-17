package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

/**
 * @Description
 * @Author sjie
 * @Date 2022/6/2 14:47
 **/

// 使用viper读取配置文件

// type Redis config

type Redis struct {
	Hostname string `mapstructure:"host" yaml:"host"`
	Port     int    `mapstructure:"port" yaml:"port"`
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
	Password string `mapstructure:"password" json:"passowrd" yaml:"passowrd"`
}

type Server struct {
	Port         int    `yaml:"port"`
	Mode         string `yaml:"mode"`
	Host         string `yaml:"host"`
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
}
type MySQL struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DbName   string `yaml:"db_name"`
	Password string `yaml:"password"`
	UserName string `yaml:"userName"`
}

type Jwt struct {
	SecretKey    string
	JwtBlackList int64 `mapstructure:"jwt_black_list" json:"jwt_black_list" yaml:"jwt_black_list"`
}

type Config struct {
	Redis
	MySQL
	Server
	Jwt
}

func NewConfig() *Config {
	// 读取配置文件
	viper.SetConfigFile("./config.yaml") // 指定配置文件路径

	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {             // 处理读取配置文件的错误
		_ = fmt.Errorf("Fatal error config file: %s \n", err)
	}

	var c *Config

	if err := viper.Unmarshal(&c); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		return nil
	}
	return c
}
