package main
import "fmt"
import "strconv"

func tostr(vet []int) string {
    if len(vet) == 0 {
        return ""
    }

    if len(vet) == 1 {
        return strconv.Itoa(vet[0])
    }

    return strconv.Itoa(vet[0]) + ", " + tostr(vet[1:])
}

func eh_primo(numero, divisor int) bool {
    if numero <= 1 {
        return false
    }

    if numero == divisor {
        return true
    }

    if numero % divisor == 0 {
        return false
    }

    return eh_primo(numero, divisor+1)
}

func numPrimos(n, contador, atual int, vet []int) string {
    if n == 1 {
        vet = append(vet, 2)
        return tostr(vet)
    }

    if contador == n {
        return tostr(vet)
    }

    if eh_primo(atual, 2) {
        vet = append(vet, atual)
        return numPrimos(n, contador+1, atual+1, vet)
    } else {
        return numPrimos(n, contador, atual+1, vet)
    }

    return tostr(vet)
}

func main() {
    var n int
    fmt.Scan(&n)

    vet := make([]int, 0)
    fmt.Println("[" + numPrimos(n, 0, 2, vet) + "]")
}
