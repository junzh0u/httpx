package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/benbjohnson/phantomjs"
	"github.com/junzh0u/httpx"
)

func get(url string, getcontent httpx.GetContentFunc) {
	content, err := getcontent(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(content)
}

func main() {
	phantomjs.DefaultProcess.Open()
	defer phantomjs.DefaultProcess.Close()

	flag.Parse()

	switch flag.Arg(0) {
	case "insecure":
		get(flag.Arg(1), httpx.GetContentInUTF8(httpx.GetInsecure))
	case "phantomjs":
		get(flag.Arg(1), httpx.GetContentViaPhantomJS([]*http.Cookie{}, 0, "", ""))
	case "standalone":
		get(flag.Arg(1), httpx.GetContentViaStandalonePhantomJS())
	case "jav":
		javCookie := http.Cookie{
			Name:     "over18",
			Value:    "18",
			Domain:   "www.javlibrary.com",
			Path:     "/",
			HttpOnly: false,
			Secure:   false,
			Expires:  time.Now().Add(1000 * time.Hour),
		}
		get(flag.Arg(1), httpx.GetContentViaPhantomJS([]*http.Cookie{&javCookie}, 6*time.Second, "", ""))
	case "mgs":
		mgsCookie := http.Cookie{
			Name:     "adc",
			Value:    "1",
			Domain:   ".mgstage.com",
			Path:     "/",
			HttpOnly: false,
			Secure:   false,
			Expires:  time.Now().Add(1000 * time.Hour),
		}
		get(flag.Arg(1), httpx.GetContentViaPhantomJS([]*http.Cookie{&mgsCookie}, 0, "", ""))
	default:
		get(flag.Arg(0), httpx.GetContentInUTF8(http.Get))
	}
}
