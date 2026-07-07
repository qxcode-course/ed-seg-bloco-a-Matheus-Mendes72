package main
import "fmt"

func quadrado(n int) int {
    if n == 1 {
        fmt.Println("1^2 = 1")
        return 1
    }

    fmt.Printf("%d^2 = %d^2 + 2*%d + 1 = ?\n", n, (n-1), (n-1))

    // quadrado_anterior := quadrado(n-1)
    resultado := quadrado(n-1) + 2*(n-1) + 1
    fmt.Printf("%d^2 = %d^2 + 2*%d + 1 = %d\n", n, (n-1), (n-1), resultado)

    return resultado
}

func main() {
    var n int
    fmt.Scan(&n)

    quadrado(n)
}
