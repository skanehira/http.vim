package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b, _ := httputil.DumpRequest(r, true)
		fmt.Println(string(b))
		fmt.Fprintln(w, "Hello Gorilla")
	})

	http.ListenAndServe(":80", nil)
}
