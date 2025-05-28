package main

import "fmt"

// 1. Variáveis, Tipos e Constantes

// - Crie uma função que recebe duas variáveis inteiras e retorna a soma.
func soma(a, b int) int {
	return a + b
}

// - Converta um int para float64 e vice-versa.
func switchIntFloat(x any) any {
	switch t := x.(type) {
	case int:
		return float64(t)
	case float64:
		return int(t)
	default:
		fmt.Println("Tipo não suportado")
		return nil
	}
}

// - Escreva uma função que receba um byte e retorne sua representação como rune e string
func byteToRuneAndString(b byte) (rune, string) {
	return rune(b), string(b)
}

func main() {
	// - Crie uma função que recebe duas variáveis inteiras e retorna a soma.
	fmt.Println("## Exercício 1.1")
	fmt.Println("Soma de 10 e 20:", soma(10, 20))

	// - Converta um int para float64 e vice-versa.
	fmt.Println("## Exercício 1.2")
	fmt.Println("Convertendo 10 para float64:", switchIntFloat(10))
	fmt.Println("Convertendo 10.5 para int:", switchIntFloat(10.5))

	// - Use constantes tipadas e não tipadas em um if.
	fmt.Println("## Exercício 1.3")
	const a int = 10
	const b = 20
	if a < b {
		fmt.Println("a é menor que b")
	}

	// - Escreva uma função que receba um byte e retorne sua representação como rune e string
	fmt.Println("## Exercício 1.4")
	c, r := byteToRuneAndString(65)
	fmt.Printf("byte 65 como rune: %q, string: %s", c, r)
}
