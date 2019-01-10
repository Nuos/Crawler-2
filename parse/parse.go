package parse

import (
	"Crawler/util"
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func ParseFromGithub(document *goquery.Document) {
	document.Find("div.explore-content ol.repo-list li").Each(func(i int, selection *goquery.Selection) {
		RespName, _ := util.HandleCommon(selection.Find("div h3 a").Text())
		URL, _ := util.HandleUrl(selection.Find("div").Eq(0).Find("h3 a").AttrOr("href", "None"))
		Descript, _ := util.HandleCommon(selection.Find("div").Eq(2).Find("p").Text())
		Stars, _ := util.HandleCommon(selection.Find("div").Eq(3).Find("a").Eq(0).Text())
		Fork, _ := util.HandleCommon(selection.Find("div").Eq(3).Find("a").Eq(1).Text())
		TodayStars, _ := util.HandleCommon(selection.Find("div").Eq(3).Find("span").Eq(1).Text())
		fmt.Println("RespName:", RespName)
		fmt.Println("URL:", URL)
		fmt.Println("Descript:", Descript)
		fmt.Println("Stars:", Stars)
		fmt.Println("Fork", Fork)
		fmt.Println("TodayStars:", TodayStars)
	})
}
func ParseForDevelops(document *goquery.Document) {

}
