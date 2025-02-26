package scraper

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// GetBestObminRate отримує курс купівлі та продажу USD/UAH з BestObmin
func GetBestObminRate() (buyRate, sellRate float64, err error) {
	url := "https://bestobmin.com.ua/"

	// Завантажуємо сторінку
	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, fmt.Errorf("не вдалося отримати сторінку: %v", err)
	}
	defer resp.Body.Close()

	// Перевіряємо статус відповіді
	if resp.StatusCode != http.StatusOK {
		return 0, 0, fmt.Errorf("неочікуваний статус відповіді: %d", resp.StatusCode)
	}

	// Парсимо HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return 0, 0, fmt.Errorf("не вдалося розпарсити HTML: %v", err)
	}

	// Шукаємо всі блоки з курсами для купівлі
	buyRateElement := doc.Find("div.digit_bg.left_digit_bg p")
	if buyRateElement.Length() == 0 {
		return 0, 0, fmt.Errorf("не знайдено курс купівлі")
	}

	// Отримуємо курс купівлі
	buyRateStr := strings.TrimSpace(buyRateElement.Eq(0).Text())
	buyRate, err = strconv.ParseFloat(buyRateStr, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("помилка парсингу курсу купівлі: %v", err)
	}

	// Шукаємо всі блоки з курсами для продажу
	sellRateElement := doc.Find("div.digit_bg.right_digit_bg p")
	if sellRateElement.Length() == 0 {
		return 0, 0, fmt.Errorf("не знайдено курс продажу")
	}

	// Отримуємо курс продажу
	sellRateStr := strings.TrimSpace(sellRateElement.Eq(0).Text())
	sellRate, err = strconv.ParseFloat(sellRateStr, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("помилка парсингу курсу продажу: %v", err)
	}

	return buyRate, sellRate, nil
}
