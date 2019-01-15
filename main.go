// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Docker")
	})

	fmt.Println("Listening on :8085")
	log.Fatal(http.ListenAndServe(":8085", nil))
}
