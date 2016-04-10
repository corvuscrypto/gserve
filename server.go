package gserve

import (
	"net/http"

	"github.com/corvuscrypto/ggi"
)

//StartServer initializes a default server to begin handling requests via ggi handlers
func StartServer() {
	http.HandleFunc("/", ggi.HandleRequest)
	http.ListenAndServe(":80", nil)
}
