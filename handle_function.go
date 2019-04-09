package main

import "fmt"
import "net/http"

func main() {
	handlerIndex := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	}

	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Data"))
	})

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

