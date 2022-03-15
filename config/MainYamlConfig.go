package config

type MainYamlConfig struct {
	Core CoreConfig
}

type CoreConfig struct {
	Server ServerConfig `yaml:"server"`
	User   UserConfig   `yaml:"user"`
}

type ServerConfig struct {
	IsSsl string `yaml:"ssl"`
	Host  string `yaml:"host"`
	Port  int    `yaml:"port"`
}

type UserConfig struct {
	Username string `yaml:"username"`
	Token    string `yaml:"token"`
}
