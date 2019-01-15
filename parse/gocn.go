package parse

import (
	"errors"

	"github.com/wentaojia2014/Crawler/download"

	"github.com/wentaojia2014/Crawler/util"

	"github.com/PuerkitoBio/goquery"
	"github.com/wentaojia2014/Crawler/domain"
)

var (
	ErrorParse = errors.New("parse error")
)

func TitleParse(document *goquery.Document) domain.Contents {
	var (
		allContents domain.Contents
		oneContent  domain.Content
	)
	document.Find("div.aw-common-list div.aw-item").Each(func(i int, selection *goquery.Selection) {
		var (
			user           string
			userHome       string
			url            string
			passageContent string
			title          string
			tag            string
			lastAnsewe     string
			one            int
			two            int
			three          int
		)
		user = util.StringSplit(selection.Find("a").Eq(0).AttrOr("href", "None"))
		userHome = selection.Find("a").Eq(0).AttrOr("href", "None")
		url = selection.Find("div h4 a").AttrOr("href", "None")
		doc, _ := download.Download2(url)
		passageContent, _ = PassageParse(doc)
		title = util.StringTrim(selection.Find("div h4").Text())
		tag = selection.Find("p a").Eq(0).Text()
		lastAnsewe = selection.Find("p a").Eq(1).Text()
		one, two, three = util.StringSplitByDot(selection.Find("p span").Eq(0).Text())

		oneContent = domain.Content{
			URL:          url,
			Title:        title,
			Tag:          tag,
			Reporter:     user,
			ReporterHome: userHome,
			Followers:    one,
			Answers:      two,
			Answerer:     lastAnsewe,
			See:          three,
			Passage:      passageContent,
		}
		allContents = append(allContents, oneContent)
	})
	return allContents
}
func PassageParse(doc *goquery.Document) (string, error) {
	var content string
	content = util.StringTrim(doc.Find("div.content.markitup-box").Text())
	return util.StringsReplace(content), nil
}
