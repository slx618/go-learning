package main

import (
	"net/http"
	"net/http/cgi"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler := new (cgi.Handler)
		//handler.Path = os.Getenv("GOPATH")
		handler.Path = "D:\\Program Files\\Go\\bin\\go.exe"
		script := "../Web/" + r.URL.Path
		handler.Dir = "../Web/"
		args := []string{"run", script}
		handler.Args = append(handler.Args, args...)
		handler.ServeHTTP(w, r)
	})
	http.ListenAndServe(":8989",nil)
}
