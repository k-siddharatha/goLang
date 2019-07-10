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

func handler() http.Handler {
	r := http.NewServeMux()
	endpoints []Handle{singletonHandler{}, factoryMethodHandler{}, postService{}}
	for i:= range(endpoints) {
		r.HandleFunc("/singleton", i.Handle())
	}
//	r.HandleFunc("/singleton", (singletonHandler{}).Handle())
//	r.HandleFunc("/factorymethod", (factoryMethodHandler{}).Handle())
//	r.HandleFunc("/post/add", (postService{}).Handle())

	return r
}

type Handler interface {
	GetEndpoint() string
	Handle(command string) func(http.ResponseWriter, *http.Request)
}


type message struct {
	username string `json:"username"`
	content  string `json:"content"`
}

type postService struct {
	id      *int64
	postmap map[int64]message
}

func (ps postService) GetEndpoint() string {
	return "/post/add"
}

func (ps postService) Handle() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		t := new(message)
		body, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(body, &t)
		if err != nil {
			fmt.Fprintf(w, "Panic")
			panic(err)
		}
		fmt.Fprintf(w, fmt.Sprintf("t.username:%v, and t.content:%v", t.username, t.content))

	}
}

type singletonHandler struct{}
func (sh singletonHandler) GetEndpoint() string {
	return "/singleton"
}
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

func (fm factoryMethodHandler) GetEndpoint() string {
	return "/factorymethod"
}

func (fmh factoryMethodHandler) Handle() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, fmt.Sprintf("\n %v \n %v \n %v",
			factoryMethod.NewStore(1).Open("x"), factoryMethod.NewStore(2).Open("x"), factoryMethod.NewStore(4).Open("x")))
	}
}

