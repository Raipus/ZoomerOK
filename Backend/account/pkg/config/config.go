package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type ConfigType struct {
	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	PostgresPort     int    `mapstructure:"POSTGRES_PORT"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresDb       string `mapstructure:"POSTGRES_DB"`

	Host     string `mapstructure:"HOST"`
	HttpPort int    `mapstructure:"HTTP_PORT"`
	Prefix   string `mapstructure:"PREFIX"`
}

func LoadConfig() (c *ConfigType) {
	if !strings.HasSuffix(os.Args[0], ".test") {
		viper.SetConfigName("prod")
		viper.AddConfigPath("./pkg/config/envs")
	} else {
		if !strings.HasSuffix(os.Args[0], "db.test") {
			viper.AddConfigPath("../config/envs")
		} else {
			viper.AddConfigPath("../config/envs")
		}
		viper.SetConfigName("test")
	}

	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}

	viper.SetConfigName("prod.db") // Имя второго файла без расширения
	if err := viper.MergeInConfig(); err != nil {
		panic(fmt.Errorf("fatal error merging config file: %s \n", err))
	}

	if err := viper.Unmarshal(&c); err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}

	return
}

var Config *ConfigType = LoadConfig()
