package urlshortner

import (
	"gopkg.in/yaml.v2"
	"net/http"
)

/*
	MapHandler will return an http.HandlerFunc (which also
	implements http.Handler) that will attempt to map any
	paths (keys in the map) to their corresponding URL (values
	that each key in the map points to, in string format).
	if the path is not provided in the map, then the fallback
	http.Handler will be called instead
*/
func MapHandler(pathToUrls map[string]string, fallback http.Handler) http.HandlerFunc  {

	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		// if we match a path -> we redirect to it
		if dest, ok := pathToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		// else
		fallback.ServeHTTP(w, r)
	}
}

/*
	YAMLHandler will parse the provided YAML and then return
	and http.HandlerFunc (which also implements http.Handler)
	that will attempt to map any paths to their corresponding
	URL. If the path is not provided in the YAML, then the
	fallback http.Handler will be called instead.

	YAML is expected to be in the format:
		- path: /some-path
		  url: https://www.some-url.com/demo
	The only errors that can be returned all related to having
	invalid YAML data.

	See MapHandler to create a similar http.HandlerFunc via
	a mapping YAML data.
*/
func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// parse yaml -> convert to map -> return map handler
	var pathUrls []pathUrl
	err := yaml.Unmarshal(yamlBytes, &pathUrls)
	if err != nil {
		return nil, err
	}

	pathsToUrls := make(map[string]string)
	for _, pu := range pathUrls {
		pathsToUrls[pu.Path] = pu.URL
	}

	return MapHandler(pathsToUrls, fallback), nil
}

type pathUrl struct {
	Path string `yaml:"path"`
	URL string `yaml:"url"`
}