package config

type AppConfig struct {
	Port          string `env:"PORT, default=8080"`
	AppName       string `env:"APP_NAME, default="`
	MongoUrl      string `env:"MONGO_URL, default="`
	ProjectId     string `env:"PROJECT_ID, default="`
	MongoDatabase string `env:"MONGO_DATABASE, default="`
}
