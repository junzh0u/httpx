package httpx

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Download gets the remote srcURL then writes to a local file under destPath
func Download(srcURL string, destPath string) error {
	fmt.Printf("\tDownloading %s to %s\n", srcURL, destPath)
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
