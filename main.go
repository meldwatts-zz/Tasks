package main

import (
	"log"
	"net/http"
)

func main() {
	const port = "127.0.0.1:8080"
	log.Print("Running server on " + port)
	http.HandleFunc("/", completedTaskFunc)
	log.Fatal(http.ListenAndServe(port, nil))
}

func completedTaskFunc(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(r.URL.Path))

}
