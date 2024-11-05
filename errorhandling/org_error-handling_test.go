package errorhandling

import (
	"errors"
	"fmt"
	"testing"
)

func f_1() error {
	return errors.New("error from f1")
}

func f_2() error {
	if err := f_1(); err != nil {
		err1 := fmt.Errorf("from f2: %w", err)
		return err1
	}
	return nil
}

func f_3() error {
	if err := f_2(); err != nil {
		err1 := fmt.Errorf("from f3: %w", err)
		return err1
	}
	return nil
}

func TestWrapErrorOrg(t *testing.T) {
	if err := f_3(); err != nil {
		fmt.Printf("%+v\n", err)
	} else {
		fmt.Println("no error occurred")
	}
}
