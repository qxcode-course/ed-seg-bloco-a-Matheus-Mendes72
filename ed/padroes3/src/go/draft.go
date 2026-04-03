package main
import "fmt"

func padrao(n, m, soma int) int {
    if m == 0 {
        return soma
    }

    soma = 
    
    return padrao(n-1, m-1, soma)
}

func main() {
    var n, m int
    fmt.Scan(&n, &m)

    fmt.Println(padrao(n, m, 1))
}
