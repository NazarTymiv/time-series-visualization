package helpers

import (
	"math/rand"
	"time"
)

type DataPoint struct {
	Time time.Time `json:"time"`
	Value float64 `json:"value"`
}

func GenerateTimeSeries(pointsNum int) []DataPoint {
	data := make([]DataPoint, pointsNum)

	initialValue := rand.Float64() * 100
	currentValue := initialValue

	currentTime := time.Now()

	for i := 0; i < pointsNum; i++ {
		var delta float64

		for currentValue + delta < 0 || delta == 0{
			delta = 2 - rand.Float64() * 4
		}

		currentValue += delta

		data[i] = DataPoint{
			Time: currentTime.Add(time.Duration(i) * (time.Hour * 24)),
			Value: currentValue,
		}
	}

	return data
}
