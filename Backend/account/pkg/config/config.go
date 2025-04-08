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

	Host                string `mapstructure:"HOST"`
	HttpPort            int    `mapstructure:"HTTP_PORT"`
	Prefix              string `mapstructure:"PREFIX"`
	SecretKey           string `mapstructure:"SECRET_KEY"`
	TimeCache           int    `mapstructure:"TIME_CACHE"`
	ConfirmationCache   string `mapstructure:"CONFIRMATION_CACHE"`
	ResetCache          string `mapstructure:"RESET_CACHE"`
	GenerateLinkLength  int    `mapstructure:"GENERATE_LINK_LENGTH"`
	GenerateLinkCharset string `mapstructure:"GENERATE_LINK_CHARSET"`
	Photo               PhotoConfig
}

type PhotoConfig struct {
	ImagePath string `mapstructure:"IMAGE_PATH"`
	Default   string `mapstructure:"DEFAULT"`
	Image     string
	ByteImage []byte
	Small     uint `mapstructure:"SMALL"`
	Large     uint `mapstructure:"LARGE"`
}

func LoadConfig() (c *ConfigType) {
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if !strings.HasSuffix(os.Args[0], ".test") {
		viper.AddConfigPath("./pkg/config/envs")

		viper.SetConfigName("prod")
		if err := viper.ReadInConfig(); err != nil {
			fmt.Printf("Ошибка при чтении файла prod: %s\n", err)
		}

		viper.SetConfigName("prod.db")
		if err := viper.MergeInConfig(); err != nil {
			fmt.Printf("Ошибка при объединении файла prod.db.env: %s\n", err)
		}
	} else {
		viper.AddConfigPath("../config/envs")
		viper.SetConfigName("test")
		if err := viper.ReadInConfig(); err != nil {
			fmt.Printf("Ошибка при чтении файла test: %s\n", err)
		}
	}

	viper.SetConfigName("image")
	if err := viper.MergeInConfig(); err != nil {
		fmt.Printf("Ошибка при объединении файла image.env: %s\n", err)
	}

	viper.SetConfigName("smtp")
	if err := viper.MergeInConfig(); err != nil {
		fmt.Printf("Ошибка при объединении файла smtp.env: %s\n", err)
	}

	c = new(ConfigType)

	if err := viper.Unmarshal(&c); err != nil {
		panic(fmt.Errorf("невозможно декодировать в структуру: %w", err))
	}

	fmt.Println("Viper config dump:")
	allSettings := viper.AllSettings()
	for key, value := range allSettings {
		fmt.Printf("%s: %v\n", key, value)
	}

	c.Photo.Large = viper.GetUint("large")
	c.Photo.Small = viper.GetUint("small")
	c.Photo.ImagePath = viper.GetString("image_path")
	c.Photo.Default = viper.GetString("default")

	c.Photo.Image = c.Photo.ImagePath + "/" + c.Photo.Default

	byteImage := getByteImage(c)
	c.Photo.ByteImage = byteImage
	return
}

var Config *ConfigType = LoadConfig()
