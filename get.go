package httpx

import (
	"crypto/tls"
	"net/http"
)

// GetFunc is a function type that takes a url and returns http response,
// similar to http.Get.
type GetFunc func(string) (*http.Response, error)

// GetInsecure is http.Get with InsecureSkipVerify turned on.
func GetInsecure(url string) (*http.Response, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	return client.Get(url)
}

// GetWithCookies returns a GetFunc that sends cookies with request.
func GetWithCookies(cookies []*http.Cookie) GetFunc {
	return func(url string) (*http.Response, error) {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
		for _, cookie := range cookies {
			req.AddCookie(cookie)
		}
		return http.DefaultClient.Do(req)
	}
}
