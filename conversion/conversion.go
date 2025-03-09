package conversion

import (
	"errors"
	"fmt"
	"strconv"
)

func StringToFloat(strings []string) ([]float64, error) {
	floats := make([]float64, len(strings))
	for stringIndex, stringVal := range strings {
		floatPrice, err := strconv.ParseFloat(stringVal, 64)
		if err != nil {
			return nil, errors.New(fmt.Sprint("Failed to convert string to float: ", stringVal))
		}
		floats[stringIndex] = floatPrice
	}
	return floats, nil
}
