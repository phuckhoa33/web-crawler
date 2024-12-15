package parser

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type StockData struct {
	Symbol string
	Price  string
	Change string
}

func ParseStockData(html string) []StockData {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Println("Error parsing HTML:", err)
		return nil
	}

	var stocks []StockData
	doc.Find(".stock-row").Each(func(i int, s *goquery.Selection) {
		symbol := s.Find(".symbol").Text()
		price := s.Find(".price").Text()
		change := s.Find(".change").Text()

		stock := StockData{
			Symbol: symbol,
			Price:  price,
			Change: change,
		}
		stocks = append(stocks, stock)
	})

	return stocks
}
