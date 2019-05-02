package httpx

import (
	"testing"

	"github.com/junzh0u/ioutilx"
)

func TestDownload(t *testing.T) {
	tmpfile, err := ioutilx.TempFile("", "httpx")
	if err != nil {
		t.Error(err)
	}
	err = Download("https://picsum.photos/200", tmpfile)
	if err != nil {
		t.Error(err)
	}
}
