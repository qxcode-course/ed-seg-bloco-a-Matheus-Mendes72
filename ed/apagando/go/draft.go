package main
import "fmt"
func main() {
	var pessoasFila, deixaramFila int
	fmt.Scan(&pessoasFila)

	fila := make([]int, pessoasFila)
	for i := 0; i < pessoasFila; i++ {
		fmt.Scan(&fila[i])
	}

	fmt.Scan(&deixaramFila)

	filaSaida := make([]int, deixaramFila)
	for i := 0; i < deixaramFila; i++ {
		fmt.Scan(&filaSaida[i])
	}

	fmt.Println(pessoasFila, fila, deixaramFila, filaSaida)
}
