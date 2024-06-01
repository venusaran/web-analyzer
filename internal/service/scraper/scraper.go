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

func (srv *ScraperService) RetrieveData(url string) (linkStats map[string]bool, data interfaces.PageData, err error) {
	linkStats = make(map[string]bool)
	foundUrls := make(map[string]bool)

	// make channels
	chUrls := make(chan string)
	chFinished := make(chan bool)

	// fetch URLs and extract data from the initial page
	go fetcher(url, chUrls, chFinished, &data)

	// subscribe to both channels
	var wg sync.WaitGroup
	for finished := false; !finished; {
		select {
		case url := <-chUrls:
			if !foundUrls[url] {
				foundUrls[url] = true
				wg.Add(1)
				go checkUrl(url, linkStats, &wg)
			}
		case <-chFinished:
			finished = true
		}
	}

	wg.Wait()
	close(chUrls)
	return
}
