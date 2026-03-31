package main
import "fmt"
func main() {
    var tamanho, elemDesl int
    fmt.Scan(&tamanho, &elemDesl) 

    vetor := make([]int, tamanho)
    for i := 0; i < tamanho; i++ {
        fmt.Scan(&vetor[i])
    }

    vetAux := make([]int, tamanho)
    elemDesl = elemDesl % tamanho

    for i := 0; i < tamanho; i++ {
        novaPos := (i + elemDesl) % tamanho
        vetAux[novaPos] = vetor[i]
    }
    
    fmt.Print("[ ")
    for i := 0; i < tamanho; i++ {
        fmt.Print(vetAux[i], " ")
    }
    fmt.Println("]")
}
