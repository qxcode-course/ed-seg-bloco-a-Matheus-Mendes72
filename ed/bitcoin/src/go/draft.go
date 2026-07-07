package main
import "fmt"

func ativos(n, k int) int {
    if n <= k {
        return 1
    }

    return ativos(n/2, k) + ativos(n/2, k)
}

func main() {
    var n, k int
    fmt.Scan(&n, &k)

    fmt.Println(ativos(n, k))
}