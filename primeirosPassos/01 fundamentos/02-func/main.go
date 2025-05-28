package main

import "fmt"

// Função sem parâmetros e sem retorno
func digaOi() {
	fmt.Println("Oi")
}

// Função com parâmetros e retorno
func somar(a int, b int) int {
	return a + b
}

// Função com múltiplos retornos
func swap(a, b int) (int, int) {
	return b, a
}

// Função com múltiplos retornos e nomeação
func dividir(a, b int) (res int, rem int) {
	res = a / b
	rem = a % b
	return
}

// Higher-order function: retorna uma função
func somarCom(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

// Função variádica
func somarTudo(nums ...int) int {
	var out int
	for _, num := range nums {
		out += num
	}
	return out
}

// First-class function
func multiplicar(a, b int) int {
	return a * b
}

// Closure
func contador() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// 1. digaOi
	digaOi()

	// 2. somar
	fmt.Println("somar(1, 2):", somar(1, 2))

	// 3. swap
	a, b := swap(3, 4)
	fmt.Println("swap(3, 4):", a, b)

	// 4. dividir
	res, rem := dividir(25, 4)
	fmt.Println("dividir(25, 4):", res, rem)

	// 5. higher-order function
	result := somarCom(10)(5)
	fmt.Println("somarCom(10)(5):", result)

	// 6. função variádica
	fmt.Println("somarTudo(1, 2, 3, 4, 5):", somarTudo(1, 2, 3, 4, 5))

	// 7. função como valor (first-class function)
	var op func(int, int) int = multiplicar
	fmt.Println("op(3, 4):", op(3, 4))

	// 8. closure
	proximo := contador()
	fmt.Println("contador():", proximo()) // 1
	fmt.Println("contador():", proximo()) // 2

	// 9. função anônima
	func(msg string) {
		fmt.Println("Função anônima:", msg)
	}("Olá, mundo!")
}
