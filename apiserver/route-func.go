package apiserver

const (
	startURL = "https://gocn.io/sort_type-new__day-0__is_recommend-0"
	rootURL  = "https://gocn.io/sort_type-new__day-0__is_recommend-0__page-%s"
)
func GetContents(request *restful.Request, response *restful.Response) {
	var (
		page string
		url string
	)
	page = request.PathParameter("page")
	if page == "" {
		url = startURL
	} else {
		url = fmt.Sprintf(rootURL, page)
	}
	contents, _ := engine.Run(
		domain.Request{
			Url: url,
			ParseFunc: gocn.TitleParse,
		},
	)
	response.WriteEntity(contents)
}
