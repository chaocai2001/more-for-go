package fn_field

import (
	"fmt"
	"testing"
)

type IsNumber func(int) bool

func LargeThanTwo(op int) bool {
	return op > 2
}

// Coding Style 1
func (f IsNumber) Apply(s ...int) []bool {
	ret := []bool{}
	for _, e := range s {
		ret = append(ret, f(e))
	}
	return ret
}

func TestMemberApply(t *testing.T) {
	// Coding Style 1
	fmt.Println(IsNumber(LargeThanTwo).Apply(1, 2, 3, 4))
}

// Coding Style 2
func Apply(f IsNumber, s ...int) []bool {
	ret := []bool{}
	for _, e := range s {
		ret = append(ret, f(e))
	}
	return ret
}

func TestApply(t *testing.T) {
	// Coding Style 2
	fmt.Println(Apply(LargeThanTwo, 1, 2, 3, 4))
}
