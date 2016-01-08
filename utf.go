package httpx

import (
  "io/ioutil"
  "net/http"

  "golang.org/x/net/html/charset"
)

func RespBodyInUTF8(resp *http.Response) (string, error) {
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return "", err
  }
  contenttype := resp.Header.Get("Content-Type")
  encoding, _, _ := charset.DetermineEncoding(body, contenttype)
  utfbody, err := encoding.NewDecoder().Bytes(body)
  if err != nil {
    return "", err
  }
  return string(utfbody), nil
}