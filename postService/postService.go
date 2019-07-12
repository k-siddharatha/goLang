package postService

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
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

func (ps PostService) GetPostFromRequest(rObject http.Request) Message {
  body, err := ioutil.ReadAll(rObject.Body)
  defer rObject.Body.Close()
  if err != nil {
    return Message{}
  }
  var t Message
  err = json.Unmarshal(body, &t)
  if err != nil {
    return Message{}
  }
  return t
}
