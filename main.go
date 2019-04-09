package main 

import "fmt"
import "net/http"

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var message = "Welcome Index"
	w.Write([]byte(message))
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	var message = "Hello World"
	w.Write([]byte(message))
}

//main router for handling index and hello

func main() {
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/hello", handlerHello)

	var server = "localhost:9000"
	fmt.Println("server started at", server)
	err := http.ListenAndServe(server, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
