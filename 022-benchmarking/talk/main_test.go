package talk

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	exp := "Hello Bond"
	s := Greet("Bond")
	if s != exp {
		t.Error("Got", s, "Expected", exp)
	}
}

func ExampleGreet() {
	fmt.Println(Greet("James"))
	// Output:
	// Hello James
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("James")
	}
}
