package main
import "fmt"

func subset_sum(numeros []int, pos int, soma int, alvo int) bool {
    if soma == alvo {
        return true
    }

    if pos == len(numeros) {
        return false
    }

    if subset_sum(numeros, pos+1, soma+numeros[pos], alvo) {
        return true
    }

    if subset_sum(numeros, pos+1, soma, alvo) {
        return true
    }

    return false
}

func main() {
    var n, alvo int
    fmt.Scan(&n, &alvo)

    numeros := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&numeros[i])
    }

    if subset_sum(numeros, 0, 0, alvo) {
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }
}