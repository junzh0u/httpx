package utfhttp

import (
  "bytes"
  "fmt"
  "io/ioutil"
  "net/http"
  "strings"

  "github.com/golang/glog"
  "github.com/PuerkitoBio/goquery"
  "golang.org/x/net/html/charset"
)

func GetBody(url string) (string, error) {
  resp, err := http.Get(url)
  if err != nil {
    return "", err
  }
  if resp.StatusCode != http.StatusOK {
    return "", fmt.Errorf("Unexpected status code %d from %s", resp.StatusCode, url)
  }
  body, err := ioutil.ReadAll(resp.Body)
  resp.Body.Close()
  if err != nil {
    return "", err
  }
  rawdoc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
  if err != nil {
    return "", err
  }
  contenttype, _ := rawdoc.Find("meta[http-equiv=content-type]").Attr("content")
  encoding, name, certain := charset.DetermineEncoding(body, contenttype)
  if certain {
    glog.V(1).Infof("[UTFHTTP] Encoding of %v is %v", url, name)
  } else {
    glog.V(1).Infof("[UTFHTTP] Guess encoding of %v is %v", url, name)
  }
  utfbody, err := encoding.NewDecoder().Bytes(body)
  if err != nil {
    return "", err
  }
  return string(utfbody), nil
}

func GetDocument(url string) (*goquery.Document, error) {
  body, err := GetBody(url)
  if err != nil {
    return nil, err
  }
  return goquery.NewDocumentFromReader(strings.NewReader(body))
}
