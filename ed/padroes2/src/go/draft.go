package main
import "fmt"

func padrao(n, soma, contador int) int {
    if n <= 1 {
        return soma
    }

    soma = contador*contador + 2*contador
    
    return padrao(n-1, soma, contador+1)
}

func main() {
    var n int
    fmt.Scan(&n)

    fmt.Println(padrao(n, 3, 2))
}
