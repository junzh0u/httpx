package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/junzh0u/httpx"
)

func exec(url string, readbody httpx.ReadBodyFunc) {
	body, err := readbody(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(body)
}

func main() {
	flag.Parse()

	switch flag.Arg(0) {
	case "insecure":
		exec(flag.Arg(1), httpx.ReadBodyInUTF8(httpx.GetInsecure))
	default:
		exec(flag.Arg(0), httpx.ReadBodyInUTF8(http.Get))
	}
}
