package must_test

import (
	"fmt"

	"github.com/dancantos/must"
)

func ExampleMust() {
	fn := func() (string, error) {
		return "Hello World", nil
	}
	value := must.Must(fn())
	fmt.Println(value)
	// Output: Hello World
}

func ExampleOk() {
	fn := func() (string, bool) {
		return "Hello World", true
	}
	value := must.Ok(fn())
	fmt.Println(value)
	// Output: Hello World
}
