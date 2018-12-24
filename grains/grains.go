package grains

import (
	"errors"
	"math"
	"strconv"
)

func Square(boxNum int) (uint64, error) {
	if boxNum < 1 || boxNum > 64 {
		return 0, errors.New("Box number <" + strconv.Itoa(boxNum) + "> is not valid")
	}
	if boxNum == 1 {
		return 1, nil
	}
	return uint64(math.Pow(2, float64(boxNum-1))), nil
}

func Total() uint64 {
	var sum uint64 = 1
	for i := 2; i < 65; i++ {
		sum += uint64(math.Pow(2, float64(i-1)))
	}
	return sum
}
