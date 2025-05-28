package main

import (
	"fmt"
	"math"
	"time"
)

// 📌 Example01 – Switch básico
func doBasic(x int) {
	switch x {
	case 1:
		fmt.Println("x é 1")
	case 2:
		fmt.Println("x é 2")
	default:
		fmt.Println("x não é 1 nem 2")
	}
}

// 📌 Example02 – Fallthrough
func doFallthrough(x int) {
	switch x {
	case 1:
		fmt.Println("x é 1")
		fallthrough
	case 2:
		fmt.Println("x é 2")
	default:
		fmt.Println("x não é 1 nem 2")
	}
}

// 📌 Example03 – Switch sem expressão
func doWithCondition(x int) {
	switch {
	case x == 1:
		fmt.Println("x é 1")
	case "abc" == "foo":
		fmt.Println("x é foo")
	default:
		fmt.Println("nenhuma condição satisfeita")
	}
}

// 📌 Example04 – Switch sem expressão com time
func isWeekday(t time.Time) bool {
	switch {
	case t.Weekday() > 0 && t.Weekday() < 6:
		return true
	default:
		return false
	}
}

// 📌 Example05 – Switch com declaração de variável
func switchWithDeclaration() {
	switch x := math.Sqrt(4); x {
	case 2:
		fmt.Println("resultado é 2")
	default:
		fmt.Println("algo deu errado")
	}
}

// 📌 Example06 – Múltiplos valores no mesmo case
func isWeekend(t time.Time) bool {
	switch t.Weekday() {
	case time.Saturday, time.Sunday:
		return true
	default:
		return false
	}
}

// 📌 Example07 – Type switch
func typeSwitch(x any) {
	switch x.(type) {
	case string:
		fmt.Println("tipo string")
	case int:
		fmt.Println("tipo int")
	case nil:
		fmt.Println("tipo nil")
	default:
		fmt.Println("outro tipo")
	}
}

// 📌 Example08 – Type switch com atribuição
func typeSwitchWithAssign(x any) {
	switch t := x.(type) {
	case string:
		takeString(t)
	case int:
		fmt.Println("int")
	case nil:
		fmt.Println("nil")
	default:
		fmt.Println("outro tipo")
	}
}

func takeString(s string) {
	fmt.Println("String recebida:", s)
}

// 🧪 Execução dos exemplos
func main() {
	fmt.Println("Exemplos de switch em Go:")

	// doBasic(1)
	// doBasic(3)

	// doFallthrough(1)

	// doWithCondition(1)

	// fmt.Println("É dia útil?", isWeekday(time.Now()))

	// switchWithDeclaration()

	// fmt.Println("É final de semana?", isWeekend(time.Now()))

	// typeSwitch("abc")
	// typeSwitch(42)
	// typeSwitch(nil)

	// typeSwitchWithAssign("abc")
}
