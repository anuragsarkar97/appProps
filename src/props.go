package src

import (
	"log"
)



type Config struct {
	appProp map[string]string
}

func (c *Config) Get(key string) string {
	i, ok := c.appProp[key]
	if ok {
		return i
	}
	return ""
}

func (c *Config) Print() {
	for key, value := range c.appProp {
		log.Printf("key is : %s and value is : %s", "\033[31m"+key+"\033[0m", "\033[31m"+value+"\033[0m")
	}
}

func newConfig(appProps map[string]string) Config {
	c := new(Config)
	c.appProp = appProps
	return *c
}

