package main
import "fmt"

func doacao(n, c float64) float64 {
    
    /*
    x reais
    2x
    2x - c
    2*(2x-c)
    2*(2*(2x-c)) = 4 * (2x-c)

    na penultima doacao ele tem o valor de c em dinheiro
    se n = 3 e c = 20
    n = 2; x = 10 = 30 - 20
    n = 1; x = 30 / 2 = 15 = 35 - 20 
    n = 0; 35 / 2 = 17.50

    F(0) = (F(1) + C) / 2; F(1) = (15 + 20) 
    */

    if n == 0 {
        return 0
    }

    return (doacao(n-1, c) + c) / 2
}

func main() {
    var n, c float64
    fmt.Scan(&n, &c)

    dinheiro := doacao(n, c)

    fmt.Printf("%.2f\n", dinheiro)
}