package scraper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ExchangeRate описує структуру курсу валют
type ExchangeRate struct {
	Type        string `json:"@type"`
	Currency    string `json:"currency"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Rate        struct {
		Price string `json:"price"`
	} `json:"currentExchangeRate"`
}

// GetMinfinRate отримує курс купівлі та продажу для вказаної валюти
func GetMinfinRate(currency string) (buyRate float64, sellRate float64, err error) {
	url := "https://minfin.com.ua/ua/currency/ivano-frankovsk/"

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return 0, 0, fmt.Errorf("помилка запиту: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return 0, 0, fmt.Errorf("отримано статус-код %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return 0, 0, fmt.Errorf("помилка розбору HTML: %v", err)
	}

	var jsonData string
	doc.Find("script[type='application/ld+json']").Each(func(i int, s *goquery.Selection) {
		content := strings.TrimSpace(s.Text())
		if strings.Contains(content, "\"@type\":\"ItemList\"") {
			jsonData = content
		}
	})

	if jsonData == "" {
		return 0, 0, fmt.Errorf("не знайдено потрібного JSON у script")
	}

	var data struct {
		MainEntity struct {
			ItemListElement []ExchangeRate `json:"itemListElement"`
		} `json:"mainEntity"`
	}

	err = json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		return 0, 0, fmt.Errorf("помилка розбору JSON: %v", err)
	}

	var buy, sell float64

	for _, item := range data.MainEntity.ItemListElement {
		if item.Type == "ExchangeRateSpecification" && item.Currency == currency {
			priceStr := strings.TrimSpace(item.Rate.Price)
			if priceStr == "-" || priceStr == "" {
				continue // Пропускаємо відсутні значення
			}
			price, err := strconv.ParseFloat(priceStr, 64)
			if err != nil {
				return 0, 0, fmt.Errorf("помилка конвертації price: %v", err)
			}

			if item.Description == "Курс купівлі" {
				buy = price
			} else if item.Description == "Курс продажу" {
				sell = price
			}
		}
	}

	if buy == 0 || sell == 0 {
		return 0, 0, fmt.Errorf("не вдалося знайти курс %s", currency)
	}

	return buy, sell, nil
}
