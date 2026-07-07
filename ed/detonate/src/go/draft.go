package main
import "fmt"

func dfs(graph [][]int, inicio int, visitado []bool) int {
    visitado[inicio] = true

    qtd_bomba := 1

    for _, ref := range graph[inicio] {
        if !visitado[ref] {
            qtd_bomba += dfs(graph, ref, visitado)
        }
    }

    return qtd_bomba
}

func maximumDetonation(bombs [][]int) int {
    n := len(bombs)

    graph := make([][]int, n)
    for i := 0; i < n; i++ {
        x1 := bombs[i][0]
        y1 := bombs[i][1]
        raio := bombs[i][2]
        for j := 0; j < n; j++ {
            if i == j {
                continue
            }

            x2 := bombs[j][0]
            y2 := bombs[j][1]

            distx := x1 - x2
            disty := y1 - y2

            // Fórmula de distancia entre dois pontos
            dist := distx*distx + disty*disty

            if dist <= raio*raio {
                graph[i] = append(graph[i], j)
            }
        }
    }

    resposta := 0
    for inicio := 0; inicio < n; inicio++ {
        visitado := make([]bool, n)

        qtd_bomba := dfs(graph, inicio, visitado)

        if qtd_bomba > resposta {
            resposta = qtd_bomba
        }
    }

    return resposta
}

func main() {
    var n, m int
    fmt.Scan(&n, &m)

    bombs := make([][]int, n)
    for i := 0; i < n; i++ {
        bombs[i] = make([]int, m)
        fmt.Scan(&bombs[i][0], &bombs[i][1], &bombs[i][2])
    }

    fmt.Println(maximumDetonation(bombs))
}