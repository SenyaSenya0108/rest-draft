package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc(`/product/{id:\d+}`, product)

	err := http.ListenAndServe(":8081", router)
	if err != nil {
		panic(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Home")
}

func product(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "Product ID %s", id)
}
