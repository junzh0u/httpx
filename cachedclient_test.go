package utfhttp

import (
  "testing"
)

func BenchmarkCachedClientGetSuccess(b *testing.B) {
  succ_cases := []string {
    "http://junz.info",
  }

  cc, _ := NewCachedClient("/tmp/utfhttp.cache.json")
  defer cc.Close()
  for _, url := range succ_cases {
    for i := 0; i < b.N; i++ {
      _, err := cc.GetDocument(url)
      if err != nil {
        b.Errorf("Failed: %s", url)
      }
    }
  }
}

func BenchmarkCachedClientGetFail(b *testing.B) {
  fail_cases := []string {
    "http://junz.info/404",
  }

  cc, _ := NewCachedClient("/tmp/utfhttp.cache.json")
  defer cc.Close()
  for _, url := range fail_cases {
    for i := 0; i < b.N; i++ {
      _, err := cc.GetDocument(url)
      if err == nil {
        b.Errorf("Should fail while not: %s", url)
      }
    }
  }
}
