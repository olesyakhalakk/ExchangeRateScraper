package scraper

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// parseKytObmin parses the HTML page to extract the USD/UAH exchange rate
func parseKytObmin(doc *goquery.Document) (float64, float64, error) {
	var rates []float64

	doc.Find("a[href*='usd-uah']").Each(func(i int, s *goquery.Selection) {
		rateStr := strings.TrimSpace(strings.ReplaceAll(s.Text(), ",", "."))
		if rateStr == "" {
			return
		}

		rate, err := strconv.ParseFloat(rateStr, 64)
		if err == nil {
			rates = append(rates, rate)
		}
	})

	if len(rates) < 2 {
		return 0, 0, ErrNoDataFound
	}

	return rates[0], rates[1], nil
}

// GetKytObminRate retrieves the USD/UAH exchange rate from KytObmin
func GetKytObminRate() (float64, float64, error) {
	doc, err := FetchHTML("https://kyt-obmin.if.ua")
	if err != nil {
		return 0, 0, err
	}

	return parseKytObmin(doc)
}
