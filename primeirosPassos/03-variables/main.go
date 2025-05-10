package main

import "fmt"

// Variável global (pode ser declarada sem ser usada)
var idade int = 30

func main() {
	// 1. Declaração com tipo explícito (sem inicialização)
	var nome string = "Smilly"
	var sobrenome string = "Delmondes"
	fmt.Println("1. Nome:", nome, "| Sobrenome:", sobrenome, "| Idade:", idade)

	// 2. Declaração com inferência de tipo (inicialização)
	var cidade, estado = "São Paulo", "SP"
	fmt.Println("2. Cidade:", cidade, "| Estado:", estado)

	// 3. Agrupamento de variáveis
	var (
		pais       = "Brasil"
		continente = "América do Sul"
	)
	fmt.Println("3. País:", pais, "| Continente:", continente)

	// 4. Short syntax (somente dentro da função)
	profissao := "Desenvolvedor"
	empresa := "TechCorp"
	fmt.Println("4. Profissão:", profissao, "| Empresa:", empresa)

	// 5. Declaração mista com tipos diferentes
	var linguagem, ano = "Go", 2009
	fmt.Println("5. Linguagem:", linguagem, "| Ano:", ano)

	// 6. Short syntax com múltiplos valores
	linguagem2, ano2 := "Python", 1991
	fmt.Println("6. Linguagem2:", linguagem2, "| Ano2:", ano2)

	// 7. Demonstrando valores zero
	var (
		numero int
		texto  string
		ativo  bool
	)
	fmt.Println("7. Valores zerados -> Número:", numero, "| Texto:", texto, "| Ativo:", ativo)
}
