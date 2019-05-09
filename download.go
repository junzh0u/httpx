package httpx

import (
	"io/ioutil"
	"net/http"
	"os"
)

// Download gets the remote srcURL then writes to a local file under destPath
func Download(srcURL string, destPath string) error {
	resp, err := http.Get(srcURL)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	ioutil.WriteFile(destPath, data, os.ModePerm)
	return nil
}
