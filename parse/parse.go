package parse

import (
	"fmt"
	"os"

	"github.com/wentaojia2014/Crawler/util"

	"github.com/PuerkitoBio/goquery"
)

//parse https://github.com/trending
func ParseFromGithub(document *goquery.Document) {
	githubFile := "github.txt"
	fout, err := os.Create(githubFile)
	defer fout.Close()
	if err != nil {
		fmt.Println(githubFile, err)
		return
	}
	document.Find("div.explore-content ol.repo-list li").Each(func(i int, selection *goquery.Selection) {
		RespName, _ := util.HandleCommon(selection.Find("div h3 a").Text())
		URL, _ := util.HandleUrl(selection.Find("div").Eq(0).Find("h3 a").AttrOr("href", "None"))
		Descript, _ := util.HandleCommon(selection.Find("div").Eq(2).Find("p").Text())
		Stars, _ := util.HandleCommon(selection.Find("div").Eq(3).Find("a").Eq(0).Text())
		Fork, _ := util.HandleCommon(selection.Find("div").Eq(3).Find("a").Eq(1).Text())
		TodayStars, _ := util.HandleCommon(selection.Find("div").Eq(3).Find("span").Eq(1).Text())
		fmt.Println("---------begin---------")
		fmt.Println("RespName:", RespName)
		fmt.Println("URL:", URL)
		fmt.Println("Descript:", Descript)
		fmt.Println("Stars:", Stars)
		fmt.Println("Fork", Fork)
		fmt.Println("TodayStars:", TodayStars)
		fmt.Println("---------end---------")

		fout.WriteString("\n---------begin---------")
		fout.WriteString("\nRespName:" + RespName)
		fout.WriteString("\nURL:" + URL)
		fout.WriteString("\nDescript:" + Descript)
		fout.WriteString("\nStars:" + Stars)
		fout.WriteString("\nFork" + Fork)
		fout.WriteString("\nTodayStars:" + TodayStars)
		fout.WriteString("\n---------end---------")
	})
}

//parse https://github.com/trending/developers
func ParseForDevelops(document *goquery.Document) {
	developerFile := "developer.txt"
	fout, err := os.Create(developerFile)
	defer fout.Close()
	if err != nil {
		fmt.Println(developerFile, err)
		return
	}
	document.Find("div.explore-content ol li").Each(func(i int, selection *goquery.Selection) {
		DevName, _ := util.HandleCommon(selection.Find("li div div").Eq(1).Find("h2 a").Text())
		Descript, _ := util.HandleCommon(selection.Find("li div div").Eq(1).Find("a span").Text())
		URL, _ := util.HandleCommon(selection.Find("li div div").Eq(1).Find("h2 a").AttrOr("href", "None"))
		fmt.Println("---------begin---------")
		fmt.Println("DevName:", DevName)
		fmt.Println("Descript:", Descript)
		fmt.Println("URL:", URL)
		fmt.Println("---------end---------")
		fout.WriteString("\n---------begin---------")
		fout.WriteString("\nDevName:" + DevName)
		fout.WriteString("\nDescript:" + Descript)
		fout.WriteString("\nURL:" + URL)
		fout.WriteString("\n---------end---------")
	})
}
