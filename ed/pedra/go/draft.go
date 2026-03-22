package main
import "fmt"
//import "math"

type jogada struct {
    pa int 
    pb int
}

func abs(diferenca int) int {
    if (diferenca < 0) {
        return -diferenca
    }
    return diferenca
}

func ganhador(jogadas []jogada) int {
    indice := -1
    ganhador := 1000

    for i, valor := range jogadas {
        if (valor.pa < 10 || valor.pb < 10) {
            continue
        }

        diferenca := abs(valor.pa - valor.pb)
        if (diferenca < ganhador) {
            ganhador = diferenca
            indice = i
        }
    }
    return indice
}

func main() {
    var n int
    fmt.Scan(&n)

    jogadas := make([]jogada, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&jogadas[i].pa, &jogadas[i].pb)
    }

    indice := ganhador(jogadas)
    if (indice == -1) {
        fmt.Println("sem ganhador")
    } else {
        fmt.Println(indice)
    }
}





// type jogada struct {
//     pa int
//     pb int
// }

// func abs(diferenca int) int {
//     if diferenca < 0 {
//         return -diferenca
//     }
//     return diferenca
// }

// func verificacao(jogadas []jogada) int {
//     indice := -1
//     ganhador := 1000

//     for i, valor := range jogadas {
//         if valor.pa < 10 || valor.pb < 10 {
//             continue
//         }
//         pontuacao := abs(valor.pa - valor.pb)
//         if pontuacao < ganhador {
//             ganhador = pontuacao
//             indice = i
//         }
//     }
//     return indice
// }

// func main() {
//     var n int
//     fmt.Scan(&n)

//     jogadas := make([]jogada, n)
//     for i := 0; i < n; i++ {
//         fmt.Scan(&jogadas[i].pa, &jogadas[i].pb)
//     }

//     indice := verificacao(jogadas)
//     if indice == -1 {
//         fmt.Println("sem ganhador")
//     } else {
//         fmt.Println(indice)
//     }
// }





// func valido(p1, p2 int) bool {
//     return p1 >= 10 && p2 >= 10
// }

// func diferenca(p1, p2 int) int {
//     if p1 > p2 {
//         return p1 - p2
//     } else {
//         return p2 - p1
//     }
// }

// func vencedor(n int) int {
//     var p1, p2 int
//     indice := -1
//     ganhador := 1000
//     for i := range n {
//         fmt.Scan(&p1, &p2)
        
//         val := valido(p1, p2)
//         if !val {
//             continue
//         }
        
//         pontuacao := diferenca(p1, p2)
//         if pontuacao < ganhador {
//             ganhador = pontuacao
//             indice = i
//         }
//     }
//     return indice
// }

// func main() {
//     var n int
//     fmt.Scan(&n)
    
//     indice := vencedor(n)
//     if (indice == -1) {
//         fmt.Println("sem ganhador")
//     } else {
//         fmt.Println(indice)
//     }
// } 