package service

import (
	"log"
)

// ExchangeSource represents an exchange rate source
type ExchangeSource struct {
	Name  string
	Fetch func() (float64, float64, error)
}

// CollectExchangeRates gathers exchange rates from various sources
func CollectExchangeRates(sources []ExchangeSource) (map[string]float64, map[string]float64) {
	buyRates := make(map[string]float64)
	sellRates := make(map[string]float64)

	for _, source := range sources {
		buy, sell, err := source.Fetch()
		if err == nil {
			buyRates[source.Name] = buy
			sellRates[source.Name] = sell
		} else {
			log.Println("❌ Parsing error", source.Name, ":", err)
		}
	}

	return buyRates, sellRates
}
