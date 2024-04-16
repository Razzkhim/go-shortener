package main

import (
	"net/http"
)

type Handler struct{}

func (h Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	data := []byte("Hello!")
	res.Write(data)
}

func main() {
	var h Handler

	err := http.ListenAndServe(`:8080`, h)
	if err != nil {
		panic(err)
	}

}
