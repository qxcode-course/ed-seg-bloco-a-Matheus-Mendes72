package main
import "fmt"

type Pos struct {
    l, c int
}

func inside(grid [][]int, p Pos) bool {
    linhas := len(grid)
    colunas := len(grid[0])

    return p.l >= 0 && p.l < linhas && p.c >= 0 && p.c < colunas
}

func getNeig(p Pos) []Pos {
    return []Pos{
        {p.l+1, p.c},
        {p.l-1, p.c},
        {p.l, p.c+1},
        {p.l, p.c-1},
    }
}

func laranjas(grid [][]int) int {
    fila := []Pos{}
    frescas := 0

    for i := range grid {
        for j := range grid[i] {
            if grid[i][j] == 2 {
                fila = append(fila, Pos{i, j})
            }

            if grid[i][j] == 1 {
                frescas++
            }
        }
    }

    if frescas == 0 {
        return 0
    }

    minutos := 0

    for len(fila) > 0 && frescas > 0 {
        for i := 0; i < len(fila); i++ {
            atual := fila[0]
            fila = fila[1:]

            for _, viz := range getNeig(atual) {
                if !inside(grid, viz) {
                    continue
                }

                if grid[viz.l][viz.c] == 1 {
                    grid[viz.l][viz.c] = 2
                    frescas--
                    fila = append(fila, viz)
                    // fmt.Println(fila)
                }
            }
        }
        minutos++
    }

    if frescas > 0 {
        return -1
    }

    return minutos
}

func main() {
    var l, c int
    fmt.Scan(&l, &c)

    grid := make([][]int, l)
    for i := 0; i < l; i++ {
        grid[i] = make([]int, c)
        
        for j := 0; j < c; j++ {
            fmt.Scan(&grid[i][j])
        }
    }

    fmt.Println(laranjas(grid))
}