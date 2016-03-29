package httpx

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os/exec"

	"github.com/junzh0u/ioutilx"
)

const (
	uaiPhone6Plus string = "Mozilla/5.0 (iPhone; CPU iPhone OS 8_0 like Mac OS X) AppleWebKit/600.1.3 (KHTML, like Gecko) Version/8.0 Mobile/12A4345d Safari/600.1.4"
	savePageJS    string = `var system = require('system');
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

// GetWithUA takes an UA string, and returns a GetFunc
// which behaves like http.Get, but with specified UA
func GetWithUA(ua string) GetFunc {
	return func(url string) (*http.Response, error) {
		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("User-Agent", ua)
		return client.Do(req)
	}
}

// GetMobile ia a wrapper of GetWithUA with UA of iPhone 6 Plus
func GetMobile(url string) (*http.Response, error) {
	return GetWithUA(uaiPhone6Plus)(url)
}

// GetWithPhantomJS takes a string of js script, and returns a GetFunc
// which behaves like http.Get, but run the script using PhantomJS underneath to
// get the result
func GetWithPhantomJS(script string) GetFunc {
	return func(url string) (*http.Response, error) {
		resp, err := http.Get(url)
		if err != nil {
			return resp, err
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
		resp.Body = ioutil.NopCloser(bytes.NewBuffer(stdout))
		return resp, nil
	}
}

// GetFullPage is a wrapper of GetWithPhantomJS, with default savepage.js
func GetFullPage(url string) (*http.Response, error) {
	return GetWithPhantomJS(savePageJS)(url)
}
