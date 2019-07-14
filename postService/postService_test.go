package postService

import (
  "testing"
  "encoding/json"
)

func TestGetPostFromRequest(t *testing.T) {
  message := map[string]interface{}{
		"username": "aksinghdce",
		"content":  "This is the content",
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		t.Errorf("Error in Marshal:%v",err)
	}

  m := GetMessageFromBody(bytesRepresentation)
  expected_username := message["username"]
  expected_content := message["content"]
  if m.username != expected_username {
    t.Errorf("Byte Representation of Input:%v", bytesRepresentation)
    t.Errorf("Username mismatch, expected:%v, output:%v", expected_username, m.username)
  }
  if m.content != expected_content {
    t.Errorf("Content mismatch, expected:%v, output:%v", expected_content, m.content)
  }
}
