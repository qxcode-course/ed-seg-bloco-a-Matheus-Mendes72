// for len(jogo) > 0 {
//         espada(jogo, escolhido)

//         if escolhido +1 == len(jogo) {
//             escolhido = escolhido % len(jogo)
//         }
//         if escolhido +2 <= len(jogo) {
//             jogo = append(jogo[:escolhido], jogo[escolhido+1:]...)
//         }

//         if len(jogo) == 0 {
//             break
//         }
//     }
package main
import "fmt"
// import "strconv"
// import "strings"

func espada(jogo []int, escolhido int) {
    fmt.Print("[ ")
    for i, valor := range jogo {
        if i == escolhido {
            fmt.Printf("%v> ", valor)
        } else {
            fmt.Printf("%v ", valor)
        }
    }
    fmt.Println("]")
}

func dinamismo(jogo []int, escolhido int) {
    escolhido--
    // remover := escolhido

    for len(jogo) > 0 {
        espada(jogo, escolhido)
        
        if len(jogo) == 1 {
            break
        }

        remover := (escolhido +1) % len(jogo)

        jogo = append(jogo[:remover], jogo[remover+1:]...)

        escolhido = remover % len(jogo)
    }
}

func main() {
    var qtd, escolhido int
    fmt.Scan(&qtd, &escolhido)

    jogo := make([]int, qtd)
    for i := 0; i < qtd; i++ {
        jogo[i] = i +1;
    }
    
    dinamismo(jogo, escolhido)

}

// var aux string
// for _, valor := range jogo {
//     aux += strconv.Itoa(valor)
// }

// for i, valor := range jogo {
//     if i+1 == escolhido {
//         aux = strings.Join(strconv.Itoa(aux), ">")
//     }
// }