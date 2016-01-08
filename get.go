package httpx

import (
  "net/http"
)

const (
  UA_iPhone_6_Plus string = "Mozilla/5.0 (iPhone; CPU iPhone OS 8_0 like Mac OS X) AppleWebKit/600.1.3 (KHTML, like Gecko) Version/8.0 Mobile/12A4345d Safari/600.1.4"
)

func GetWithUA(url string, ua string) (*http.Response, error) {
  client := &http.Client{}
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    return nil, err
  }
  req.Header.Set("User-Agent", ua)
  return client.Do(req)
}

func GetMobile(url string) (*http.Response, error) {
  return GetWithUA(url, UA_iPhone_6_Plus)
}
