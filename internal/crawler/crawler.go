package crawler

import (
	"fmt"
	"sync"

	"github.com/phuckhoa33/web-crawler/internal/fetcher"
	parser "github.com/phuckhoa33/web-crawler/internal/parsers"
)

type Crawler struct {
	urls       []string
	maxWorkers int
	wg         sync.WaitGroup
}

func NewCrawler(urls []string, maxWorkers int) *Crawler {
	return &Crawler{
		urls:       urls,
		maxWorkers: maxWorkers,
	}
}

func (c *Crawler) Run() {
	urlQueue := make(chan string, 100)

	// Thêm tất cả URL vào hàng đợi
	for _, url := range c.urls {
		urlQueue <- url
	}

	// Khởi tạo các worker để xử lý song song
	for i := 0; i < c.maxWorkers; i++ {
		c.wg.Add(1)
		go c.worker(urlQueue)
	}

	close(urlQueue)
	c.wg.Wait()
}

func (c *Crawler) worker(urlQueue chan string) {
	defer c.wg.Done()

	for url := range urlQueue {
		body, err := fetcher.Fetch(url)
		fmt.Println(url)
		if err != nil {
			fmt.Printf("Error fetching %s: %v\n", url, err)
			continue
		}

		// Parse dữ liệu chứng khoán từ HTML
		data := parser.ParseStockData(body)
		fmt.Printf("Crawled data from %s: %+v\n", url, data)

		// Ở đây, bạn có thể lưu trữ dữ liệu vào DB hoặc file
	}
}
