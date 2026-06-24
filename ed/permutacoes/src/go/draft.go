package main
import "fmt"
import "sort"

func permutacao(posicao int, palavra string, usado []bool, atual []rune, resultado []string) []string {
    if posicao == len(palavra) {
        return append(resultado, string(atual))
    }

    for i := 0; i < len(palavra); i++ {
        if !usado[i] {
            usado[i] = true
            atual[posicao] = rune(palavra[i])

            resultado = permutacao(posicao+1, palavra, usado, atual, resultado)

            usado[i] = false        
        }
    }

    return resultado
}

func main() {
    var palavra string
    fmt.Scan(&palavra)

    usado := make([]bool, len(palavra))
    atual := make([]rune, len(palavra))

    resultado := []string{}

    resultado = permutacao(0, palavra, usado, atual, resultado)

    sort.Strings(resultado)

    for _, i := range resultado {
        fmt.Println(i)
    }
}