package main

import (
	"basic-http/config"
	ctr "basic-http/controller"
	"fmt"
	"log"
	"net/http"
)

func routes() {
	fileServer := http.FileServer(http.Dir("./view")) // New code
	http.Handle("/", fileServer)

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong!")
	})

	http.HandleFunc("/upload", ctr.UploadHandler)
}

func main() {
	routes()
	fmt.Printf("Starting server at port %s\n", config.PORT)
	if err := http.ListenAndServe(":"+config.PORT, nil); err != nil {
		log.Fatal(err)
	}
}
