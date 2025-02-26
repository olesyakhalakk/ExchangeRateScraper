package scraper

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// GetKytObminRate отримує курси купівлі та продажу USD/UAH з KytObmin
func GetKytObminRate() (buyRate, sellRate float64, err error) {
	url := "https://kyt-obmin.if.ua"

	// Виконуємо HTTP-запит
	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, fmt.Errorf("помилка отримання сторінки: %v", err)
	}
	defer resp.Body.Close()

	// Перевіряємо статус відповіді
	if resp.StatusCode != 200 {
		return 0, 0, fmt.Errorf("неочікуваний статус відповіді: %d %s", resp.StatusCode, resp.Status)
	}

	// Парсимо HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return 0, 0, fmt.Errorf("помилка парсингу HTML: %v", err)
	}

	// Шукаємо курс купівлі та продажу USD/UAH
	var rates []float64
	doc.Find("a[href*='usd-uah']").Each(func(i int, s *goquery.Selection) {
		rateStr := strings.TrimSpace(s.Text())
		rateStr = strings.ReplaceAll(rateStr, ",", ".")

		// Конвертуємо у float64
		rate, err := strconv.ParseFloat(rateStr, 64)
		if err == nil {
			rates = append(rates, rate)
		}
	})

	if len(rates) < 2 {
		return 0, 0, fmt.Errorf("не вдалося знайти курс USD/UAH")
	}

	return rates[0], rates[1], nil
}
