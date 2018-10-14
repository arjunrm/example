package main

import "fmt"

func main() {
	var arr [5]int
	arr[3] = 10
	fmt.Println(arr)
	fmt.Printf("%T\n", arr)

	// x := type{values} //composite literals
	slice1 := []int{1, 2, 3, 4}
	fmt.Println(slice1)

	for i, v := range slice1 {
		fmt.Printf("%v-%v|", i, v)
	}
	fmt.Println("")

	fmt.Println(arr[3:5])
	fmt.Println(slice1[1:4])

	slice1 = append(slice1, 5, 6, 7)
	fmt.Println(slice1)

	// composite literal
	slice2 := []int{23, 34, 45, 56, 67}
	fmt.Println(slice2)

	slice1 = append(slice1, slice2...)
	fmt.Println(slice1)

	slice1 = append(slice1[:2], slice1[4:]...)
	fmt.Println(slice1)

	slice2 = append(slice2, arr[:]...)
	fmt.Println(slice2)

	slice2 = append(slice2, 12, 23, 34, 45)
	fmt.Println(slice2)

	slice3 := make([]int, 10, 10)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
	slice3 = append(slice3, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	mulslice1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(mulslice1)

	mulslice2 := [][]int{slice1, slice2}
	fmt.Println(mulslice2)

	mulslice3 := make([][]int, 3)
	x := 2
	for i := range mulslice3 {
		mulslice3[i] = make([]int, 3)
		for j := range mulslice3[i] {
			mulslice3[i][j] = (j + 1) * x
			x++
		}
		fmt.Println(mulslice3[i])
	}
	fmt.Println(mulslice3)
	fmt.Printf("%T\n", mulslice3)

	map1 := map[string]int{
		"James": 32,
		"Bond":  50,
	}
	fmt.Println(map1)
	fmt.Println(map1["James"])

	v, ok := map1["Barbra"]
	fmt.Println(v, ok)

	if v, ok = map1["Dummy"]; !ok {
		fmt.Println("NOT OK!! Checking whether key is present in the map")
	}

	map1["Todd"] = 23

	for k, v := range map1 {
		fmt.Println(k, v)
	}

	delete(map1, "James")
	fmt.Println(map1)
	fmt.Printf("%T\n", map1)

	m2 := map[string][]string{
		"Key1": {"1", "2", "3"},
		"Key2": {"A", "B", "C"},
	}
	fmt.Println(m2)
	fmt.Printf("%T\n", m2)

	for i, sl := range m2 {
		fmt.Printf("%v:\n", i)
		for j, v := range sl {
			fmt.Printf("%v,%v | ", j, v)
		}
		fmt.Println("")
	}
}
