package main
import "fmt"

func quadrado(n int) int {
    if n == 1 {
        return 1
    }

    fmt.Printf("%d^2 = %d^2 + 2*%d + 1 = ?", n, (n-1), (n-1))
}

func main() {
    var n int
    fmt.Scan(&n)

    quadrado(n)
}
