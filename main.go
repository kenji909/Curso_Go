package main

import (
	"Ejercicio1.go"
	"Ejercicio2.go" 
	"Ejercicio3.go"
)
	
func main2() {
	fmt.Println("----- Ejercicio 1 -----")
	VariablesSlice()
	fmt.Println("Ingrese la cantidad de datos a cargar:")
	cantidad := 0
	fmt.Scanln(&cantidad)
	slice0 := [cantidad]int()
	valor := 0
	for i,_ range slice0{
		fmt.Println("Ingrese el valor del dato en la posicion",i,":")
		slice0[i] = fmt.Scanln(&valor)
	}
	resultado1 := PromedioSlice(slice0)
	fmt.Println("El promedio es:",resultado1)
	fmt.Println("----- Ejercicio 2 -----")
	fmt.Println("----- Ejercicio 3 -----")
}
