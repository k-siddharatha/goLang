package postService

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	"fmt"
)

type Message struct {
	username string `json:"username"`
	content  string `json:"content"`
}

type PostService struct {
	id      *int64
	postmap map[int64]Message
}

func NewPostService() PostService {
	var ps PostService
	return ps
}
func GetMessageFromBody(b []byte) Message {
	var t Message
	fmt.Printf("b:%v", b)
  err := json.Unmarshal(b, &t)
  if err != nil {
    return Message{}
  }
  return t
}
func (ps PostService) GetPostFromRequest(rObject http.Request) Message {
  body, err := ioutil.ReadAll(rObject.Body)
  defer rObject.Body.Close()
  if err != nil {
    return Message{}
  }
	m := GetMessageFromBody(body)
	fmt.Printf("Message:%v", m)
	return m
}
