package main

import "fmt"

func main() {
	var cantidad int
	fmt.Print("Ingrese la Cantidad: ")
	fmt.Scanln(&cantidad)
	fmt.Println("Cantidad vale:", cantidad)
	slice0 := []int{}
	var valor int
	for i := 0; i < cantidad; i++ {
		fmt.Println("Ingrese el valor del dato en la posicion", i, ":")
		fmt.Scanln(&valor)
		slice0 = append(slice0, valor)
	}
}
