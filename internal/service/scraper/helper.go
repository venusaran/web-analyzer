package scraper

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	interfaces "github.com/venusaran/web-analyzer/pkg/interfaces"
	"golang.org/x/net/html"
)

// TODO: Add logs
func fetcher(url string, ch chan string, chFinished chan bool, data *interfaces.PageData) {
	defer func() {
		chFinished <- true
	}()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("ERROR: Failed to fetch:", url)
		return
	}
	defer resp.Body.Close()

	// check for HTML version from the doctype
	if docType, err := getHTMLVersion(resp); err == nil {
		data.HTMLVersion = docType
	}

	z := html.NewTokenizer(resp.Body)
	data.Headings = make(map[string]int)
	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			return
		case tt == html.DoctypeToken:
			data.HTMLVersion = "HTML5" // TODO:assuming HTML5 if doctype is found
		case tt == html.StartTagToken:
			t := z.Token()
			switch t.Data {
			case "title":
				tt = z.Next()
				if tt == html.TextToken {
					data.Title = z.Token().Data
				}
			case "h1", "h2", "h3", "h4", "h5", "h6":
				data.Headings[t.Data]++
			case "form":
				for _, attr := range t.Attr {
					if attr.Key == "action" {
						data.LoginForm = true
					}
				}
			case "a":
				ok, url := getHref(t)
				if ok && strings.HasPrefix(url, "http") {
					ch <- url
				}
			}
		}
	}
}

func getHTMLVersion(resp *http.Response) (string, error) {
	z := html.NewTokenizer(resp.Body)
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return "", fmt.Errorf("doctype not found")
		case html.DoctypeToken:
			t := z.Token()
			return strings.ToUpper(t.Data), nil
		}
	}
}

func getHref(t html.Token) (ok bool, href string) {
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
			ok = true
		}
	}
	return
}

func checkUrl(url string, results map[string]bool, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Head(url)
	if err != nil {
		results[url] = false
		return
	}
	defer resp.Body.Close()
	results[url] = resp.StatusCode == http.StatusOK
}
