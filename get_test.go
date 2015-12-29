package utfhttp

import (
  "testing"
)

func TestGetDocument(t *testing.T) {
  succ_cases := []string {
    "http://junz.info",
  }
  fail_cases := []string {
    "http://junz.info/404",
  }

  for _, url := range succ_cases {
    _, err := GetDocument(url)
    if err != nil {
      t.Errorf("Failed: %s", url)
    }
  }
  for _, url := range fail_cases {
    _, err := GetDocument(url)
    if err == nil {
      t.Errorf("Should fail while not: %s", url)
    }
  }
}
