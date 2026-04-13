package main
import "fmt"
func main() {
	var qtdPessoas, qtdPessoas_deixaramFila int
	fmt.Scan(&qtdPessoas)

	fila := make([]int, qtdPessoas)
	for i := 0; i < qtdPessoas; i++ {
		fmt.Scan(&fila[i])
	}

	fmt.Scan(&qtdPessoas_deixaramFila)

	pessoasImpacientes := make(map[int]bool)
	for i := 0; i < qtdPessoas_deixaramFila; i++ {
		indice := 0
		fmt.Scan(&indice)
		pessoasImpacientes[indice] = true
	}

	for _, valor := range fila {
		impaciente, _ := pessoasImpacientes[valor]

		if impaciente == false {
			fmt.Printf("%v ", valor)
		}
	}

	fmt.Println()
}
