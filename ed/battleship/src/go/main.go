package main

import (
	"bufio"
	"fmt"
	"os"
)

// Função que será chamada no LeetCode
func countBattleships(board [][]byte) int {
	if len(board) == 0 {
		return 0
	}

	linhas := len(board)
	colunas := len(board[0])

	qtd_navios := 0
	var dfs func(int, int)
	dfs = func(i, j int) {
		if i < 0 || i >= linhas || j < 0 || j >= colunas {
			return 
		}

		if board[i][j] != 'X' {
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
			if board[i][j] == 'X' {
				qtd_navios++
				dfs(i, j)
			}
		}
	}
	
	return qtd_navios
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)

	board := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}

	result := countBattleships(board)
	fmt.Println(result)
}
