package src

import (
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

type ByLen []os.FileInfo

func (a ByLen) Len() int {
	return len(a)
}

func (a ByLen) Less(i, j int) bool {
	return len(a[i].Name()) < len(a[j].Name())
}

func (a ByLen) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

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

// remove file and use resource folder ...

func UseResource(resourceFolderPath string) Config {
	var propList = make(map[string]string)
	var accesptedProp string
	_, err := os.Stat(resourceFolderPath)
	if os.IsNotExist(err) {
		log.Printf("Folder does not exist, %s", err)
		return newConfig(propList)
	} else {
		files, err := ioutil.ReadDir(resourceFolderPath)
		if err != nil {
			return newConfig(propList)
		}

		sort.Sort(ByLen(files))

		for _, file := range files {
			if file.Name() == "application.properties" {
				availableConfig := UseProps(resourceFolderPath + "/" + file.Name())
				propList = availableConfig.appProp
				accesptedProp = availableConfig.Get("use.profile")
			} else {
				if accesptedProp == getAppPropsType(file.Name()) {
					newProps := UseProps(resourceFolderPath + "/" + file.Name())
					updateProps(propList, newProps.appProp)
				}

			}
			strings.Split(strings.Split(file.Name(), ".")[0], "-")
			log.Println(file.Name())
		}
	}
	return newConfig(propList)
}

func getAppPropsType(fileName string) string {
	return strings.Split(strings.Split(fileName, ".")[0], "-")[1]
}

func updateProps(currentProps, updatedProps map[string]string) map[string]string {
	for i, v := range updatedProps {
		currentProps[i] = v
	}
	return currentProps
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

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
