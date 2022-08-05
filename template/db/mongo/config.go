package mongo

type Config struct {
	URI  string `koanf:"uri"`
	Name string `koanf:"name"`
}
