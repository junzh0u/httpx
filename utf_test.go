package httpx

import (
	"net/http"
	"testing"
)

func TestRespBodyInUTF8(t *testing.T) {
	resp, _ := http.Get("http://www.google.co.jp")
	_, err := RespBodyInUTF8(resp)
	if err != nil {
		t.Error(err)
	}
}
