package httpx

import (
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/benbjohnson/phantomjs"
)

func TestGetContentInUTF8(t *testing.T) {
	_, err := GetContentInUTF8(http.Get)("http://www.google.co.jp")
	if err != nil {
		t.Error(err)
	}
}

func TestGetContentViaPhantomJSWithCookie(t *testing.T) {
	phantomjs.DefaultProcess.Open()
	defer phantomjs.DefaultProcess.Close()

	mgsCookie := http.Cookie{
		Name:     "adc",
		Value:    "1",
		Domain:   ".mgstage.com",
		Path:     "/",
		HttpOnly: false,
		Secure:   false,
		Expires:  time.Now().Add(1000 * time.Hour),
	}
	content, err := GetContentViaPhantomJS([]*http.Cookie{&mgsCookie}, 0, "", "")(
		"http://www.mgstage.com/search/search.php?search_word=SIRO-1715&search_shop_id=shiroutotv")
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(content, "年齢認証") {
		t.Fatal("Age verification required, cookie not working")
	}
	if !strings.Contains(content, "video/SIRO-1715") {
		t.Fatal("No link found, page might not fully load")
	}
}

func TestGetContentViaPhantomJSWithWait(t *testing.T) {
	phantomjs.DefaultProcess.Open()
	defer phantomjs.DefaultProcess.Close()

	javCookie := http.Cookie{
		Name:     "over18",
		Value:    "18",
		Domain:   "www.javlibrary.com",
		Path:     "/",
		HttpOnly: false,
		Secure:   false,
		Expires:  time.Now().Add(1000 * time.Hour),
	}
	content, err := GetContentViaPhantomJS([]*http.Cookie{&javCookie}, 6*time.Second, "", "")(
		"http://www.javlibrary.com/ja/?v=javlimtiza")
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(content, "cf-im-under-attack") {
		t.Fatal("Challenged by Cloudfare")
	}
	if !strings.Contains(content, "SDDE-222") {
		t.Fatal("Page not fully loaded")
	}
}

func TestGetContentViaStandalonePhantomJS(t *testing.T) {
	content, err := GetContentViaStandalonePhantomJS()("http://m.1pondo.tv/movies/1/")
	if err != nil {
		t.Error(err)
	}
	if len(content) < 1000 {
		t.Error("Content too short, could be broken")
	}
}

func BenchmarkGetContentViaPhantomJS(b *testing.B) {
	phantomjs.DefaultProcess.Open()
	defer phantomjs.DefaultProcess.Close()

	for i := 0; i < b.N; i++ {
		GetContentViaPhantomJS([]*http.Cookie{}, 0, "", "")("http://m.1pondo.tv/movies/1/")
	}
}

func BenchmarkGetContentFullPage(b *testing.B) {
	phantomjs.DefaultProcess.Open()
	defer phantomjs.DefaultProcess.Close()

	for i := 0; i < b.N; i++ {
		GetContentFullPage("http://m.1pondo.tv/movies/1/")
	}
}

func BenchmarkGetContentViaStandalonePhantomJS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetContentViaStandalonePhantomJS()("http://m.1pondo.tv/movies/1/")
	}
}
