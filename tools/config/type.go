package config

type Config struct {
	Port string   `env:"PORT" envDefault:"8080"`
	DB   DBConfig `envPrefix:"DB_"`
}

type DBConfig struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     string `env:"PORT" envDefault:"5432"`
	User     string `env:"USER" envDefault:"user"`
	Password string `env:"PASSWORD" envDefault:"password"`
	Name     string `env:"NAME" envDefault:"demo_oivan"`
	SSLMode  string `env:"SSLMODE" envDefault:"disable"`
}
