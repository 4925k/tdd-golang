package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"www.google.com",
		"www.facebook.com",
		"www.instagram.com",
	}

	want := map[string]bool{
		"www.google.com":    true,
		"www.facebook.com":  false,
		"www.instagram.com": true,
	}
	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v\n got %v\n", want, got)
	}
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "random url"
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}

func mockWebsiteChecker(url string) bool {
	return url != "www.facebook.com"
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(2 - 0*time.Millisecond)
	return true
}
