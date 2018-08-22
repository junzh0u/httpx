package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/junzh0u/httpx"
)

func get(url string, getfunc httpx.GetFunc) {
	fmt.Println(url)
	resp, err := getfunc(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Status)
	body, err := httpx.RespBodyInUTF8(resp)
	if err != nil {
		panic(err)
	}
	fmt.Println(body)
}

func main() {
	flag.Parse()

	switch flag.Arg(0) {
	case "insecure":
		get(flag.Arg(1), httpx.GetInsecure)
	default:
		get(flag.Arg(0), http.Get)
	}
}
