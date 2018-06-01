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

func Distance(distance float64, c Conversion) string {

	var rv string
	if distance <= 0 {
		return fmt.Sprintf("0%s", c)
	}

	switch c {
	case Conversions.K:
		rv = Round(distance / 1000)

	case Conversions.M:
		rv = Round(distance * float64(0.00062137))
	}

	return fmt.Sprintf("%s%s", rv, c)
}

func Round(val float64) string {
	var round float64
	roundOn := float64(.5)
	places := float64(2)

	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	_div := math.Copysign(div, val)
	_roundOn := math.Copysign(roundOn, val)
	if _div >= _roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	return strconv.FormatFloat(round/pow, 'f', -1, 64)
}
