package httpx

import (
	"net/http"
	"testing"

	"github.com/syndtr/goleveldb/leveldb"
)

func TestCached(t *testing.T) {
	db, err := leveldb.OpenFile("/tmp/httpx.test.db", nil)
	defer db.Close()
	if err != nil {
		t.Error(err)
	}
	cachedGet := Cached(db)(http.Get)
	for i := 0; i < 10; i++ {
		_, err := cachedGet("http://www.google.com")
		if err != nil {
			t.Error(err)
		}
	}
}

func BenchmarkCached(b *testing.B) {
	db, err := leveldb.OpenFile("/tmp/httpx.test.db", nil)
	defer db.Close()
	if err != nil {
		b.Error(err)
	}
	cachedGet := Cached(db)(http.Get)
	for i := 0; i < b.N; i++ {
		_, err := cachedGet("http://www.google.com")
		if err != nil {
			b.Error(err)
		}
	}
}
