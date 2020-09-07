package app

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"sync/atomic"
)

type Collector struct {
	next    *Collector
	name    string
	runtime int
	client  *http.Client
	link    *Link
	parse   *parser
	over    int32 // 0-结束 1-运行
	stopCh  chan struct{}
}

type parser struct {
	pattern string
	f       func(e *HtmlElement) string
}

func InitCollector(link *Link, name string, runtime int, stopCh chan struct{}) *Collector {
	c := &Collector{}
	c.name = name
	c.runtime = runtime
	c.client = &http.Client{}
	c.link = link
	c.over = 0
	c.stopCh = stopCh
	return c
}

func (c *Collector) SetParse(pattern string, f func(e *HtmlElement) string) {
	c.parse = &parser{pattern: pattern, f: f}
}

func (c *Collector) collect(in chan string) chan string {
	out := make(chan string)
	go func() {
		defer func() {
			close(out)
			atomic.CompareAndSwapInt32(&c.over, 1, 0)
		}()
		for url := range in {
			// stop check
			select {
			case <-c.link.stopCh:
				return
			default:
			}
			atomic.CompareAndSwapInt32(&c.over, 0, 1)
			// request
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				return
			}
			resp, err := c.client.Do(req)
			if err != nil {
				return
			}
			// parse
			doc, err := goquery.NewDocumentFromReader(resp.Body)
			if err != nil {
				return
			}
			doc.Find(c.parse.pattern).Each(func(_ int, s *goquery.Selection) {
				for _, n := range s.Nodes {
					e := NewHtmlElement(resp, s, n)
					// save
					u := c.parse.f(e)
					out <- u
				}
			})
		}
	}()
	return out
}
