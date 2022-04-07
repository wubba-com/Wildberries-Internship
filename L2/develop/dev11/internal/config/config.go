package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

const (
	config = "config.yml"
)

type Config struct {
	configFile string
	DataFile   string `yaml:"data_file"`
	Listen     struct {
		Port   string `yaml:"port"`
		BindIP string `yaml:"bind_ip"`
	} `yaml:"listen"`
	DB struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Name     string `yaml:"name"`
	} `yaml:"db"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		err := cleanenv.ReadConfig(config, instance)
		if err != nil {
			log.Printf("err: %s\n", err.Error())
			return
		}
		instance.configFile = config
	})

	return instance
}
