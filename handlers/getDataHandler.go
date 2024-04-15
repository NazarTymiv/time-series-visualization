package handlers

import (
	"net/http"

	"github.com/NazarTymiv/time-series-visualization/helpers"
)

type Data struct {
	Points []helpers.DataPoint `json:"points"`
	Delta float64 `json:"delta"`
	CurrentData float64 `json:"current_data"`
}

func GetDataHandler(w http.ResponseWriter, r *http.Request) {
	dataPoints := helpers.GenerateTimeSeries(1000)

	startValue := dataPoints[0].Value
	endValue := dataPoints[len(dataPoints) - 1].Value

	data := Data{
		Points: dataPoints,
		Delta: endValue - startValue,
		CurrentData: endValue,
	}

	helpers.RespondWithJSON(w, 200, data)
}