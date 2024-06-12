package must_test

import (
	"errors"
	"testing"

	"github.com/dancantos/must"
)

func TestMust(t *testing.T) {
	if result := must.Must(1, nil); result != 1 {
		t.Fatalf("must.Must(1, nil) == %d, expected %d", result, 1)
	}
}

func TestMustPanics(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Fatal("must.Must(0, err) failed to panic")
		}
	}()
	must.Must(0, errors.New("bad things"))
}

func TestOk(t *testing.T) {
	if result := must.Ok(1, true); result != 1 {
		t.Fatalf("must.Must(1, nil) == %d, expected %d", result, 1)
	}
}

func TestOkPanics(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Fatal("must.Must(0, false) failed to panic")
		}
	}()
	must.Ok(0, false)
}
