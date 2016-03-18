package httpx

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os/exec"
)

const (
	uaiPhone6Plus string = "Mozilla/5.0 (iPhone; CPU iPhone OS 8_0 like Mac OS X) AppleWebKit/600.1.3 (KHTML, like Gecko) Version/8.0 Mobile/12A4345d Safari/600.1.4"
)

// GetWithUA is a wrapper of http.Get with specified UA
func GetWithUA(url string, ua string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", ua)
	return client.Do(req)
}

// GetMobile ia a wrapper of GetWithUA with UA of iPhone 6 Plus
func GetMobile(url string) (*http.Response, error) {
	return GetWithUA(url, uaiPhone6Plus)
}

func mkSavePageJS() (string, error) {
	jsfile, err := ioutil.TempFile("/tmp", "phatomjs.savepage")
	if err != nil {
		return "", err
	}
	err = jsfile.Close()
	if err != nil {
		return "", err
	}
	jsScriptPath := jsfile.Name()
	jscontent := `var system = require('system');
	var page = require('webpage').create();

	page.onError = function(msg, trace) {
	// do nothing
	};

	page.open(system.args[1], function(status) {
		console.log(page.content);
		phantom.exit();
	});`
	err = ioutil.WriteFile(jsScriptPath, []byte(jscontent), 0644)
	if err != nil {
		return "", err
	}
	return jsScriptPath, nil
}

// GetFullPage behaves like http.Get, but use PhantomJS underneath to get the full rendered page
func GetFullPage(url string, phantomJsBinPath string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return resp, err
	}
	if resp.StatusCode != 200 {
		return resp, err
	}

	jsScriptPath, err := mkSavePageJS()
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(
		phantomJsBinPath,
		jsScriptPath,
		url,
	)
	stdout, _ := cmd.Output()
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(stdout))
	return resp, nil
}
