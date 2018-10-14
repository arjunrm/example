package main

import "fmt"

func main() {
	for i := 65; i < 123; i++ {
		// print value, hex, unicode with character
		fmt.Printf("%v\t%#x\t%#U\n", i, i, i)
	}

	for i := 6500; i < 6600; i++ {
		// print value, hex, unicode with character
		fmt.Printf("%v\t%#x\t%#U\n", i, i, i)
	}

	switch {
	case false:
		fmt.Println("False switch")
	case (3 == 3):
		fmt.Println("3==3")
		fallthrough // fall through is not by default
	case (4 == 4):
		fmt.Println("4==4")
	case (5 == 5):
		fmt.Println("5==5")
	default:
		fmt.Println("Default")
	}
}
