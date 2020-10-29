package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World")
	})

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Running Hello Handler")

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Error reading body", err)

			http.Error(rw, "Unable to read request body", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(rw, "Hello %s", b)
	})

	// Listen for connections on all ip addresses (0.0.0.0)
	// port 9090
	log.Println("Starting Server")
	err := http.ListenAndServe(":9090", nil)
	log.Fatal(err)
}
