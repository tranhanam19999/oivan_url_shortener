package config

type Config struct {
	Stage string    `env:"STAGE" envDefault:"local"`
	App   AppConfig `envPrefix:"APP_"`
	DB    DBConfig  `envPrefix:"DB_"`
}

type AppConfig struct {
	Host    string `env:"HOST" envDefault:"localhost"`
	Port    string `env:"PORT" envDefault:"8080"`
	BaseURL string `env:"BASE_URL" envDefault:"http://localhost:8080"`
}

type DBConfig struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     string `env:"PORT" envDefault:"5432"`
	User     string `env:"USER" envDefault:"user"`
	Password string `env:"PASSWORD" envDefault:"password"`
	Name     string `env:"NAME" envDefault:"demo_oivan"`
	SSLMode  string `env:"SSLMODE" envDefault:"disable"`
}
