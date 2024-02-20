package util

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Host       string `yaml:"host"`
		Port       int    `yaml:"port"`
		MPDBaseURL string `yaml:"mpd_base_url"`
	} `yaml:"server"`
	Database struct {
		Username string `yaml:"user"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Name     string `yaml:"name"`
		Password string `yaml:"pass" envconfig:"DB_PASSWORD"`
	} `yaml:"database"`
}

var Cfg *Config

func Load_config() *Config {
	cfg := new(Config)
	var cfgPath string
	env := os.Getenv("SERVER_ENV")
	if env != "dev" && env != "prod" && env != "cloud" {
		ErrorLogger.Fatalln("SERVER_ENV environment variable must be set to dev, prod, or cloud")
	}
	ex, err := os.Executable()
	if err != nil {
		ErrorLogger.Fatalln(err)
	}
	cfgPath = filepath.Join(
		filepath.Dir(filepath.Dir(ex)),
		"config",
		fmt.Sprintf("%v.config.yaml", env),
	)
	f, err := os.Open(cfgPath)
	if err != nil {
		ErrorLogger.Fatalln(err)
	}
	defer f.Close()
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		ErrorLogger.Fatalln(err)
	}
	err = envconfig.Process("", cfg)
	if err != nil {
		ErrorLogger.Fatalln(err)
	}
	if cfg.Database.Password == "" {
		ErrorLogger.Fatalln("The DB_PASSWORD environment variable is not set")
	}
	return cfg
}
