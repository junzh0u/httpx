package httpx

import (
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html/charset"
)

// ReadBody takes a http.Response and returns its body in UTF8 string.
func ReadBody(resp *http.Response) (string, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	contentType := resp.Header.Get("Content-Type")
	encoding, _, _ := charset.DetermineEncoding(body, contentType)
	content, err := encoding.NewDecoder().Bytes(body)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// ReadBodyX works just like ReadBody, except it takes an extra error param,
// so that it can be chained with http.Get or its alternatives.
func ReadBodyX(resp *http.Response, err error) (string, error) {
	if err != nil {
		return "", err
	}
	return ReadBody(resp)
}
