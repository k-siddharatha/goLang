package main

import (
	"fmt"
	"net/http"

	"app/factoryMethod"
	"app/singleton"
)

func main() {
	http.ListenAndServe("0.0.0.0:8080", handler())
}

func handler() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/singleton", (singletonHandler{}).Handle())
	r.HandleFunc("/factorymethod", (factoryMethodHandler{}).Handle())
	return r
}

type Handler interface {
	Handle(command string) func(http.ResponseWriter, *http.Request)
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

func commandHandler(resWriter http.ResponseWriter, r *http.Request) {
	s := singleton.New()
	s["Amit"] = "Kumar"
	s2 := singleton.New()
	fmt.Fprint(resWriter, fmt.Sprintf("\n Received a request singleton2['Amit']=%v \n %v \n %v \n %v",
		s2["Amit"], factoryMethod.NewStore(1).Open("x"), factoryMethod.NewStore(2).Open("x"), factoryMethod.NewStore(4).Open("x")))
}
