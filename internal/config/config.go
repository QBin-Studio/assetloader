package config

var config = Config{
	AssetDir:    "",
	AssetPrefix: "IMG_",
	OutputFile:  "index.ts",
}

type Config struct {
	AssetDir    string
	AssetPrefix string
	OutputFile  string
}

func GetConfig() *Config {
	return &config
}

func SetConfig(c Config) {
	config = c
}
