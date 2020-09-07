package app

import (
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"net/http"
)

type HtmlElement struct {
	resp  *http.Response
	s     *goquery.Selection
	attrs []html.Attribute
	text  string
}

func NewHtmlElement(resp *http.Response, s *goquery.Selection, n *html.Node) *HtmlElement {
	return &HtmlElement{resp: resp, s: s, attrs: n.Attr, text: goquery.NewDocumentFromNode(n).Text()}
}

func (h *HtmlElement) Attr(k string) string {
	for _, attr := range h.attrs {
		if attr.Key == k {
			return attr.Val
		}
	}
	return ""
}

func (h *HtmlElement) Text() string {
	return h.text
}
