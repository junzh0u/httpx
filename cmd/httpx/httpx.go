package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/benbjohnson/phantomjs"
	"github.com/junzh0u/httpx"
)

func get(url string, getbody httpx.ReadBodyFunc) {
	body, err := getbody(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(body)
}

func main() {
	phantomjs.DefaultProcess.Open()
	defer phantomjs.DefaultProcess.Close()

	flag.Parse()

	switch flag.Arg(0) {
	case "insecure":
		get(flag.Arg(1), httpx.ReadBodyInUTF8(httpx.GetInsecure))
	default:
		get(flag.Arg(0), httpx.ReadBodyInUTF8(http.Get))
	}
}
