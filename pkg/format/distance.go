package format

import (
	"fmt"
	"math"
	"strconv"
)

type Conversion string

var Conversions = struct {
	K Conversion
	M Conversion
}{
	"k",
	"mi",
}

var toMiles = float64(0.00062137)

func Distance(distance float64, c Conversion) string {

	var rv string
	if distance <= 0 {
		return fmt.Sprintf("0%s", c)
	}

	switch c {
	case Conversions.K:
		rv = fixFloat(distance / 1000)
	case Conversions.M:
		rv = fixFloat(distance * float64(0.00062137))
	}

	return fmt.Sprintf("%s%s", rv, c)
}

func fixFloat(in float64) string {
	ceil := math.Ceil(in*100) / 100
	return strconv.FormatFloat(ceil, 'f', -1, 64)
}
