package config

import "fmt"

// PrintRates – Prints a list of exchange rates
func PrintRates(title string, rates map[string]float64) {
	fmt.Println(title + ":")
	for site, rate := range rates {
		fmt.Printf("%s: %.2f UAH/USD\n", site, rate)
	}
	fmt.Println()
}

// PrintBestRates – Determines and prints the best buy and sell rates
func PrintBestRates(buyRates, sellRates map[string]float64) {
	var bestBuySite string
	var bestBuyRate float64 = -1

	for site, rate := range buyRates {
		if rate > bestBuyRate {
			bestBuyRate = rate
			bestBuySite = site
		}
	}

	var bestSellSite string
	var bestSellRate float64 = -1

	for site, rate := range sellRates {
		if bestSellRate == -1 || rate < bestSellRate {
			bestSellRate = rate
			bestSellSite = site
		}
	}

	if bestBuyRate > 0 {
		fmt.Printf("\n✅ Best buy rate: %.2f UAH/USD at %s\n", bestBuyRate, bestBuySite)
	} else {
		fmt.Println("\n❌ Failed to retrieve any buy rates.")
	}

	if bestSellRate > 0 {
		fmt.Printf("✅ Best sell rate: %.2f UAH/USD at %s\n", bestSellRate, bestSellSite)
	} else {
		fmt.Println("❌ Failed to retrieve any sell rates.")
	}
}
