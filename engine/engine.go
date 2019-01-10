package engine

import (
	"errors"
	"fmt"

	"github.com/wentaojia2014/Crawler/download"
)

var (
	ErrorDocWrong = errors.New("document Wrong")
)

type Trending struct {
}

func (t Trending) Run(request RequestForGithub) {
	doc, err := download.Download(request.URL)
	if err != nil {
		fmt.Println(ErrorDocWrong)
		return
	}
	if doc != nil {
		fmt.Println("Game start!")
		request.ParseFunc(doc)
	} else {
		fmt.Println("Game over!")
	}
}
