package scraper

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

// Scraper - structure for general parsing functionality
type Scraper struct {
	URL       string
	UserAgent string
	ParseFunc func(*goquery.Document) (float64, float64, error)
}

// FetchData retrieves the HTML page and passes it to the parsing function
func (s *Scraper) FetchData() (float64, float64, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", s.URL, nil)
	if err != nil {
		return 0, 0, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("User-Agent", s.UserAgent)

	resp, err := client.Do(req)
	if err != nil {
		return 0, 0, fmt.Errorf("error fetching page: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return 0, 0, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return 0, 0, fmt.Errorf("error parsing HTML: %v", err)
	}

	return s.ParseFunc(doc)
}
