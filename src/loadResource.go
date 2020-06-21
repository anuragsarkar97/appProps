package src

import (
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

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
				availableConfig := useProps(resourceFolderPath + "/" + file.Name())
				propList = availableConfig.appProp
				accesptedProp = availableConfig.Get("use.profile")
			} else {
				if accesptedProp == getAppPropsType(file.Name()) {
					newProps := useProps(resourceFolderPath + "/" + file.Name())
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

func useProps(filePath string) Config {
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