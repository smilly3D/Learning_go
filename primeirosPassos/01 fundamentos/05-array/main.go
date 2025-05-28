package main

import "fmt"

func main() {
	fmt.Println("🧪 Testes com arrays em Go")

	// ------------------------------
	// 📌 Example01 – Declaração de array vazio e inicializado
	// ------------------------------

	// var arr1 = [5]int{}
	// fmt.Println("arr1:", arr1) // [0 0 0 0 0]

	// var arr2 = [5]int{1, 2, 3, 4, 5}
	// fmt.Println("arr2:", arr2) // [1 2 3 4 5]

	// arr3 := [5]int{}
	// fmt.Println("arr3:", arr3) // [0 0 0 0 0]

	// arr4 := [5]int{1, 2, 3, 4, 5}
	// fmt.Println("arr4:", arr4) // [1 2 3 4 5]

	// ------------------------------
	// 📌 Example02 – Inicialização com índice específico
	// ------------------------------

	// arr := [3]int{1}
	// fmt.Println("arr:", arr) // [1 0 0]

	// arr2 := [10]int{5: 400}
	// fmt.Println("arr2:", arr2) // [0 0 0 0 0 400 0 0 0 0]

	// arr3 := [10]int{1: 100, 5: 400}
	// fmt.Println("arr3:", arr3) // [0 100 0 0 0 400 0 0 0 0]

	// arr4 := [5]string{"a", "b", "c", "d", "e"}
	// fmt.Println("arr4:", arr4) // [a b c d e]

	// arr5 := [10]string{1: "hello", 5: "world"}
	// fmt.Println("arr5:", arr5) // [  hello      world       ]

	// ------------------------------
	// 📌 Example03 – Usando constantes para tamanho do array
	// ------------------------------

	// var tamanho = 5
	// var arrErro [tamanho]int // erro: tamanho não é constante
	// fmt.Println(arrErro)

	// const tamanhoConst = 5
	// var arrOk [tamanhoConst]int
	// fmt.Println("arrOk:", arrOk) // [0 0 0 0 0]
}
