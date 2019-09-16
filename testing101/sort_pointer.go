package testing101

//package blah

func MySort(arr *[]int) {

	len := len(*arr)
	for i := 0; i < len; i++ {
		for j := 0; j < len; j++ {
			if (*arr)[i] < (*arr)[j] {
				c := (*arr)[i]
				(*arr)[i] = (*arr)[j]
				(*arr)[j] = c
			}
		}
	}
}

/*
func main() {

	a := []int{32, 14, 25, 6, 1, 123, 44, 54, 23}

	fmt.Println(a)

	sort(&a)

	fmt.Println(a)
}
*/
