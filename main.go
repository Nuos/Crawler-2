package main

import (
	"github.com/wentaojia2014/Crawler/engine"

	"github.com/wentaojia2014/Crawler/parse"
)

func main() {
	var simplerTest engine.Trending
	simplerTest.Run(
		engine.RequestForGithub{
			URL:       "https://github.com/trending",
			ParseFunc: parse.ParseFromGithub,
		},
	)
	simplerTest.Run(
		engine.RequestForGithub{
			URL:       "https://github.com/trending/developers",
			ParseFunc: parse.ParseForDevelops,
		},
	)
}
