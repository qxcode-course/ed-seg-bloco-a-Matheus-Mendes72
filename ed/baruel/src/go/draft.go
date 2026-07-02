package main
import "fmt"
func main() {
    var qtd_total, qtd_fig int
    fmt.Scan(&qtd_total, &qtd_fig)

    baruel := make([]int, qtd_total+1)
    for i := 0; i < qtd_fig; i++ {
        var freq int
        fmt.Scan(&freq)
        baruel[freq]++
    }

    primeira := true
    for i := 1; i <= qtd_total; i++ {
        for j := 1; j < baruel[i]; j++ {
            if !primeira {
                fmt.Print(" ")
            }
            fmt.Print(i)
            primeira = false
        }
    }

    if primeira {
        fmt.Print("N")
    }
    fmt.Println()
    primeira = true

    for i := 1; i <= qtd_total; i++ {
        if baruel[i] == 0 {
            if !primeira {
                fmt.Print(" ")
            }
            fmt.Print(i)
            primeira = false
        }
    }

    if primeira {
        fmt.Print("N")
    }
    fmt.Println()
}