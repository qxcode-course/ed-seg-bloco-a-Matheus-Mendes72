package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string) bool {
	linha := len(grid)
	coluna := len(grid[0])

	if len(word) == 0 || len(grid) == 0 {
		return false
	}

	var buscar func(int, int, int) bool

	buscar = func(i int, j int, indice int) bool {
		if (indice == len(word)) {
			return true
		}

		if i < 0 || i >= linha || j < 0 || j >= coluna {
			return false
		}

		if grid[i][j] != word[indice] {
			return false
		}

		aux := grid[i][j]

		grid[i][j] = '*'

		if buscar(i+1, j, indice+1) || buscar(i, j+1, indice+1) || buscar(i-1, j, indice+1) || buscar(i, j-1, indice+1) {
			return true
		}

		grid[i][j] = aux
		
		return false
	}

	for i := 0; i < linha; i++ {
		for j := 0; j < coluna; j++ {
			if buscar(i, j, 0) {
				return true
			}
		}
	}

	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
