package main
import "fmt"

func div(numero int) {
    divisao := numero / 2
    resto := numero % 2

    if divisao == 0 {
        fmt.Println(divisao, resto)
        return
    }

    div(divisao)
    fmt.Println(divisao, resto)
}

func main() {
    var numero int
    fmt.Scan(&numero)

    div(numero)
}
