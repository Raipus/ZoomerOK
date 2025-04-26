package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
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
	PageSize            int    `mapstructure:"PAGE_SIZE"`
	TimeCache           int    `mapstructure:"TIME_CACHE"`
	ConfirmationCache   string `mapstructure:"CONFIRMATION_CACHE"`
	ResetCache          string `mapstructure:"RESET_CACHE"`
	GenerateLinkLength  int    `mapstructure:"GENERATE_LINK_LENGTH"`
	GenerateLinkCharset string `mapstructure:"GENERATE_LINK_CHARSET"`

	KafkaReaderTopic string `mapstructure:"KAFKA_READER_TOPIC"`
	KafkaWriterTopic string `mapstructure:"KAFKA_WRITER_TOPIC"`
	KafkaBrokerHost  string `mapstructure:"KAFKA_BROKER_HOST"`
	KafkaBrokerPort  int    `mapstructure:"KAFKA_BROKER_PORT"`
	KafkaBrokerUrl   string
}

func LoadConfig() (c *ConfigType) {
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if !strings.HasSuffix(os.Args[0], ".test") {
		gin.SetMode(gin.ReleaseMode)
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
		gin.SetMode(gin.TestMode)
		viper.AddConfigPath("../config/envs")
		viper.SetConfigName("test")
		if err := viper.ReadInConfig(); err != nil {
			fmt.Printf("Ошибка при чтении файла test: %s\n", err)
		}
	}

	viper.SetConfigName("kafka")
	if err := viper.MergeInConfig(); err != nil {
		fmt.Printf("Ошибка при объединении файла kafka.env: %s\n", err)
	}

	c = new(ConfigType)

	if err := viper.Unmarshal(&c); err != nil {
		panic(fmt.Errorf("невозможно декодировать в структуру: %w", err))
	}

	c.KafkaBrokerUrl = c.KafkaBrokerHost + ":" + strconv.Itoa(c.KafkaBrokerPort)
	fmt.Println("Viper config dump:")
	allSettings := viper.AllSettings()
	for key, value := range allSettings {
		fmt.Printf("%s: %v\n", key, value)
	}
	return
}

var Config *ConfigType = LoadConfig()
