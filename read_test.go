package httpx

import (
	"net/http"
	"testing"
)

func TestReadBodyX(t *testing.T) {
	_, err := ReadBodyX(http.Get("http://www.google.co.jp"))
	if err != nil {
		t.Error(err)
	}
}
