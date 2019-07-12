package singleton

import (
  "testing"
)

func TestNew(t *testing.T) {
  singleton1 := New()
  singleton1["Amit"] = "Singh"
  singleton2 := New()
  expected := "Singh"
  if singleton2["Amit"] != expected {
    t.Error("Singleton Failed!")
  }
}
