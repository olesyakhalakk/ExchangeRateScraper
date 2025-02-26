package scraper

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strconv"
	"strings"
)

// GetTopObminRate парсить сайт TopObmin і повертає курс купівлі та продажу USD/UAH
func GetTopObminRate() (buyRate, sellRate float64, err error) {
	url := "https://topobmin.com/"

	// Додаємо User-Agent
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return 0, 0, fmt.Errorf("помилка завантаження сторінки: %v", err)
	}
	defer resp.Body.Close()

	tokenizer := html.NewTokenizer(resp.Body)
	var rates []float64

	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			if len(rates) < 2 {
				return 0, 0, fmt.Errorf("курс не знайдено")
			}
			return rates[0], rates[1], nil
		case html.StartTagToken:
			token := tokenizer.Token()

			// Шукаємо <td class="green"> (курс купівлі) або <td class="red"> (курс продажу)
			if token.Data == "td" {
				for _, attr := range token.Attr {
					if attr.Key == "class" && (attr.Val == "green" || attr.Val == "red") {
						tokenizer.Next()
						rateStr := strings.TrimSpace(tokenizer.Token().Data)
						rateStr = strings.ReplaceAll(rateStr, ",", ".")

						// Конвертуємо в float64
						rate, err := strconv.ParseFloat(rateStr, 64)
						if err == nil {
							rates = append(rates, rate)
						}
					}
				}
			}
		}
	}
}
