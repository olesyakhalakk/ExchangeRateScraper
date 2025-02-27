package main

import (
	"fmt"
	"price-scraper/config"
	"price-scraper/scraper"
	"price-scraper/service"
	"time"
)

func main() {
	for {
		fmt.Println("📌 Starting to collect exchange rates...")

		buyRates, sellRates := service.CollectExchangeRates([]service.ExchangeSource{
			{Name: "Minfin", Fetch: scraper.GetMinfinRate},
			{Name: "BestObmin", Fetch: scraper.GetBestObminRate},
			{Name: "KytObmin", Fetch: scraper.GetKytObminRate},
			{Name: "TopObmin", Fetch: scraper.GetTopObminRate},
		})

		// Display results
		fmt.Println("\n📊 Exchange Rates:")
		config.PrintRates("💵 Buy", buyRates)
		config.PrintRates("💰 Sell", sellRates)

		// Find the best exchange rates
		config.PrintBestRates(buyRates, sellRates)

		fmt.Println("\n🔄 Next update in 10 minutes...")
		time.Sleep(10 * time.Minute)
	}
}
