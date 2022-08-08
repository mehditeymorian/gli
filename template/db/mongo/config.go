package db

type Config struct {
	URI  string `koanf:"uri"`
	Name string `koanf:"name"`
}
