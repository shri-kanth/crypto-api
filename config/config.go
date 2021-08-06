package config

var API_KEY string = "903552914cec4c6d4c38c2da7849746d20bdb231"

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Charset  string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Host:     "127.0.0.1",
			Port:     3306,
			Username: "guest",
			Password: "guest!",
			Name:     "todoapp",
			Charset:  "utf8",
		},
	}
}
