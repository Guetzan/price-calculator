package converter

import (
	"errors"
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error) {
	var floatValues []float64

	for _, value := range strings {
		convertedValue, err := strconv.ParseFloat(value, 64)

		if err != nil {
			return []float64{}, errors.New("Error at converting string to float")
		}

		floatValues = append(floatValues, convertedValue)
	}

	return floatValues, nil
}
