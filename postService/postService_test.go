package postService

import (
  "testing"
  "bytes"
  "fmt"
)

func get_string_from_map(m map[string]string) string {
  v := new(bytes.Buffer)
  for key, val := range(m) {
    fmt.Fprintf(v, "%s=\"%s\",\n", key, val)
  }
  return v.String()
}

func TestGetPostFromRequest(t *testing.T) {
  msg := map[string]string{
    "username":"aksinghdce",
    "content":"Some Content",
  }

  var m Message = GetMessageFromBody([]byte(get_string_from_map(msg)))
  if m.username != "aksinghdce" {
    t.Errorf("Username mismatch, expected:%v, output:%v", msg["username"], m.username)
  }
  if m.content != "some content" {
    t.Errorf("Content mismatch, expected:%v, output:%v", msg["content"], m.content)
  }
}
