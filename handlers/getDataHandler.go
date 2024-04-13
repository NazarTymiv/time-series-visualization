package handlers

import (
	"net/http"

	"github.com/NazarTymiv/time-series-visualization/helpers"
)

func GetDataHandler(w http.ResponseWriter, r *http.Request) {
	data := helpers.GenerateTimeSeries(10)

	helpers.RespondWithJSON(w, 200, data)
}