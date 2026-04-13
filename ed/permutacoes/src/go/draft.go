package main
import "fmt"
import "sort"

func backtrack(pos, tamanho int, ref string, usado []bool, atual []rune, resultado []string) []string {
    if (pos == tamanho) {
        return append(resultado, string(atual))
    }

    for i := 0; i < tamanho; i++ {
        if (!usado[i]) {
            usado[i] = true
            atual[pos] = rune(ref[i])
            
            resultado = backtrack(pos+1, tamanho, ref, usado, atual, resultado)
            usado[i] = false

        }
    }

    return resultado
}

func main() {
    var ref string
    fmt.Scan(&ref)

    tamanho := len(ref)
    usado := make([]bool, tamanho)
    atual := make([]rune, tamanho)
    var resultado []string

    resultado = backtrack(0, tamanho, ref, usado, atual, resultado)

    sort.Strings(resultado)

    for _, res := range resultado {
        fmt.Println(res)
    }
}
