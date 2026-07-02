package main
import "fmt"
func main() {
    var N int
    fmt.Scan(&N)

    bomba := make([]int, N)
    distancia := make([]int, N)

    for i := 0; i < N; i++ {
        fmt.Scan(&bomba[i], &distancia[i])
    }

    head := 0
    current := 0
    for i := 0; i < N; i++ {
        diferenca := bomba[i] - distancia[i]
        current += diferenca

        if current < 0 {
            current = 0
            head = i + 1
        }
    }

    fmt.Println(head)
}