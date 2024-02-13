package util

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Username string `yaml:"user"`
		Host string `yaml:"host"`
		Port int `yaml:"port"`
		Password string `yaml:"pass" envconfig:"DB_USERNAME"`
	} `yaml:"database"`
}

func Load_config() *Config {
	var cfg *Config
	var cfgPath string
	env := os.Getenv("SERVER_ENV")
	if env != "dev" && env != "prod" {
		log.Fatalln("SERVER_ENV environment variable must be set to dev or prod")
	}
	ex, err := os.Executable()
	if err != nil {
		log.Fatalln(err)
	}
	cfgPath = filepath.Join(
		filepath.Dir(ex),
		"config",
		fmt.Sprintf("%v.config.yaml", env),
	)
	f, err := os.Open(cfgPath)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	err = envconfig.Process("", cfg)
	if err != nil {
		log.Fatalln(err)
	}
	return cfg
}
