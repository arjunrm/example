package talk

import "fmt"

// Greet is used to greet a person
func Greet(s string) string {
	return fmt.Sprint("Hello ", s)
}

// HowAreYou is dummy func
func HowAreYou() {
	fmt.Println("How are you")
}
