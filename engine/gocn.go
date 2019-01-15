package engine

import (
	"errors"
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/wentaojia2014/Crawler/domain"
	"github.com/wentaojia2014/Crawler/download"
	"github.com/wentaojia2014/Crawler/parse"
)

var (
	ErrorRun = errors.New("engine error")
)
var (
	ContentsNil = domain.Contents{}
)

func Run(request domain.Request) (domain.Contents, error) {
	var (
		doc *goquery.Document
		err error
	)
	if doc, err = download.Download2(request.Url); err != nil {
		fmt.Println(err)
		return ContentsNil, ErrorRun
	}
	return parse.TitleParse(doc), nil
}
