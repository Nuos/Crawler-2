package main

import (
	"Crawler/download"
	"Crawler/parse"
	"fmt"
)

func main() {
	fmt.Println("Hello Crawler")
	doc, err := download.Download("https://github.com/trending")
	if err != nil {
		fmt.Println(err.Error())
	}
	parse.ParseFromGithub(doc)
}