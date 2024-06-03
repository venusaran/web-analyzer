package scraper

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"

	interfaces "github.com/venusaran/web-analyzer/pkg/interfaces"
	"golang.org/x/net/html"
)

func fetcher(pageURL string, ch chan string, chFinished chan bool, pageInfo *interfaces.PageData) {
	defer func() {
		chFinished <- true
	}()

	// check the URL is valid one first
	baseURL, err := url.Parse(pageURL)
	if err != nil {
		fmt.Println("ERROR: Invalid base URL:", pageURL)
		return
	}

	// make initial page fetch
	resp, err := http.Get(pageURL)
	if err != nil {
		fmt.Println("ERROR: Failed to fetch:", pageURL)
		return
	}
	defer resp.Body.Close()

	z := html.NewTokenizer(resp.Body)
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return
		case html.DoctypeToken:
			pageInfo.HTMLVersion, _ = getHTMLVersion(z.Token())
		case html.StartTagToken, html.SelfClosingTagToken:
			t := z.Token()
			switch t.Data {
			case "title":
				pageInfo.Title = getTitle(z)
			case "h1", "h2", "h3", "h4", "h5", "h6":
				pageInfo.Headings[t.Data]++
			case "a":
				ok, href := getHref(t)
				if ok && strings.HasPrefix(href, "http") {
					parsedURL, err := url.Parse(href)
					if err == nil {
						if parsedURL.Host == baseURL.Host {
							pageInfo.InternalLinks++
						} else {
							pageInfo.ExternalLinks++
						}
					}
					ch <- href
				}
			case "form":
				pageInfo.LoginForm = pageInfo.LoginForm || isLoginForm(t)
			}
		}
	}
}

func getHTMLVersion(token html.Token) (string, error) {
	docType := strings.ToLower(token.Data)

	if strings.Contains(docType, "html") {
		if strings.Contains(docType, "4.01") {
			return "HTML 4.01", nil
		} else if strings.Contains(docType, "xhtml 1.0") {
			return "XHTML 1.0", nil
		} else if strings.Contains(docType, "xhtml 1.1") {
			return "XHTML 1.1", nil
		} else if strings.Contains(docType, "xhtml 1.2") {
			return "XHTML 1.2", nil
		} else {
			return "HTML5", nil
		}
	}
	return "Unknown HTML version", nil
}

func getTitle(z *html.Tokenizer) string {
	tt := z.Next()
	if tt == html.TextToken {
		return z.Token().Data
	}
	return ""
}

func getHref(t html.Token) (bool, string) {
	for _, a := range t.Attr {
		if a.Key == "href" {
			return true, a.Val
		}
	}
	return false, ""
}

func isLoginForm(t html.Token) bool {
	for _, a := range t.Attr {
		if a.Key == "id" && strings.Contains(strings.ToLower(a.Val), "login") {
			return true
		}
	}
	return false
}

func checkUrl(url string, d *interfaces.PageData, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url)
	// count it as InaccessibleLink if there is an error while fetching
	if err != nil {
		d.AccessibleURLs[url] = false
		d.InaccessibleLinks++
		return
	}

	// also count it as InaccessibleLink if the http status code is 200
	if resp.StatusCode != 200 {
		d.AccessibleURLs[url] = false
		d.InaccessibleLinks++
		return
	}

	defer resp.Body.Close()
	d.AccessibleURLs[url] = resp.StatusCode == http.StatusOK
}
