package httpx

import (
  "testing"
)

func TestGetWithUA(t *testing.T) {
  succ_cases := []string {
    "http://m.facebook.com",
  }
  fail_cases := []string {
    "NOT_AN_URL",
  }

  for _, url := range succ_cases {
    _, err := GetWithUA(url, UA_iPhone_6_Plus)
    if err != nil {
      t.Errorf("Failed: %s", url)
    }
  }
  for _, url := range fail_cases {
    _, err := GetWithUA(url, UA_iPhone_6_Plus)
    if err == nil {
      t.Errorf("Should fail while not: %s", url)
    }
  }
}
