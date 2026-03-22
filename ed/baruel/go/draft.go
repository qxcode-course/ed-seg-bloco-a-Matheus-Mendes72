package main
import "fmt"
// func main() {
//     qtd_album := 0
//     qtd_fig := 0
//     fmt.Scan(&qtd_album, &qtd_fig)
//     album := make([]int, qtd_fig)
//     unicos := make(map[int]bool)
//     repetidos := make([]int, 0, qtd_fig)

//     for i := range album {
//         fmt.Scan(&album[i])
//     }

//     for _, fig := range album {
//         if unicos[fig] {
//             repetidos = append(repetidos, fig)
//         } else {
//             unicos[fig] = true
//         }
//     }

//     if len(repetidos) == 0 {
//         fmt.Print("N")
//     } else {
//         for i, valor := range repetidos {
//             if i != 0 {
//                 fmt.Print(" ")
//             }
//             fmt.Printf("%v", valor)
//         }
//     }
//     fmt.Println("")
//     saida := ""
//     for i := 1; i < qtd_album; i++ {
//         if !unicos[i] {
//             saida += fmt.Sprintf("%v ", i)
//         }
//     }

//     if saida == "" {
//         fmt.Println("N")
//     } else {
//         fmt.Println(saida[:len(saida)-1])
//     }
// }

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
