package httpx

import (
	"io/ioutil"
	"testing"
)

func TestGetWithUA(t *testing.T) {
	succCases := []string{
		"http://m.facebook.com",
	}
	failCases := []string{
		"NOT_AN_URL",
	}

	for _, url := range succCases {
		_, err := GetWithUA(url, uaiPhone6Plus)
		if err != nil {
			t.Errorf("Failed: %s", url)
		}
	}
	for _, url := range failCases {
		_, err := GetWithUA(url, uaiPhone6Plus)
		if err == nil {
			t.Errorf("Should fail while not: %s", url)
		}
	}
}

func TestGetFullPage(t *testing.T) {
	resp, err := GetFullPage("phantomjs")("http://m.1pondo.tv/movies/1/")
	if err != nil {
		t.Error(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	if len(body) < 1000 {
		t.Errorf("Body too short, could be broken:\n%s\n", string(body))
	}
}
