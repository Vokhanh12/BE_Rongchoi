package main

import "net/http"

// handlerErr used to [main]
func handlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "Something went wrong")
}
