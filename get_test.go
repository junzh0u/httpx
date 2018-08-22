package httpx

import (
	"io/ioutil"
	"testing"
)

func assertGettable(t *testing.T, getfunc GetFunc, urls []string) {
	for _, url := range urls {
		_, err := getfunc(url)
		if err != nil {
			t.Errorf("Failed: %s", url)
		}
	}
}

func assertNotGettable(t *testing.T, getfunc GetFunc, urls []string) {
	for _, url := range urls {
		_, err := getfunc(url)
		if err == nil {
			t.Errorf("Should fail while not: %s", url)
		}
	}
}

func TestGetFullPage(t *testing.T) {
	resp, err := GetFullPage("http://m.1pondo.tv/movies/1/")
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

func TestGetInsecure(t *testing.T) {
	assertGettable(t, GetInsecure, []string{
		"https://www.tokyo-hot.com/product/?q=n0110",
	})
	assertNotGettable(t, GetInsecure, []string{
		"NOT_AN_URL",
	})
}
