package main
import "fmt"

func espada(jogo []int, escolhido int) {
    fmt.Print("[ ")
    for i, valor := range jogo {
        if i == escolhido {
            if valor > 0 {
                fmt.Printf("%v> ", valor)
            } else {
                fmt.Printf("<%v ", valor)
            }
        } else {
            fmt.Printf("%v ", valor)     
        }
    }
    fmt.Println("]")
}

func dinamismo(jogo []int, escolhido int) {
    escolhido--
    var remover int

    for len(jogo) > 0 {
        espada(jogo, escolhido)
            
        if len(jogo) == 1 {
        break
        }

        if jogo[escolhido] > 0 {
            remover = (escolhido +1) % len(jogo)

            jogo = append(jogo[:remover], jogo[remover+1:]...)

            escolhido = remover % len(jogo)
        } else {
            remover = (escolhido -1 + len(jogo)) % len(jogo)
            
            jogo = append(jogo[:remover], jogo[remover+1:]...)
            
            escolhido = (remover -1 + len(jogo)) % len(jogo)
            //jogo = append(jogo[:remover+1], jogo[remover:]...)
        }
        
    }
}

func main() {
    var qtd, escolhido, fase int
    fmt.Scan(&qtd, &escolhido, &fase)

    jogo := make([]int, qtd)
    for i := 0; i < qtd; i++ {
        valor := i +1;

        if fase == 1 {
            if valor%2 == 0 {
                jogo[i] = -valor
            } else {
                jogo[i] = valor
            }
        } else {
            if valor%2 == 1 {
                jogo[i] = -valor
            } else {
                jogo[i] = valor
            }
        }
    }
    
    dinamismo(jogo, escolhido)
}
