package httpx

import (
	"net/http"
	"testing"

	"github.com/benbjohnson/phantomjs"
)

func assertGettable(t *testing.T, getfunc GetFunc, urls []string) {
	for _, url := range urls {
		_, err := getfunc(url)
		if err != nil {
			t.Errorf("Failed: %s", url)
		}
	}
}

func assertNotGettable(t *testing.T, getfunc GetFunc, urls []string) {
	for _, url := range urls {
		_, err := getfunc(url)
		if err == nil {
			t.Errorf("Should fail while not: %s", url)
		}
	}
}

func TestGetInsecure(t *testing.T) {
	assertGettable(t, GetInsecure, []string{
		"https://www.tokyo-hot.com/product/?q=n0110",
	})
	assertNotGettable(t, GetInsecure, []string{
		"NOT_AN_URL",
	})
}

func BenchmarkGetViaPhantomJS(b *testing.B) {
	phantomjs.DefaultProcess.Open()
	defer phantomjs.DefaultProcess.Close()

	for i := 0; i < b.N; i++ {
		GetViaPhantomJS([]*http.Cookie{}, 0)("http://m.1pondo.tv/movies/1/")
	}
}

func BenchmarkGetViaStandalonePhantomJS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetViaStandalonePhantomJS()("http://m.1pondo.tv/movies/1/")
	}
}
