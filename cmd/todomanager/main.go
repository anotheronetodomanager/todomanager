package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var port = flag.String("port", "8080", "server port")

// example echo server
func main() {
	// parse all flags
	flag.Parse()

	log.Println("Starting echo server on port ", *port)

	// start simple echo server
	err := http.ListenAndServe(fmt.Sprint(":", *port), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// read body bytes
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			writeResponse(w, "", http.StatusInternalServerError)
		}

		// close body to avoid memory leaks
		if err = r.Body.Close(); err != nil {
			log.Println("Failed to close request body:", err)
		}

		// print body to logs
		body := string(bodyBytes)
		log.Println("Received:", body)

		// send echo response
		writeResponse(w, fmt.Sprintln("Echo:", body), http.StatusOK)
	}))
	if err != nil {
		log.Println("Failed to listen and serve:", err)
	}

	log.Println("Shut down successfully")
}

func writeResponse(w http.ResponseWriter, body string, code int) {
	w.WriteHeader(code)
	if _, err := w.Write([]byte(body)); err != nil {
		log.Println("Failed to send response:", err)
	}
}
