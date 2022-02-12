package main

import (
	"fmt"
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

	yaml := `
- path: /twtr
  url: https://twitter.com/
- path: /me
  url: https://github.com/bardenHa
`

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
