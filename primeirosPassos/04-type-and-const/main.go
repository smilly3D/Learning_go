package main

import (
	"fmt"
)

//boolean (bool)

//inteiro (int, int8, int16, int32, int64)
//inteiro sem sinal (uint, uint8, uint16, uint32, uint64)

//alias
//byte = uint8
//rune = int32

//float32, float64

//complex64, complex128

// 📌 Example01 – Tipagem forte
// func somar(a, b int) int {
// 	return a + b
// }

// 📌 Example02 – Type alias: byte
// func takeByte(b byte) {
// 	fmt.Println("takeByte (Example02):", b)
// }

// 📌 Example07 e Example08 – Constantes com int32
// func takeInt32(x int32) {
// 	fmt.Println("takeInt32 (Example07):", x)
// }

func main() {
	fmt.Println("🔎 Testes do sistema de tipos e constantes")

	// ------------------------------
	// 📌 Example01 – Tipagem forte
	// ------------------------------
	// fmt.Println(somar("a", "b")) // erro de compilação

	// ------------------------------
	// 📌 Example02 – Uso de type alias (byte)
	// ------------------------------
	// var b uint8 = 10
	// takeByte(b)

	// ------------------------------
	// 📌 Example03 – int para float64 e tentativa para bool
	// ------------------------------
	// var x int = 65
	// f := float64(x)
	// b := bool(x) // erro
	// fmt.Println("float64:", f, "bool:", b)

	// ------------------------------
	// 📌 Example04 – int para string (ASCII)
	// ------------------------------
	// var x int = 65
	// s := string(x)
	// fmt.Println("string(x):", s) // "A"

	// ------------------------------
	// 📌 Example05 – strconv.FormatInt
	// ------------------------------
	// var x int = 65
	// s := strconv.FormatInt(int64(x), 10)
	// fmt.Println("strconv:", s) // "65"

	// ------------------------------
	// 📌 Example06 – Constante tipada (comentado por não ser usado)
	// ------------------------------
	// const x int = 10

	// ------------------------------
	// 📌 Example07 – Constante tipada vs untyped
	// ------------------------------
	// const a int = 10
	// takeInt32(a) // erro

	// const b = 10
	// takeInt32(b) // funciona

	// ------------------------------
	// 📌 Example08 – Constante literal
	// ------------------------------
	// takeInt32(10)
}
