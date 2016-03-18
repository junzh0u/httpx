package httpx

import (
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html/charset"
)

// RespBodyInUTF8 takes a http.Response and returns its body as UTF8 string
func RespBodyInUTF8(resp *http.Response) (string, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	contentType := resp.Header.Get("Content-Type")
	encoding, _, _ := charset.DetermineEncoding(body, contentType)
	utfBody, err := encoding.NewDecoder().Bytes(body)
	if err != nil {
		return "", err
	}
	return string(utfBody), nil
}
