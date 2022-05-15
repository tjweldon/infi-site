package web

import (
	"log"
	"net/http"
)

func LogRequests(controller Controller) (wrapped Controller) {
	wrapped = func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Serving %s", r.URL.String())
		controller(w, r)
	}

	return wrapped
}