package main

import "fmt"

func main() {
	s := "Hello world"
	fmt.Println(s)

	// string as slice of byte
	// var bs []byte
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%#U\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U", s[i])
	}

	fmt.Println("")

	// returns index and byte value
	for i, v := range s {
		fmt.Print(i, v)
		fmt.Printf("\t%s\t", string(v))
		fmt.Printf("%#x\t", v)
		fmt.Println("")
	}

	var arr [5]int
	for i := range arr {
		arr[i] = i
	}
	fmt.Println(arr)

	const (
		i        = 23
		j uint16 = 90
		k string = "Constant"
	)
	fmt.Println(i, j, k)

	const (
		l = iota
		m
		n
	)
	fmt.Println(l, m, n)
	fmt.Printf("%T\n", l)

	const (
		c1         = iota
		c2 float32 = float32(iota)
		c3 int64   = int64(iota)
	)
	fmt.Println(c1, c2, c3)

	const (
		_  = iota // skip iota 0
		kb = 1 << (iota * 10)
		mb = 1 << (iota * 10)
		gb = 1 << (iota * 10)
	)
	fmt.Println(kb, mb, gb)
}
