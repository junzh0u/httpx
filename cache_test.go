package httpx

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/syndtr/goleveldb/leveldb"
)

func TestCached(t *testing.T) {
	tmpdir, err := ioutil.TempDir("", "httpx.test.")
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(tmpdir)
	db, err := leveldb.OpenFile(tmpdir, nil)
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	cachedGet := Cached(db)(http.Get)
	for i := 0; i < 10; i++ {
		_, err := cachedGet("http://www.google.com")
		if err != nil {
			t.Error(err)
		}
	}
}

func BenchmarkCached(b *testing.B) {
	tmpdir, err := ioutil.TempDir("", "httpx.test.")
	if err != nil {
		b.Error(err)
	}
	defer os.RemoveAll(tmpdir)
	db, err := leveldb.OpenFile(tmpdir, nil)
	if err != nil {
		b.Error(err)
	}
	defer db.Close()
	cachedGet := Cached(db)(http.Get)
	for i := 0; i < b.N; i++ {
		_, err := cachedGet("http://www.google.com")
		if err != nil {
			b.Error(err)
		}
	}
}
