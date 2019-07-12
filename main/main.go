package main

import (
	"app/factoryMethod"
	"app/singleton"
	"app/postService"
	"fmt"
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
	r.HandleFunc("/post/add", (postHandler{}).Handle())

	return r
}



type postHandler struct{}

func (ps postHandler) Handle() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		/*Create a factory to take http.Request object and return post object*/
		ps := postService.NewPostService()
		msg := ps.GetPostFromRequest(*r)
		fmt.Fprintf(w, fmt.Sprintf("\nmsg:%v\n", msg))
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
