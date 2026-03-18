package main
import "fmt"

func main() {
    var qtd int
    fmt.Scan(&qtd)

    animal := make([]int, qtd)
    for i := 0; i < qtd; i++ {
        fmt.Scan(&animal[i])
    }

    solteiros := make(map[int]int) 
    casais := 0
    for _, animal := range animal {

        // Método para percorrer Map
        qtd, existe := solteiros[-animal]
    
        if existe && qtd > 0 {
            // Remove o valor do Map
            solteiros[-animal]--
            casais += 1
        } else {
            // Guarda o valor no Map
            solteiros[animal]++
        }
    }
    fmt.Println(casais)
}






// func main() {
//     solteiros := make(map[int]int)
//     qtd := 0
//     fmt.Scan(&qtd)
//     pares := 0
//     for range qtd {
//         animal := 0
//         fmt.Scan(&animal) 
//         qtd, existe := solteiros[-animal]
//         if existe && qtd > 0 {
//             solteiros[-animal] = qtd - 1
//             pares += 1
//         } else {
//             solteiros[animal] ++
//         }
//     }

//     fmt.Println(pares)
// }


// func casais(qtd int) int {
//     casais := 0
//     slice := make(map[int]int)
//     for i := 0; i < qtd; i++ {
//         var animal int
//         fmt.Scan(&animal)

//         if slice[-animal] > 0 {
//             slice[-animal]--
//             casais++
//         } else {
//             slice[animal]++
//         }
//     }
//     return casais
// }
// func main() {
//     var qtd int
//     //vetor := []int{}
//     fmt.Scan(&qtd)

//     casais_formados := casais(qtd)
//     fmt.Println(casais_formados)
// }

