package main

import "fmt"

func main() {
	var num int
	var n1 int
	var negativo bool
	negativo = false
	fmt.Println("Cuantos numeros deseas agregar ")
	fmt.Scan(&num)
	if num > 0 {
		//s := make([]int, 0, num)
		s := []int{}
		//  fmt.Println(len(s), cap(s))
		for i := 1; i <= num; i++ {
			fmt.Println("Ingresa el numero", i)
			fmt.Scan(&n1)
			if n1 > 0 {
				s = append(s, n1)
				//s[i] = n1
			} else {
				negativo = true
				fmt.Println("No son validos los numeros negativos")
				break
			}
		}

		if !negativo {
			//	fmt.Println(s)
			bubleSort(num, s)

		} else {
			fmt.Println("Intenta de nuevo.")
		}

	} else {
		fmt.Println("Intenta de nuevo.")
	}

}

func bubleSort(num int, s []int) {
	var aux int
	fmt.Println("Arreglo desordenado :")
	fmt.Println(s)
	for i := 0; i < num; i++ {
		for j := 0; j < num; j++ {
			if s[i] < s[j] {
				aux = s[i]
				s[i] = s[j]
				s[j] = aux
			}
		}
	}

	fmt.Println("Arreglo Ordenado :")
	for i := 0; i < num; i++ {
		fmt.Println(s[i])
	}

}
