package main

import "fmt"

func main() {
	fmt.Println("🔁 Exemplos de Loops em Go")

	// ------------------------------
	// 📌 Example01 – For padrão
	// ------------------------------
	// var res int
	// for i := 0; i < 10; i++ {
	// 	res++
	// }
	// fmt.Println("res:", res) // 10

	// ------------------------------
	// 📌 Example02 – For com omissões
	// ------------------------------

	// // Somente condição e post
	// var res, i int
	// for ; i < 10; i++ {
	// 	res++
	// }
	// fmt.Println("res:", res)

	// // Somente condição (post dentro do corpo)
	// res, i = 0, 0
	// for i < 10 {
	// 	res++
	// 	i++
	// }
	// fmt.Println("res:", res)

	// // Loop infinito com break
	// res = 0
	// for {
	// 	res++
	// 	break
	// }
	// fmt.Println("res:", res) // 1

	// ------------------------------
	// 📌 Example03 – For com range
	// ------------------------------
	// arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// // Apenas iterando
	// for range arr {
	// 	fmt.Println("loop")
	// }

	// // Apenas índice
	// for i := range arr {
	// 	fmt.Println("índice:", i)
	// }

	// // Índice e valor
	// for i, elem := range arr {
	// 	fmt.Println("i:", i, "valor:", elem)
	// }

	// // Apenas valor (ignorando índice)
	// for _, elem := range arr {
	// 	fmt.Println("valor:", elem)
	// }

	// ------------------------------
	// 📌 Example04 – range sobre inteiro (Go 1.22+)
	// ------------------------------

	// for range 5 {
	// 	fmt.Println("loop")
	// }

	// for i := range 5 {
	// 	fmt.Println("índice:", i)
	// }
}
