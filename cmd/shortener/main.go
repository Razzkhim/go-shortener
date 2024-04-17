package main

import "net/http"

//type Handler struct{}
//
//func (Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
//	data := []byte("hi")
//	res.Write(data)
//}

func mainPage(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello"))
}

func apiPage(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("APIIIIIIIIIII"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(`/asdsa`, apiPage)
	mux.HandleFunc(`/`, mainPage)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
