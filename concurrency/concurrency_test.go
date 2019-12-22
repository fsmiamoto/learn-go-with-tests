package concurrency

import (
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "http://omaehamoushindeiru.com" {
		return false
	}

	return true
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsite(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a random url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}

func TestCheckWebsite(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gyspydave5.com",
		"http://github.com",
		"http://omaehamoushindeiru.com",
	}

	want := map[string]bool{
		"http://google.com":             true,
		"http://blog.gyspydave5.com":    true,
		"http://github.com":             true,
		"http://omaehamoushindeiru.com": false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	for _, url := range websites {
		if want[url] != got[url] {
			t.Errorf("Wanted %v but got %v for %v", want[url], got[url], url)
		}
	}
}
