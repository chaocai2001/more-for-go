package generator

import (
	"fmt"
	"testing"
)

func carGenerator(itr func(int) int, lower int, upper int) func() (int, bool) {
	return func() (int, bool) {
		lower = itr(lower)
		return lower, lower > upper
	}
}

func iterator(i int) int {
	i += 1
	return i
}

func TestPrintCar(t *testing.T) {
	lower := 10
	for i := 1; i < 5; i++ {
		lower, _ = carGenerator(iterator, lower, 20)()
		fmt.Println(lower)
	}
}
