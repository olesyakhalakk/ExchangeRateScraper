package scraper

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

// parseTopObmin parses the HTML and extracts the USD/UAH exchange rate
func parseTopObmin(doc *goquery.Document) (float64, float64, error) {
	var rates []float64

	doc.Find("td.green, td.red").Each(func(i int, s *goquery.Selection) {
		rate, err := CleanPriceString(s.Text())
		if err == nil {
			rates = append(rates, rate)
		}
	})

	if len(rates) < 2 {
		return 0, 0, fmt.Errorf("failed to find USD/UAH exchange rate")
	}

	return rates[0], rates[1], nil
}

// GetTopObminRate fetches the exchange rate from TopObmin using Scraper
func GetTopObminRate() (float64, float64, error) {
	doc, err := FetchHTML("https://topobmin.com/")
	if err != nil {
		return 0, 0, err
	}
	return parseTopObmin(doc)
}
