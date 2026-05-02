package main

import (
	"bufio"
    "fmt"
    "os"
)

type Pos struct {
	l, c int
}

func search(grid[][]rune, startPos, endPos Pos) bool {
	l, c := startPos.l, startPos.c 

	if (l < 0 || l >= len(grid) || (c < 0 || c >= len(grid[0]))) {
		return false
	}

	if grid[l][c] != ' ' {
		return false
	}

	stack := NewStack[Pos]()
	stack.Push(startPos)

	for !stack.IsEmpty() {
		pos := stack.Pop()

		if (pos.l < 0 || pos.l >= len(grid)) || (pos.c < 0 || pos.c >= len(grid[0])) {
			return false
		}

		if grid[pos.l][pos.c] == 'I' || grid[pos.l][pos.c] == 'F' {
			grid[pos.l][pos.c] = '.'
			continue
		}

		if grid[pos.l][pos.c] != ' ' {
			continue
		}

		grid[pos.l][pos.c] = '.'

		if pos.l == endPos.l && pos.c == endPos.c {
			return false
		}

		stack.Push(Pos{pos.l+1, pos.c})
		stack.Push(Pos{pos.l, pos.c+1})
		stack.Push(Pos{pos.l-1, pos.c})
		stack.Push(Pos{pos.l, pos.c-1})
	}

	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nl_nc := scanner.Text()
	var nl, nc int
	fmt.Sscanf(nl_nc, "%d %d", &nl, &nc)
	grid := make([][]rune, nl)

	// Lê a gridriz
	for i := range nl {
		scanner.Scan()
		grid[i] = []rune(scanner.Text())
	}

	// Procura posições de início e endPos e conserta para _
	var startPos, endPos Pos
	for l := range nl {
		for c := range nc {
			if grid[l][c] == 'I' {
				grid[l][c] = ' '
				startPos = Pos{l, c}
			}
			if grid[l][c] == 'F' {
				grid[l][c] = ' '
				endPos = Pos{l, c}
			}
		}
	}

	search(grid, startPos, endPos)

	for _, line := range grid {
		fmt.Println(string(line))
	}
}