package main

import "net/http"

// handlerReadiness used to [main]
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}
