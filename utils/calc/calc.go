package calc

import (
	"math"
)

func DegreeToRadian(deg float64) float64 {
	return deg * (math.Pi / 180.0)
}

func CalculateDistance(lat1, lon1, lat2, lon2 float64, n float64) bool {
	const earthRadius = 6371.0

	lat1Rad := DegreeToRadian(lat1)
	lon1Rad := DegreeToRadian(lon1)
	lat2Rad := DegreeToRadian(lat2)
	lon2Rad := DegreeToRadian(lon2)

	dlon := lon2Rad - lon1Rad
	dlat := lat2Rad - lat1Rad
	a := math.Pow(math.Sin(dlat/2), 2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Pow(math.Sin(dlon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := earthRadius * c

	return distance > n
}
