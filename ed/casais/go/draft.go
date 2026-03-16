package main
import "fmt"
func main() {
    var qtd, casais int
    //vetor := []int{}
    fmt.Scan(&qtd)

    slice := make(map[int]int)
    for i := 0; i < qtd; i++ {
        var animal int
        fmt.Scan(&animal)

        if slice[-animal] > 0 {
            slice[-animal]--
            casais++
        } else {
            slice[animal]++
        }

        // encontrou := false
        // for j := 0; j < len(vetor); j++ {
        //     if vetor[j] == -animal {
        //         vetor[j] = 0
        //         casais++
        //         encontrou = true
        //         continue
        //     }
        //     if !encontrou {}
        //         vetor = append(vetor, animal)
        //     }
    }
    fmt.Println(casais)
}
