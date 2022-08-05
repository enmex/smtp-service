package config

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