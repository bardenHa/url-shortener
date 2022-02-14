package handler

import (
	"testing"
)

func TestParseYaml(t *testing.T) {
	input := `
- path: /twtr
  url: https://twitter.com
- path: /me
  url: https://github.com/bardenHa
`
	expectedOutput := []pathUrl{
		{Path: "/twtr", URL: "https://twitter.com"},
		{Path: "/me", URL: "https://github.com/bardenHa"},
	}

	output, err := parseYaml([]byte(input))

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
