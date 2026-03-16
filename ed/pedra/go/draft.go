package main
import "fmt"
import "math"
func main() {
    var n, p1, p2 int
    fmt.Scan(&n)

    indice := -1
    ganhador := 1000

    for i := range n {
        fmt.Scan(&p1, &p2)

        if (p1 < 10 || p2 < 10) {
            continue
        }
        diferenca := int(math.Abs(float64(p1 - p2)))
        if (diferenca < ganhador) {
            ganhador = diferenca
            indice = i
        }
    }

    if (indice == -1) {
        fmt.Println("sem ganhador")
    } else {
        fmt.Println(indice)
    }
}
