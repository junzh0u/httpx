package httpx

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"os/exec"

	"github.com/junzh0u/ioutilx"
)

const (
	savePageJS string = `var system = require('system');
var page = require('webpage').create();

page.onError = function(msg, trace) {
	// do nothing
};

page.open(system.args[1], function(status) {
  console.log(page.content);
  phantom.exit();
});`
)

// GetFunc is a function type that takes a url and returns http response,
// similar as http.Get
type GetFunc func(string) (*http.Response, error)

var phantomPoolSize = 20
var phantomPool = make(chan int, phantomPoolSize)

func init() {
	for i := 1; i <= phantomPoolSize; i++ {
		phantomPool <- 1
	}
}

// GetWithPhantomJS takes a string of js script, and returns a GetFunc
// which behaves like http.Get, but run the script using PhantomJS underneath to
// get the result
func GetWithPhantomJS(script string, alwaysOk bool) GetFunc {
	return func(url string) (*http.Response, error) {
		<-phantomPool
		defer func() { phantomPool <- 1 }()

		resp, err := http.Get(url)
		if err != nil {
			return resp, err
		}
		if alwaysOk {
			resp.StatusCode = 200
		}
		if resp.StatusCode != 200 {
			return resp, err
		}

		scriptPath, err := ioutilx.TempFileWithContent(
			"/tmp", "phatomjs.savepage", script)
		if err != nil {
			return nil, err
		}
		cmd := exec.Command(
			"phantomjs",
			scriptPath,
			url,
		)
		stdout, err := cmd.Output()
		if err != nil {
			return nil, err
		}
		resp.ContentLength = int64(len(stdout))
		resp.Body = ioutil.NopCloser(bytes.NewBuffer(stdout))
		return resp, nil
	}
}

// GetFullPage is a wrapper of GetWithPhantomJS, with default savepage.js
func GetFullPage(url string) (*http.Response, error) {
	return GetWithPhantomJS(savePageJS, false)(url)
}

// GetInsecure is a http.Get with InsecureSkipVerify turned on
func GetInsecure(url string) (*http.Response, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	return client.Get(url)
}
