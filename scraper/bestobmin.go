package scraper

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// parseBestObmin parses the BestObmin HTML page and extracts the exchange rate
func parseBestObmin(doc *goquery.Document) (float64, float64, error) {
	buyStr := strings.TrimSpace(doc.Find("div.digit_bg.left_digit_bg p").First().Text())
	sellStr := strings.TrimSpace(doc.Find("div.digit_bg.right_digit_bg p").First().Text())

	if buyStr == "" || sellStr == "" {
		return 0, 0, ErrNoDataFound
	}

	buy, err := strconv.ParseFloat(buyStr, 64)
	if err != nil {
		return 0, 0, err
	}

	sell, err := strconv.ParseFloat(sellStr, 64)
	if err != nil {
		return 0, 0, err
	}

	return buy, sell, nil
}

// GetBestObminRate retrieves the USD/UAH exchange rate from BestObmin
func GetBestObminRate() (float64, float64, error) {
	doc, err := FetchHTML("https://bestobmin.com.ua/")
	if err != nil {
		return 0, 0, err
	}
	return parseBestObmin(doc)
}
