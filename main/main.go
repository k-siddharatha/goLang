package main

import (
	"app/factoryMethod"
	"app/singleton"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.ListenAndServe("0.0.0.0:8080", handler())
}

type Handler interface {
	Handle() func(http.ResponseWriter, *http.Request)
}


func handler() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/singleton", (singletonHandler{}).Handle())
	r.HandleFunc("/factorymethod", (factoryMethodHandler{}).Handle())
	r.HandleFunc("/post/add", (postService{}).Handle())

	return r
}



type message struct {
	username string `json:"username"`
	content  string `json:"content"`
}

type postService struct {
	id      *int64
	postmap map[int64]message
}

func (ps postService) Handle() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, fmt.Sprintf("%v", r))
		var t message
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		err = json.Unmarshal(body, &t)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		fmt.Fprintf(w, fmt.Sprintf("\nt:%v\n", t))
		fmt.Fprintf(w, fmt.Sprintf("t.username:%v, and t.content:%v", t.username, t.content))
	}
}

type singletonHandler struct{}
func (sh singletonHandler) Handle() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		s := singleton.New()
		s["Amit"] = "Kumar"
		s2 := singleton.New()
		fmt.Fprint(w, fmt.Sprintf("\n Received a request singleton2['Amit']=%v",
			s2["Amit"]))

	}
}

type factoryMethodHandler struct{}


func (fmh factoryMethodHandler) Handle() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, fmt.Sprintf("\n %v \n %v \n %v",
			factoryMethod.NewStore(1).Open("x"), factoryMethod.NewStore(2).Open("x"), factoryMethod.NewStore(4).Open("x")))
	}
}
