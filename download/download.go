package download

import (
	"errors"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var (
	ErrorNil          = errors.New("response is nil")
	ErrorWrongCode    = errors.New("http response code is wrong")
	ErrorHttpRequest  = errors.New("http request error")
	ErrorHttpResponse = errors.New("http response error")
)

func Download(url string) (*goquery.Document, error) {
	var (
		resp *http.Response
		err  error
	)
	if resp, err = http.Get(url); err != nil {
		return nil, ErrorNil
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, ErrorWrongCode
	}
	return goquery.NewDocumentFromReader(resp.Body)
}
func Download2(url string) (*goquery.Document, error) {
	var (
		req *http.Request
		err error
	)
	if req, err = http.NewRequest("GET", url, nil); err != nil {
		return nil, ErrorHttpRequest
	}
	client := http.DefaultClient
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.139 Safari/537.36")
	var (
		resp *http.Response
	)
	if resp, err = client.Do(req); err != nil {
		return nil, ErrorHttpResponse
	}
	defer resp.Body.Close()
	return goquery.NewDocumentFromReader(resp.Body)
}
