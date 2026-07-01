package main

import (
	"bufio"
	"fmt"
	"os"
)

// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}

	linhas := len(board)
	colunas := len(board[0])

	var dfs func(int, int)
	dfs = func(i, j int) {
		if i < 0 || i >= linhas || j < 0 || j >= linhas {
			return
		}

		if board[i][j] != 'O' {
			return
		}

		board[i][j] = 'M'

		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}

	for i := 0; i < linhas; i++ {
		for j := 0; j < colunas; j++ {
			if i == 0 {
				dfs(i, j)
				dfs(linhas-1, j)
			}

			if j == 0 {
				dfs(i, j)
				dfs(i, colunas-1)
			}
		}
	}

	for i := 0; i < linhas; i++ {
		for j := 0; j < colunas; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}

			if board[i][j] == 'M' {
				board[i][j] = 'O'
			}
		}
	}
}

// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}
