package main

import (
	"fmt"
	"time"
)

// 📌 Example01 – Uso simples
func doDefer() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

// 📌 Example02 – Ordem reversa de execução
func deferReverseOrder() {
	defer fmt.Println("um")
	defer fmt.Println("dois")
	defer fmt.Println("três")
	fmt.Println("início")
}

// 📌 Example03 – Defer com função anônima
func deferWithAnonymous() {
	defer func() {
		fmt.Println("fim da função")
	}()
	fmt.Println("início da função")
}

// 📌 Example04 – Executado mesmo com panic
func deferWithPanic() {
	defer fmt.Println("sempre executa, mesmo após panic")
	panic("ocorreu um erro")
}

// 📌 Example05 – Medindo tempo de execução
func trackTime(start time.Time, name string) {
	duration := time.Since(start)
	fmt.Printf("Função %s demorou %s\n", name, duration)
}

func simulateWork() {
	defer trackTime(time.Now(), "simulateWork")
	time.Sleep(2 * time.Second)
}

// 🧪 Execução dos exemplos
func main() {
	fmt.Println("Exemplos de uso do defer:")

	// doDefer()
	// deferReverseOrder()
	// deferWithAnonymous()
	// deferWithPanic()
	// simulateWork()
}
