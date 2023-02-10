package web

type Browser struct {
	Driver *BrowserDriver
	Parser *HtmlParser
}

func ConstructBrowser(driver *BrowserDriver, parser *HtmlParser) Browser {
	return Browser{
		Driver: driver,
		Parser: parser,
	}
}
