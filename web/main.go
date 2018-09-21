package main

import (
	"log"
	"net/http"
)

func sayhello(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", sayhello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
