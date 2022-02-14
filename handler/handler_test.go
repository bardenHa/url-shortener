package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var yml = `
- path: /twtr
  url: https://twitter.com
- path: /me
  url: https://github.com/bardenHa
`

func TestParseYaml(t *testing.T) {
	expectedOutput := []pathUrl{
		{Path: "/twtr", URL: "https://twitter.com"},
		{Path: "/me", URL: "https://github.com/bardenHa"},
	}

	output, err := parseYaml([]byte(yml))

	if err != nil {
		t.Errorf("Expected no error but got: %v", err)
	}

	if output[0].Path != expectedOutput[0].Path {
		t.Errorf("Path incorrect. Got: %q, Expected: %q", output[0].Path, expectedOutput[0].Path)
	}

	if output[0].URL != expectedOutput[0].URL {
		t.Errorf("URL incorrect. Got: %q, Expected: %q", output[0].Path, expectedOutput[0].Path)
	}
}

func TestBuildMap(t *testing.T) {
	input := []pathUrl{
		{Path: "/twtr", URL: "https://twitter.com"},
		{Path: "/me", URL: "https://github.com/bardenHa"},
	}

	output := buildMap(input)

	fmt.Println(output)

	if len(output) != 2 {
		t.Errorf("Returned wrong length of map. Got: %v, Expected: %v", len(output), 2)
	}

	url, ok := output["/twtr"]

	if !ok {
		t.Errorf("Url is not returned in map. Got: %v, Expected: %v", ok, true)
	}

	if url != "https://twitter.com" {
		t.Errorf("Incorrect url returned. Got: %q, Expected: %q", url, "https://twitter.com")
	}
}

func TestYAMLHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/twtr", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	})

	yamlHandler, err := YAMLHandler([]byte(yml), mux)

	if err != nil {
		t.Errorf("Expected no error but got: %v", err)
	}

	handler := http.HandlerFunc(yamlHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != 302 {
		t.Errorf("Incorrect response code. Got: %v, Expected: %v", rr.Code, 302)
	}

	fmt.Println(rr.Header().Get("Location"))

	if rr.Header().Get("Location") != "https://twitter.com" {
		t.Errorf("Incorrect response body. Got: %q, Expected: %q", rr.Header().Get("Location"), "https://twitter.com")
	}
}
