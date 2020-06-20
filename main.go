package main

import (
	"AppProps/src"
)

func init() {

}

func main() {
	config := src.UseResource("resources/")
	config.Print()

}
