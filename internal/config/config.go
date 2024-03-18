package config

type Config struct {
	BindAddr    string `json:"bind_addr"`
	DatabaseURL string `json:"database_url"`
}

func NewConfig() *Config {

	return &Config{
		BindAddr:    ":8080",
		DatabaseURL: "user=user password=password host=postgres port=5432 dbname=mydb sslmode=disable",
	}
}
