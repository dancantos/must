package must_test

import (
	"fmt"
	"testing"

	"github.com/dancantos/must"
)

func ExampleMust(t *testing.T) {
	fn := func() (string, error) {
		return "Hello World", nil
	}
	value := must.Must(fn())
	fmt.Println(value)
	// Output: Hello World
}

func ExampleOk(t *testing.T) {
	fn := func() (string, bool) {
		return "Hello World", true
	}
	value := must.Ok(fn())
	fmt.Println(value)
	// Output: Hello World
}
