package testing101

import "testing"

func testSort(t *testing.T) {
	a := []int{2, 1, 4, 3}

	expected := []int{1, 2, 3, 4}

	//fmt.Println(a)

	MySort(&a)

	actual := a

	//fmt.Println(a)

	if actual != expected {
		t.Errorf("Expected the sum of %v to be %d but instead got %d!", a, expected, actual)
	}
}
