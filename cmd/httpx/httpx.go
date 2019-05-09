package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/junzh0u/httpx"
)

func main() {
	flag.Parse()
	body, err := httpx.ReadBodyX(http.Get(flag.Arg(0)))
	if err != nil {
		panic(err)
	}
	fmt.Println(body)
}
