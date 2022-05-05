package main

import "fmt"

func VariablesSlice() {
	slice1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice2 := make([]int, 0)

	fmt.Println("slice1:", slice1)

	for _, v := range slice1 {
		if (v % 2) != 0 {
			slice2 = append(slice2, v)
		}

	}

	fmt.Println("slice2:", slice2)

	slice3 := slice1[3:7]

	fmt.Println("slice3:", slice3)

}

func PromedioSlice(dato []int) int {
	prom := 0
	for _, v := range dato {
		prom = prom + v
	}
	prom = prom / len(dato)
	return prom
}
