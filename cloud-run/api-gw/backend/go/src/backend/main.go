package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// APIURL is for...
var APIURL string

func main() {

	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/api", apiHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	APIURL = os.Getenv("API_URL")
	if APIURL == "" {
		log.Fatalln("API_URL env variable missing!")
	}

	log.Println("** Service Started on Port " + port + " **")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	aboutMsg := fmt.Sprintf("{\"status\":\"ok\", \"api_url\": \"%s\"}", APIURL)

	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, aboutMsg)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(APIURL)
	if err != nil {
		log.Print(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", body)
}
