package core

type AppConfig struct {
	// Required envs
	Port      string `env:"PORT, default=8080"`
	AppName   string `env:"APP_NAME"`
	ProjectId string `env:"PROJECT_ID"`
	DebugMode bool   `env:"DEBUG, default=false"`
	// Please add here your envs variables and their default values
	MongoUrl      string `env:"MONGO_URL, default="`
	MongoDatabase string `env:"MONGO_DATABASE, default="`
}
