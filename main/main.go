package main

import (
	"fmt"
	"net/http"
	"urlshortner"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string {
		"/urlshort-godoc": "https://godoc.org/github.com/grophercises/urlshort",
		"/yaml-godoc": "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshortner.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the fallback
	yaml := getYAML()
	// mapHandler is the fall back for yaml
	// so request for /urlshort-godoc will go to
	// YAMLHandler -> fallback to mapHandler and served
	yamlHandler, err := urlshortner.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("starting server on 8080")
	_ = http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request)  {
	_, _ = fmt.Fprintln(w, "Hello, world!")
}

func getYAML() string {
	return `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/grophercises/urlshort/tree/solution
`
}