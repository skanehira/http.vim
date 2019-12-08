package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b, _ := httputil.DumpRequest(r, true)
		fmt.Println(string(b))
		fmt.Fprintln(w, "Hello Gorilla")
	})

	log.Fatal(http.ListenAndServe(":80", nil))
}
