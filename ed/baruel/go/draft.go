package main
import "fmt"
func main() {
    var qtd_total, baruel int
    fmt.Scan(&qtd_total, &baruel)

    figurinhas := make(map[int]int)
    for i := 0; i < baruel; i++ {
        var val_figur int
        fmt.Scan(&val_figur)
        figurinhas[val_figur]++
    }

    fig_rep := false
    prim := true
    for i := 0; i < baruel; i++ {
        if figurinhas[i] > 1 {
            for j := 0; j < figurinhas[i]-1; j++ {
                    if !prim {
                        fmt.Print(" ")
                    }
                fmt.Print(i)
                fig_rep = true
                prim = false
            }
        }
    }

    if !fig_rep {
        fmt.Print("N")
    }

    fmt.Println()

    prim = true
    faltando := false
    for i := 1; i <= qtd_total; i++ {
        if figurinhas[i] == 0 {
            if !prim {
                fmt.Print(" ")
            }
            fmt.Print(i)
            faltando = true
            prim = false
        }
    }

    if !faltando {
        fmt.Print("N")
    }

    fmt.Println()
}
