package concurrency

type WebsiteChecker func(string) bool

type result struct {
	url string
	ok  bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	c := make(chan result)

	for _, url := range urls {
		go func(u string) {
			c <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-c
		results[r.url] = r.ok
	}

	return results
}
