package httpx

import (
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestGetInsecure(t *testing.T) {
	_, err := GetInsecure("http://www.google.co.jp")
	if err != nil {
		t.Error(err)
	}
}

func TestGetWithCookies(t *testing.T) {
	mgsCookies := []*http.Cookie{&http.Cookie{
		Name:     "adc",
		Value:    "1",
		Domain:   ".mgstage.com",
		Path:     "/",
		HttpOnly: false,
		Secure:   false,
		Expires:  time.Now().Add(1000 * time.Hour),
	}}

	content, err := ReadBodyInUTF8(GetWithCookies(mgsCookies))(
		"https://www.mgstage.com/search/search.php?search_word=SIRO-1715&search_shop_id=shiroutotv")
	if err != nil {
		t.Error(err)
	}
	if strings.Contains(content, "年齢認証") {
		t.Errorf("Challenged by 年齢認証")
	}
}
