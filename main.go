package main

import (
	appProps "AppProps/src"
	"log"
)

func init() {

}

func main() {
	config := appProps.UseProps("resources/application.properties")
	log.Println(config)
}
