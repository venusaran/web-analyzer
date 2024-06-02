package scraper

import (
	"context"
	"sync"

	interfaces "github.com/venusaran/web-analyzer/pkg/interfaces"
)

type ScraperService struct{}

func NewService(ctx context.Context) (l *ScraperService, err error) {

	l = &ScraperService{}
	return
}

func (srv *ScraperService) RetrieveData(url string) (resp interfaces.PageData, err error) {
	resp.Headings = make(map[string]int)
	resp.AccessibleURLs = make(map[string]bool)

	// make channels
	chUrls := make(chan string)
	// channel to notify all the tasks are finished execution
	chFinished := make(chan bool)

	// fetch data from the initial page(the one received from user)
	go fetcher(url, chUrls, chFinished, &resp)

	// subscribe to both channels
	var wg sync.WaitGroup
	for finished := false; !finished; {
		select {
		case url := <-chUrls:
			if _, exists := resp.AccessibleURLs[url]; !exists {
				resp.AccessibleURLs[url] = false
				wg.Add(1)
				go checkUrl(url, &resp, &wg)
			}
		case <-chFinished:
			finished = true
		}
	}

	wg.Wait()
	close(chUrls)

	return
}
