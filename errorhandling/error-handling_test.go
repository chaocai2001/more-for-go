package errorhandling

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
)

func f1() error {
	return errors.New("error from f1")
}

func f2() error {
	if err := f1(); err != nil {
		err1 := errors.Wrap(err, "error from f2")
		return err1
	}
	return nil
}

func f3() error {
	if err := f2(); err != nil {
		err1 := errors.Wrap(err, "error from f3")
		return err1
	}
	return nil
}

func TestWrapError(t *testing.T) {
	if err := f3(); err != nil {
		fmt.Printf("%#v\n", err)
	} else {
		fmt.Println("no error occurred")
	}
}
