package config

import (
	"flag"
	"strings"
	"sync"
	"time"
)

type Config struct {
	Host    string
	Port    string
	Timeout time.Duration
}

var instance *Config
var once sync.Once
var timeout = flag.Duration("timeout", 10*time.Second, "таймаут на подключение к серверу")

func Get() *Config {
	host, port := "localhost", "3000"
	once.Do(func() {
		instance = &Config{}
		flag.Parse()

		resource := flag.Arg(0)
		instance.Timeout = *timeout
		instance.Host = host
		instance.Port = port

		if strings.Contains(resource, " ") {
			s := strings.Split(strings.Trim(resource, "\n"), " ")
			instance.Host, instance.Port = s[0], s[1]
		}
	})

	return instance
}
