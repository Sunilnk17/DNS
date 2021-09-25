package config

import (
	"github.com/gorilla/schema"
	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"os"
	"strings"
)

var config *Config
var requestParamsValidator *validator.Validate
var requestParamsDecoder *schema.Decoder

type Config struct {
	Server  ServerConfig
	Sectors []SectorConfig
}

type ServerConfig struct {
	Port string
}

type SectorConfig struct {
	Id        string
	Companies []CompanyConfig
}

type CompanyConfig struct {
	Name     string
	Response string
}

func GetConfig() *Config {
	return config
}

func SetDummyConfig() {
	config = &Config{}
}

func Initialize() error {
	viper.SetConfigFile(getConfigFile())
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	for _, k := range viper.AllKeys() {
		value := viper.GetString(k)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			viper.Set(k, os.Getenv((strings.TrimSuffix(strings.TrimPrefix(value, "${"), "}"))))
		}
	}
	return viper.Unmarshal(&config)
}

func GetReqParamsValidator() *validator.Validate {
	return requestParamsValidator
}

func GetReqParamsDecoder() *schema.Decoder {
	return requestParamsDecoder
}

func getConfigFile() string {
	return "app/config/dev.yml"
}

func InitializeDecoderAndValidator() {
	log.Print("Initializing params decoder & validator")
	requestParamsDecoder = schema.NewDecoder()
	requestParamsValidator = validator.New()
	requestParamsDecoder.IgnoreUnknownKeys(true)
}
