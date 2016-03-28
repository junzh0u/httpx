package httpx

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/syndtr/goleveldb/leveldb"
)

// Decorator of GetFunc
type Decorator func(GetFunc) GetFunc

// Cached returns a decorator that will cache http responses in leveldb
func Cached(db *leveldb.DB) Decorator {
	return func(get GetFunc) GetFunc {
		return func(url string) (*http.Response, error) {
			data, err := db.Get([]byte(url), nil)
			if err == nil {
				return http.ReadResponse(bufio.NewReader(bytes.NewReader(data)), nil)
			}
			resp, err := get(url)
			if err != nil {
				return resp, err
			}
			if resp.StatusCode != 200 {
				return resp, err
			}
			buffer := new(bytes.Buffer)
			err = resp.Write(buffer)
			if err != nil {
				return resp, err
			}
			cachedresp, err := ioutil.ReadAll(buffer)
			if err != nil {
				return resp, err
			}
			err = db.Put([]byte(url), cachedresp, nil)
			return resp, err
		}
	}
}
