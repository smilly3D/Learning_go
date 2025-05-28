package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("🧪 Exemplos de estruturas condicionais (if) em Go")

	// ------------------------------
	// 📌 Example01 – Declaração normal
	// ------------------------------
	// if 1 < 2 {
	// 	fmt.Println("1 é menor que 2")
	// }

	// ------------------------------
	// 📌 Example02 – Short statement no if
	// ------------------------------
	// if x := math.Sqrt(4); x < 10 {
	// 	fmt.Println("Raiz menor que 10:", x)
	// }

	// ------------------------------
	// 📌 Example03 – Short statement com else if
	// ------------------------------
	// if x := math.Sqrt(4); x < 1 {
	// 	fmt.Println("x é menor que 1")
	// } else if x > 0 {
	// 	fmt.Println("x é maior que 0")
	// }

	// ------------------------------
	// 📌 Example04 – Short statement com else if e else
	// ------------------------------
	// if x := math.Sqrt(4); x < 1 {
	// 	fmt.Println("x é menor que 1")
	// } else if x < 0 {
	// 	fmt.Println("x é menor que 0")
	// } else {
	// 	fmt.Println("nenhuma das condições anteriores foi satisfeita")
	// }

	// ------------------------------
	// 📌 Example05 – Tratamento de erro com short statement
	// ------------------------------
	// if err := doError(); err != nil {
	// 	fmt.Println("Erro detectado:", err)
	// }

	// Padrão idiomático:
	// if err := doSomething(); err != nil {
	// 	return err
	// }
}

func doError() error {
	return errors.New("erro de exemplo")
}

// func doSomething() error {
// 	// simulação de alguma ação que pode falhar
// 	return nil
// }
