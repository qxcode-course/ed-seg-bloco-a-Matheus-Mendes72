package main
import "fmt"

type gomos struct {
    x int
    y int
}

func movimento(direcao string) (int, int) {
    if direcao == "L" {
        return -1, 0
    } 
    if direcao == "R" {
        return 1, 0
    }
    if direcao == "U" {
        return 0, -1
    }
    
    return 0, 1
}

func main() {
    var qtd int
    var direcao string
    fmt.Scan(&qtd, &direcao)
    

    cobra := make([]gomos, qtd)
    for i := 0; i < qtd; i++ {
        fmt.Scan(&cobra[i].x, &cobra[i].y)
    }

    dirX, dirY := movimento(direcao)

    for i := len(cobra) -1; i > 0; i-- {
        cobra[i] = cobra[i-1]
    }

    cobra[0].x += dirX
    cobra[0].y += dirY

    for _, valor := range cobra {
        fmt.Println(valor.x, valor.y)
    }
}
