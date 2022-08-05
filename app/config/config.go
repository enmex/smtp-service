package config

import (
	"os"
	"smtp/pkg/logger"

	"gopkg.in/yaml.v2"
)

type Provider struct {
	Credentials struct {
		User     string `json:"user" yaml:"user"`
		Password string `json:"password" yaml:"password"`
	} `json:"credentials" yaml:"credentials"`
	Delivery struct {
		Host    string `json:"host" yaml:"host"`
		Address string `json:"address" yaml:"address"`
	} `json:"delivery" yaml:"delivery"`
}

var (
	SingleMode   = "SINGLE"
	MultiplyMode = "MULTIPLY"
)

// Provider should be have default provider configuration
type Config struct {
	Mode      string              `json:"mode" yaml:"mode"`
	Providers map[string]Provider `json:"providers" yaml:"providers"`
}

var config Config

func Init() error {
	confFromFile, err := confFromFile("./usr/local/bin/config.yml")
	if err != nil {
		return err
	}
	config = *confFromFile

	return nil
}

func GetConfig() Config {
	return config
}

func confFromFile(fileName string) (*Config, error) {
	file, err := os.Open(fileName)
	if err != nil {
		logger.Logger.Errorf("[Config][confFromFile] Error opening config file: %s", err)
		return nil, err
	}

	var conf Config
	defer file.Close()

	if err := yaml.NewDecoder(file).Decode(&conf); err != nil {
		logger.Logger.Errorf("[Config][confFromFile] Error decoding config :%s",err.Error())
		return nil, err
	}

	logger.Logger.Info("Config loaded successfully")
	
	return &conf, nil
}
