package scraper

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
)

// ExchangeRate represents the structure of the currency exchange rate
type ExchangeRate struct {
	Type        string `json:"@type"`
	Currency    string `json:"currency"`
	Description string `json:"description"`
	Rate        struct {
		Price string `json:"price"`
	} `json:"currentExchangeRate"`
}

// parseMinfin parses the HTML page and extracts the exchange rate
func parseMinfin(doc *goquery.Document) (float64, float64, error) {
	var jsonData string

	// Locate the JSON containing exchange rates
	doc.Find("script[type='application/ld+json']").EachWithBreak(func(i int, s *goquery.Selection) bool {
		content := strings.TrimSpace(s.Text())
		if strings.Contains(content, "\"@type\":\"ItemList\"") {
			jsonData = content
			return false // Stop iteration once the required JSON is found
		}
		return true
	})

	if jsonData == "" {
		return 0, 0, fmt.Errorf("required JSON not found in script")
	}

	// Structure for JSON parsing
	var data struct {
		MainEntity struct {
			ItemListElement []struct {
				Type                string `json:"@type"`
				Currency            string `json:"currency"`
				Name                string `json:"name"`
				Description         string `json:"description"`
				CurrentExchangeRate struct {
					Price         string `json:"price"`
					PriceCurrency string `json:"priceCurrency"`
				} `json:"currentExchangeRate"`
			} `json:"itemListElement"`
		} `json:"mainEntity"`
	}

	// Parse JSON data
	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		return 0, 0, fmt.Errorf("JSON parsing error: %v", err)
	}

	var buy, sell float64

	for _, item := range data.MainEntity.ItemListElement {
		if item.Type != "ExchangeRateSpecification" || item.Currency != "USD" {
			continue
		}

		priceStr := strings.TrimSpace(item.CurrentExchangeRate.Price)
		if priceStr == "" || priceStr == "-" {
			continue
		}

		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			continue
		}

		switch item.Description {
		case "Курс купівлі":
			buy = price
		case "Курс продажу":
			sell = price
		}
	}

	if buy == 0 || sell == 0 {
		return 0, 0, fmt.Errorf("USD exchange rate not found")
	}

	return buy, sell, nil
}

// GetMinfinRate retrieves the USD/UAH exchange rate from Minfin via Scraper
func GetMinfinRate() (float64, float64, error) {
	s := Scraper{
		URL:       "https://minfin.com.ua/ua/currency/ivano-frankovsk/",
		UserAgent: "Mozilla/5.0",
		ParseFunc: parseMinfin,
	}
	return s.FetchData()
}
