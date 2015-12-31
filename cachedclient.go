package utfhttp

import (
  "encoding/json"
  "io/ioutil"
  "os"
  "strings"
  "sync"

  "github.com/golang/glog"
  "github.com/PuerkitoBio/goquery"
)

type CachedClient struct {
  sync.Mutex
  Cache map[string]string
  DbPath string
}

func NewCachedClient(dbpath string) (*CachedClient, error) {
  client := new(CachedClient)
  client.Cache = make(map[string]string)
  client.DbPath = dbpath
  file, err := ioutil.ReadFile(dbpath)
  if err != nil {
    return client, err
  }
  err = json.Unmarshal(file, &client)
  return client, err
}

func (cc *CachedClient) Close() {
  data, err := json.Marshal(cc)
  if err == nil {
    err = ioutil.WriteFile(cc.DbPath, data, os.ModePerm)
  }
  if err != nil {
    glog.Error("[UTFHTTP] Failed to save cached data: ", err)
  }
}

func (cc *CachedClient) GetBody(url string) (string, error) {
  body, ok := cc.Cache[url]
  if ok {
    glog.V(2).Info("[UTFHTTP] Hit cache: %s", url)
    return body, nil
  }
  body, err := GetBody(url)
  if err != nil {
    return "", err
  }
  cc.Cache[url] = body
  return body, nil
}

func (cc *CachedClient) GetDocument(url string) (*goquery.Document, error) {
  body, err := cc.GetBody(url)
  if err != nil {
    return nil, err
  }
  return goquery.NewDocumentFromReader(strings.NewReader(body))
}
