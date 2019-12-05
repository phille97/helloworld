package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	name := os.Getenv("NAME")
	listen := os.Getenv("LISTEN")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello %s", name)
	})

	log.Printf("Up and running! Listening on %s", listen)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN"), nil))
}
