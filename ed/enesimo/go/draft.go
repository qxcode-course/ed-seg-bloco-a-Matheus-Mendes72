package main
import "fmt"

func eh_primo(numero, divisor int) bool {
    if numero <= 1 {
        return false
    }

    if divisor == numero {
        return true
    }

    if numero % divisor == 0 {
        return false
    }

    return eh_primo(numero, divisor+1)
}

func nPrimo(numero, contador, atual int) int {
    if contador == numero {
        return atual-1
    }

    if eh_primo(atual, 2) {
        return nPrimo(numero, contador+1, atual+1)
    } 
        
    return nPrimo(numero, contador, atual+1)
}

func main() {
    var numero int
    fmt.Scan(&numero)

    fmt.Println(nPrimo(numero, 0, 2))
}
