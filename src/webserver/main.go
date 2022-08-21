package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	// Host name of the HTTP Server
	HOST = "localhost"
	// Port of the HTTP Server
	PORT = "8080"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Web Server"))
}

func services(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GET /services")
}

func posts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatalln(err)
	}

 	// We read the response body on the line below
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

 	// Convert the body to type string
	sb := string(body)
	w.Write([]byte(sb))
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/services", services)
	http.HandleFunc("/posts", posts)

	// http.ListenAndServe(":8080", nil)
	err := http.ListenAndServe(":" + PORT, nil)
	if err != nil {
		log.Fatal("Error starting the HTTP Server: ", err)
		return
	}
}
