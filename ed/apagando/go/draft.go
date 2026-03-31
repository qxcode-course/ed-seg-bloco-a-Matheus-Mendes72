package main
import "fmt"
func main() {
	var pessoasFila, impaciente int
	fmt.Scan(&pessoasFila)

	fila := make([]int, pessoasFila)
	for i := 0; i < pessoasFila; i++ {
		fmt.Scan(&fila[i])
	}

	fmt.Scan(&impaciente)

	impacientes := make(map[int]bool)
	for i := 0; i < impaciente; i++ {
		var indice int
		fmt.Scan(&indice)
		impacientes[indice] = true
	}

	for _, valor := range fila {
		v, _ := impacientes[valor]
		if v {
			continue
		} else {
			fmt.Printf("%v ", valor)
		}
	}
	fmt.Println()
}
