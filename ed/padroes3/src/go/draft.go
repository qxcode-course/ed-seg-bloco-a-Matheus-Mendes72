package main
import "fmt"

func padrao(n, m int) int {
    if m == 1 {
        return n
    }
    
    return padrao(n, m-1) + (n-2)
}

func main() {
    var n, m int
    fmt.Scan(&n, &m)

    fmt.Println(padrao(n, m))
}
