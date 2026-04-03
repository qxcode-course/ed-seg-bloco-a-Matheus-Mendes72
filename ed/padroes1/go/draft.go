package main
import "fmt"

func padrao(n, soma int) int {
    if n <= 1 {
        return soma
    }

    return padrao(n-1, soma+8)
}

func main() {
    var n int
    fmt.Scan(&n)

    fmt.Println(padrao(n, 20))
}
