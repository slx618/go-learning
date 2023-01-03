package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func f2(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	query := r.URL.Query()
	fmt.Println(query.Get("page"))
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body))
	w.Write([]byte("ok"))
}

func main() {
	fmt.Println("http://127.0.0.1:9090")

	http.HandleFunc("/", f1)
	http.HandleFunc("/x", f2)
	http.ListenAndServe("127.0.0.1:9090", nil)
}
