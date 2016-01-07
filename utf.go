package httpx

import (
  "bytes"
  "io/ioutil"
  "net/http"

  "github.com/PuerkitoBio/goquery"
  "golang.org/x/net/html/charset"
)

func RespBodyInUTF8(resp *http.Response) (string, error) {
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return "", err
  }
  rawdoc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
  if err != nil {
    return "", err
  }
  contenttype, _ := rawdoc.Find("meta[http-equiv=content-type]").Attr("content")
  encoding, _, _ := charset.DetermineEncoding(body, contenttype)
  utfbody, err := encoding.NewDecoder().Bytes(body)
  if err != nil {
    return "", err
  }
  return string(utfbody), nil
}
