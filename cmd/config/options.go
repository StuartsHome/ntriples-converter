package config

type Namespace string
type Config struct {
	OutputPath string
	Namespace  Namespace
}

func New(ns string, outputPath string) *Config {
	return &Config{
		OutputPath: outputPath,
		Namespace:  Namespace(ns),
	}
}
