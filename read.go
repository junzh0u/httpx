package httpx

import (
	"io/ioutil"

	"golang.org/x/net/html/charset"
)

// ReadBodyFunc is a function type that takes a url and returns its body in
// string
type ReadBodyFunc func(string) (string, error)

// ReadBodyInUTF8 takes a GetFunc and returns a ReadBodyFunc which calls the
// GetFunc and returns the content of response in UTF8
func ReadBodyInUTF8(getfunc GetFunc) ReadBodyFunc {
	return func(url string) (string, error) {
		resp, err := getfunc(url)
		if err != nil {
			return "", err
		}

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
}
