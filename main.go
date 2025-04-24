package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		statusNotFoundHandler(w, r)
		return
	}

	w.Write([]byte("Hello world!"))
}

func hello(w http.ResponseWriter, r *http.Request) {
	pathRegexp := regexp.MustCompile(`^/hello/\w+$`)
	if !pathRegexp.Match([]byte(r.URL.Path)) {
		statusNotFoundHandler(w, r)
		return
	}

	name := strings.Split(r.URL.Path, "/")[2]
	w.Write([]byte(fmt.Sprintf("Hello, %s", name)))
}

func statusNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 page not found"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/hello/", hello)

	err := http.ListenAndServe(":8081", mux)

	if err != nil {
		panic(err)
	}
}
