package httpx

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"time"
)

// GetFunc is a function type that takes a url and returns http response,
// similar as http.Get
type GetFunc func(string) (*http.Response, error)

// GetInsecure is http.Get with InsecureSkipVerify turned on
func GetInsecure(url string) (*http.Response, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	return client.Get(url)
}

func headThenReplaceContent(getContent GetContentFunc) GetFunc {
	return func(url string) (*http.Response, error) {
		resp, err := http.Head(url)
		if err != nil {
			return resp, err
		}

		content, err := getContent(url)
		if err != nil {
			return nil, err
		}

		resp.ContentLength = int64(len(content))
		resp.Body = ioutil.NopCloser(bytes.NewBufferString(content))
		return resp, nil
	}
}

// GetViaPhantomJS returns a GetFunc that calls http.Head to get header, while
// using PhantomJS wrapper to get content
func GetViaPhantomJS(cookies []*http.Cookie, waitDuration time.Duration) GetFunc {
	return headThenReplaceContent(GetContentViaPhantomJS(cookies, waitDuration))
}

// GetViaStandalonePhantomJS returns a GetFunc that calls http.Head to get
// header, while using a standalone PhantomJS process to get content
func GetViaStandalonePhantomJS() GetFunc {
	return headThenReplaceContent(GetContentViaStandalonePhantomJS())
}
