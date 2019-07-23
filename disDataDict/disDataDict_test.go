package disDataDict

import (
  "testing"
  "context"
  "log"
)

func TestNew(t *testing.T) {
    kAPI := New()
    resp, err := kAPI.Set(context.Background(), "/foo", "bar", nil)
    if err != nil {
      log.Fatal(err)
      t.Error("kAPI failed")
    } else {
      log.Printf("Set is done, metadata is :%v\n", resp)
    }
}
