package packsum

import "testing"

func TestSum(t *testing.T) {
	xi := []int{1, 2, 3, 4, 5}
	v := Sum(xi...)
	if v != 15 {
		t.Fatal("Failed")
	}
}

func TestTableTest(t *testing.T) {
	type test struct {
		data []int
		ans  int
	}
	tests := []test{
		test{[]int{1, 2, 3}, 6},
		test{[]int{2, 3, 4}, 9},
		test{[]int{1, 2, 3, -1, -2, -3}, 0},
	}

	for _, v := range tests {
		x := Sum(v.data...)
		if x != v.ans {
			t.Error("Expected", v.ans, "Got", x)
		}
	}
}
