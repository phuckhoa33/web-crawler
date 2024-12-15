package main

import "github.com/phuckhoa33/web-crawler/internal/crawler"

func main() {
	startURLs := []string{
		"https://finance.yahoo.com",
		// Thêm các URL của trang web chứng khoán khác
	}

	maxWorkers := 10
	c := crawler.NewCrawler(startURLs, maxWorkers)
	c.Run()

	// Xử lý dữ liệu ở đây (ví dụ: lưu vào DB)
}
