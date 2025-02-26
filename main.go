package main

import (
	"fmt"
	"log"
	"price-scraper/scraper"
	"time"
)

func main() {
	for {
		fmt.Println("📌 Початок збору курсів валют...")

		// Карти для зберігання курсів купівлі та продажу
		buyRates := make(map[string]float64)
		sellRates := make(map[string]float64)

		// Отримуємо курси з Minfin
		usdMinfinBuy, usdMinfinSell, err := scraper.GetMinfinRate("USD")
		if err == nil {
			buyRates["Minfin"] = usdMinfinBuy
			sellRates["Minfin"] = usdMinfinSell
		} else {
			log.Println("❌ Помилка парсингу Minfin:", err)
		}

		// Отримуємо курси з BestObmin
		usdBestObminBuy, usdBestObminSell, err := scraper.GetBestObminRate()
		if err == nil {
			buyRates["BestObmin"] = usdBestObminBuy
			sellRates["BestObmin"] = usdBestObminSell
		} else {
			log.Println("❌ Помилка парсингу BestObmin:", err)
		}

		// Отримуємо курси з KytObmin
		usdKytObminBuy, usdKytObminSell, err := scraper.GetKytObminRate()
		if err == nil {
			buyRates["KytObmin"] = usdKytObminBuy
			sellRates["KytObmin"] = usdKytObminSell
		} else {
			log.Println("❌ Помилка парсингу Kyt-Obmin:", err)
		}

		// Отримуємо курси з TopObmin
		usdTopObminBuy, usdTopObminSell, err := scraper.GetTopObminRate()
		if err == nil {
			buyRates["TopObmin"] = usdTopObminBuy
			sellRates["TopObmin"] = usdTopObminSell
		} else {
			log.Println("❌ Помилка парсингу TopObmin:", err)
		}

		// Виводимо результати
		fmt.Println("\n📊 Курси валют:")
		fmt.Println("💵 Купівля:")
		for site, rate := range buyRates {
			fmt.Printf("%s: %.2f UAH/USD\n", site, rate)
		}

		fmt.Println("\n💰 Продаж:")
		for site, rate := range sellRates {
			fmt.Printf("%s: %.2f UAH/USD\n", site, rate)
		}

		// Знаходимо найкращий курс для купівлі (найбільший)
		var bestBuySite string
		var bestBuyRate float64 = -1 // Встановлюємо мінімальне значення для пошуку максимального

		for site, rate := range buyRates {
			if rate > bestBuyRate {
				bestBuyRate = rate
				bestBuySite = site
			}
		}

		// Знаходимо найкращий курс для продажу (найменший)
		var bestSellSite string
		var bestSellRate float64 = -1 // Встановлюємо мінімальне значення для пошуку мінімального

		for site, rate := range sellRates {
			if bestSellRate == -1 || rate < bestSellRate {
				bestSellRate = rate
				bestSellSite = site
			}
		}

		// Виводимо найкращі курси
		if bestBuyRate > 0 {
			fmt.Printf("\n✅ Найкращий курс для купівлі: %.2f UAH/USD на %s\n", bestBuyRate, bestBuySite)
		} else {
			fmt.Println("\n❌ Не вдалося отримати жодного курсу для купівлі.")
		}

		if bestSellRate > 0 {
			fmt.Printf("✅ Найкращий курс для продажу: %.2f UAH/USD на %s\n", bestSellRate, bestSellSite)
		} else {
			fmt.Println("❌ Не вдалося отримати жодного курсу для продажу.")
		}

		fmt.Println("\n🔄 Наступне оновлення через 10 хвилин...")
		time.Sleep(10 * time.Minute)
	}
}
