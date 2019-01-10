package main

import (
	"fmt"

	"github.com/wentaojia2014/Crawler/download"
	"github.com/wentaojia2014/Crawler/parse"
)

func main() {
	fmt.Println("Hello Crawler")
	doc, err := download.Download("https://github.com/trending")
	if err != nil {
		fmt.Println(err.Error())
	}
	parse.ParseFromGithub(doc)
}
