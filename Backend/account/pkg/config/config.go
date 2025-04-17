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

	SmtpUsername string `mapstructure:"SMTP_USERNAME"`
	SmtpPassword string `mapstructure:"SMTP_PASSWORD"`
	SmtpHost     string `mapstructure:"SMTP_HOST"`
	SmtpPort     string `mapstructure:"SMTP_PORT"`

	KafkaAccountBlogTopic string `mapstructure:"KAFKA_ACCOUNT_BLOG_TOPIC"`
	KafkaBrokerHost       string `mapstructure:"KAFKA_BROKER_HOST"`
	KafkaBrokerPort       int    `mapstructure:"KAFKA_BROKER_PORT"`
	KafkaBrokerUrl        string

	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPort     int    `mapstructure:"REDIS_PORT"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
	RedisUrl      string
}

type PhotoConfig struct {
	ImagePathProd string `mapstructure:"IMAGE_PATH_PROD"`
	ImagePathTest string `mapstructure:"IMAGE_PATH_TEST"`
	Default       string `mapstructure:"DEFAULT"`
	Image         string
	ByteImage     []byte
	Small         uint `mapstructure:"SMALL"`
	Large         uint `mapstructure:"LARGE"`
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

	viper.SetConfigName("kafka")
	if err := viper.MergeInConfig(); err != nil {
		fmt.Printf("Ошибка при объединении файла kafka.env: %s\n", err)
	}

	viper.SetConfigName("redis")
	if err := viper.MergeInConfig(); err != nil {
		fmt.Printf("Ошибка при объединении файла redis.env: %s\n", err)
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

	c.KafkaBrokerUrl = c.KafkaBrokerHost + ":" + string(c.KafkaBrokerPort)
	c.RedisUrl = c.RedisHost + ":" + string(c.RedisPort)

	c.Photo.Large = viper.GetUint("large")
	c.Photo.Small = viper.GetUint("small")
	c.Photo.ImagePathProd = viper.GetString("image_path_prod")
	c.Photo.ImagePathTest = viper.GetString("image_path_test")
	c.Photo.Default = viper.GetString("default")

	if !strings.HasSuffix(os.Args[0], ".test") {
		c.Photo.Image = c.Photo.ImagePathProd + "/" + c.Photo.Default
	} else {
		c.Photo.Image = c.Photo.ImagePathTest + "/" + c.Photo.Default
	}

	byteImage := getByteImage(c)
	c.Photo.ByteImage = byteImage
	return
}

var Config *ConfigType = LoadConfig()
