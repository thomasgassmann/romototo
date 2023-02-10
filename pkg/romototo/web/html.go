package web

import (
	"net/http"

	"github.com/anaskhan96/soup"
)

type HtmlParser struct {
	url string
	doc *soup.Root
}

type AddHeaderTransport struct{
    T http.RoundTripper
}

func (adt *AddHeaderTransport) RoundTrip(req *http.Request) (*http.Response,error) {
    req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
    return adt.T.RoundTrip(req)
}

func (parser *HtmlParser) Init(url string) {
	parser.url = url
}

func (parser *HtmlParser) Fetch() error {
	client := http.Client{}
	client.Transport = &AddHeaderTransport{http.DefaultTransport}
	resp, err := soup.GetWithClient(parser.url, &client)
	if err != nil {
		return err
	}

	doc := soup.HTMLParse(resp)
	parser.doc = &doc
	return nil
}
