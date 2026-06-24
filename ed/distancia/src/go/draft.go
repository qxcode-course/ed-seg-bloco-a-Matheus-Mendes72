package main
import "fmt"

var (
    seq string
    L int
)

func distancia(posicao int, texto []rune, ultima[]int) bool {
    if posicao == len(texto) {
        return true
    }

    if texto[posicao] != '.' {
        num := int(texto[posicao] - '0')

        if posicao-ultima[num] <= L {
            return false
        }

        antiga := ultima[num]
        ultima[num] = posicao

        if distancia(posicao+1, texto, ultima) {
            return true
        }
        ultima[num] = antiga

        return false
    }

    for i := 0; i <= L; i++ {
        if posicao-ultima[i] > L {
            texto[posicao] = rune(i + '0')
            
            antiga := ultima[i]
            ultima[i] = posicao

            if distancia(posicao+1, texto, ultima) {
                return true
            }

            ultima[i] = antiga
            texto[posicao] = '.'
        }
    }

    return false
}

func main() {
    fmt.Scan(&seq, &L) 
    
    palavra := []rune(seq)
    ultima := make([]int, L+1)

    for i := 0; i <= L; i++ {
        ultima[i] = -1000
    }

    distancia(0, palavra, ultima)

    fmt.Println(string(palavra))
}
