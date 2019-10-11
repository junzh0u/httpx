package httpx

import (
	"net/http"
	"testing"
)

func TestReadBody(t *testing.T) {
	resp, err := http.Get("http://www.google.co.jp")
	if err != nil {
		t.Error(err)
	}
	firstpass, err := ReadBody(resp)
	if err != nil {
		t.Error(err)
	}
	secondpass, err := ReadBody(resp)
	if err != nil {
		t.Error(err)
	}
	if firstpass != secondpass {
		t.Error("First pass differs from second pass")
	}
}

func TestReadBodyX(t *testing.T) {
	_, err := ReadBodyX(http.Get("http://www.google.co.jp"))
	if err != nil {
		t.Error(err)
	}
}
