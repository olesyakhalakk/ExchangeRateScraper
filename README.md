# Exchange Rate Scraper

This project is a Go-based scraper that collects exchange rates from multiple sources and displays the best available rates.

## Features
- Scrapes exchange rates from multiple sources:
    - Minfin
    - BestObmin
    - KytObmin
    - TopObmin
- Finds the best exchange rates for buying and selling USD/UAH.
- Automatically updates every 10 minutes.
- Logs errors if a source is unavailable.

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/olesyakhalakk/price-scraper.git
   cd ExchangeRateScraper
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

3. Build and run the project:
   ```sh
   go run main.go
   ```

## Project Structure

```
ExchangeRateScraper/
│── config/          # Configuration and helper functions
│── scraper/         # Scrapers for different exchange rate sources
│── service/         # Business logic for processing exchange rates
│── main.go          # Entry point of the application
│── go.mod           # Go module file
│── go.sum           # Dependency tracking
```

## Usage
The program continuously fetches exchange rates and prints them to the console.

### Example Output
```
📌 Starting to collect exchange rates...

📊 Exchange Rates:
💵 Buy:
  - Minfin: 38.10
  - BestObmin: 38.20
  - KytObmin: 38.05
  - TopObmin: 38.15

💰 Sell:
  - Minfin: 38.50
  - BestObmin: 38.40
  - KytObmin: 38.45
  - TopObmin: 38.55

✅ Best Buy Rate: BestObmin (38.20)
✅ Best Sell Rate: TopObmin (38.55)

🔄 Next update in 10 minutes...
```

## Contribution
Feel free to contribute by submitting issues or pull requests.

## License
This project is licensed under the MIT License.

