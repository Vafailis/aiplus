package domain

type Config struct {
	Database DatabaseConfig `json:"database"`
	Server   ServerConfig   `json:"server"`
}

type DatabaseConfig struct {
	Url string `json:"url"`
	Key string `json:"key"`
}

type ServerConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}
