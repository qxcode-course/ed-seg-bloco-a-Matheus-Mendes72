package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}

	linhas := len(matrix)
	colunas := len(matrix[0])

	memo := make([][]int, linhas)
	for i := range memo {
		memo[i] = make([]int, colunas)
	}

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if memo[i][j] != 0 {
			return memo[i][j]
		}

		maior := 1

		direcoes := [][]int {
			{-1, 0},
			{1, 0},
			{0, -1},
			{0, 1},
		}

		for _, direcao := range direcoes {
			ni := i + direcao[0]
			nj := j + direcao[1]

			if ni >= 0 && ni < linhas && nj >= 0 && nj < colunas && matrix[i][j] < matrix[ni][nj] {
				caminho := 1 + dfs(ni, nj)

				if caminho > maior {
					maior = caminho
				}
			}
		}

		memo[i][j] = maior
		return maior
	}

	resposta := 0

	for i := 0; i < linhas; i++ {
		for j := 0; j < colunas; j++ {
			caminho := dfs(i, j)

			if caminho > resposta {
				resposta = caminho
			}
		}
	}

	return resposta
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}
	parts := strings.Fields(scanner.Text())
	if len(parts) < 2 {
		return
	}
	nl, _ := strconv.Atoi(parts[0])
	nc, _ := strconv.Atoi(parts[1])

	matrix := make([][]int, nl)
	for i := 0; i < nl; i++ {
		if !scanner.Scan() {
			return
		}
		tokens := strings.Fields(scanner.Text())
		row := make([]int, nc)
		for j := 0; j < nc && j < len(tokens); j++ {
			v, _ := strconv.Atoi(tokens[j])
			row[j] = v
		}
		matrix[i] = row
	}

	fmt.Println(longestIncreasingPath(matrix))
}
