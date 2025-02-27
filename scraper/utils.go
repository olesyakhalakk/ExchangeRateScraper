package scraper

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var ErrNoDataFound = errors.New("required data not found")

// FetchHTML retrieves an HTML page and returns a goquery.Document
func FetchHTML(url string) (*goquery.Document, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error fetching page: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received status code %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error parsing HTML: %v", err)
	}

	return doc, nil
}

// ExtractJSON finds the required JSON in <script type='application/ld+json'>
func ExtractJSON(doc *goquery.Document) (string, error) {
	var jsonData string
	doc.Find("script[type='application/ld+json']").EachWithBreak(func(i int, s *goquery.Selection) bool {
		content := s.Text()
		if content != "" && jsonData == "" {
			jsonData = content
			return false // Stop iteration once the required JSON is found
		}
		return true
	})

	if jsonData == "" {
		return "", fmt.Errorf("required JSON not found in script")
	}

	return jsonData, nil
}

// CleanPriceString cleans and converts a string representation of a number into float64
func CleanPriceString(priceStr string) (float64, error) {
	priceStr = strings.TrimSpace(strings.ReplaceAll(priceStr, ",", "."))
	if priceStr == "" || priceStr == "-" {
		return 0, ErrNoDataFound
	}

	return strconv.ParseFloat(priceStr, 64)
}
