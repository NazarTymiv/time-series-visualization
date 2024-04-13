package main

import (
	"net/http"
)

func getDataHandler(w http.ResponseWriter, r *http.Request) {
	data := generateTimeSeries(10)

	respondWithJSON(w, 200, data)
}