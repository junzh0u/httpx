package httpx

import (
	"net/http"
	"testing"
)

func TestReadBodyInUTF8(t *testing.T) {
	_, err := ReadBodyInUTF8(http.Get)("http://www.google.co.jp")
	if err != nil {
		t.Error(err)
	}
}
