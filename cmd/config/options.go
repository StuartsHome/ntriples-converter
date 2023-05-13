package config

type Namespace string
type Config struct {
	Namespace  Namespace
	InputPath  string
	OutputPath string
}

func New(ns Namespace, inputPath, outputPath string) *Config {
	return &Config{
		Namespace:  Namespace(ns),
		InputPath:  inputPath,
		OutputPath: outputPath,
	}
}
