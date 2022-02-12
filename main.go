package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bardenHa/urlshortener/handler"
)

var pathsToUrls map[string]string = map[string]string{
	"/ggle":  "https://www.google.co.uk",
	"/ytube": "https://www.youtube.com",
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", welcome)

	mapHandler := handler.MapHandler(pathsToUrls, mux)

	yaml, err := ioutil.ReadFile("shortenURLs.yaml")
	if err == nil {
		panic(err)
	}

	yamlHandler, err := handler.YAMLHandler([]byte(yaml), mapHandler)

	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func welcome(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "welcome")
}
