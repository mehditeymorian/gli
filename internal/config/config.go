package config

type Config struct {
	Versions []string
	DB       []string
	HTTP     []string
	Log      []string
}

func Load() Config {
	return Default()
}
