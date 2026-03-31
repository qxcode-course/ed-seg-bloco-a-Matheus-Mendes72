package main
import "fmt"
func main() {
    var tamCons, tamBusc int
    fmt.Scan(&tamCons)

    vetCons := make([]string, tamCons)
    for i := range vetCons {
        fmt.Scan(&vetCons[i])
    }

    fmt.Scan(&tamBusc)

    vet := make([]string, tamBusc)
    vetBusc := make(map[string]int)
    for i := range vet {
        fmt.Scan(&vet[i])
        vetBusc[vet[i]] = 0
    }

    // for i := 0; i < tamBusc; i++ {
    //     var s string
    //     fmt.Scan(&s)
    //     vetBusc[s] = 0
    // }
    
    //contador := make([]int, tamBusc)
    for _, valor := range vetCons {
        _, existe := vetBusc[valor]
        if existe {
            vetBusc[valor] += 1
        }
    }
    //fmt.Println(vetBusc)

    //contador := 0
    for i := 0; i < len(vet); i++ {
        if i != len(vet) -1 {
            fmt.Print(vetBusc[vet[i]], " ")
        } else {
            fmt.Println(vetBusc[vet[i]])
        }
        //contador++
    }

    // fmt.Println()

    //fmt.Println(vetCons, vetBusc)
}
