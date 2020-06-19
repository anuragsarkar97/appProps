package appProps

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
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

func newConfig(appProps map[string]string) Config {
	c := new(Config)
	c.appProp = appProps
	return *c 
}

func UseProps(filePath string) Config {
	var propList = make(map[string]string)
	if valid(filePath) {
		data := loadData(filePath)
		setVariables(data, propList)
	}
	return newConfig(propList)

}

func setVariables(data []byte, propList map[string]string) {
	stringData := string(data)
	splitArray := strings.Split(stringData, "\n")
	for _, v := range splitArray {
		if validResourceEntry(v) {
			addResourceEntry(v, propList)
			//log.Println(v)
		}
	}
	//log.Println(propList)
}

func addResourceEntry(pair string, propList map[string]string) {
	data := strings.Split(pair, "=")
	propList[data[0]] = envReplacer(data[1])
}

func validResourceEntry(v string) bool {
	if len(strings.Split(v, "=")) == 2 {
		return true
	}
	return false
}

func valid(filePath string) bool {

	fileSplit := strings.Split(filePath, ".")
	if "properties" != fileSplit[len(fileSplit)-1] {
		return false
	}
	return true
}

func loadData(filePath string) []byte {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("cannot load resource propeties %s", err.Error())
	}
	return data

}

func envReplacer(value string) string {
	if len(value) < 3 {
		return value
	}

	if value[0:2] == "${" {
		value = value[2 : len(value)-1]
		valueSplit := strings.Split(value, ":")

		value = getEnv(valueSplit[0], strings.Join(valueSplit[1:], ""))
		return value
	}
	return value

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
