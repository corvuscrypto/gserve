package gserve

import (
	"encoding/json"
	"log"
	"os"

	"github.com/corvuscrypto/ggi"
)

//ScanRoutes scans the file at the given path and adds handlers to the ggi
//server
func ScanRoutes(filepath string) {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	//decode into a string-string map
	var m map[string]string
	err = json.NewDecoder(f).Decode(m)
	if err != nil {
		log.Fatal(err)
	}
	for route, file := range m {
		ggi.RegisterRoute(route, file)
	}
}

//AddRoute registers a single route to a filepath
func AddRoute(route, filepath string) {
	ggi.RegisterRoute(route, filepath)
}
