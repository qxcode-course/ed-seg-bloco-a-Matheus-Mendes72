package main
import "fmt"
// import "math"

// type Jogada struct {
//     pa, pb int
// }

// func main() {
//     qtd := 0
//     fmt.Scan(&qtd)
//     jogadas := make([]Jogada, qtd)

//     for i := range jogadas {
//         jog := &jogadas[i]
//         fmt.Scan(&jog.pa, &jog.pb)
//     }

//     ind_melhor := -1
//     valor_melhor := math.MaxInt32
//     for i, jog := range jogadas {
//         if jog.pa < 10 || jog.pb < 10 {
//             continue
//         }

//         pontuacao := jog.pa - jog.pb
//         if pontuacao < 0 {
//             pontuacao = -pontuacao
//         }

//         if pontuacao < valor_melhor {
//             ind_melhor = i
//             valor_melhor = pontuacao
//         } 
         
//         if ind_melhor == -1 {
//             fmt.Println("sem ganhador")
//         }
//     }
//     fmt.Println(ind_melhor)
// }





func valido(p1, p2 int) bool {
    return p1 >= 10 && p2 >= 10
}

func diferenca(p1, p2 int) int {
    if p1 > p2 {
        return p1 - p2
    } else {
        return p2 - p1
    }
}

func vencedor(n int) int {
    var p1, p2 int
    indice := -1
    ganhador := 1000
    for i := range n {
        fmt.Scan(&p1, &p2)
        
        val := valido(p1, p2)
        if !val {
            continue
        }
        
        pontuacao := diferenca(p1, p2)
        if pontuacao < ganhador {
            ganhador = pontuacao
            indice = i
        }
    }
    return indice
}

func main() {
    var n int
    fmt.Scan(&n)
    
    indice := vencedor(n)
    if (indice == -1) {
        fmt.Println("sem ganhador")
    } else {
        fmt.Println(indice)
    }
} 