package config

type Config struct {
	Host     string
	Address  string
}

var config Config

func Init() {

	config = Config{
		Host:     "smtp.gmail.com:",
		Address:  "smtp.gmail.com:465",
	}
}

func GetConfig() Config {
	return config
}
