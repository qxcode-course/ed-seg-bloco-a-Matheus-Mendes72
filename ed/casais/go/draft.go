package main
import "fmt"
func casais(qtd int) int {
    casais := 0
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
    }
    return casais
}
func main() {
    var qtd int
    //vetor := []int{}
    fmt.Scan(&qtd)

    casais_formados := casais(qtd)
    fmt.Println(casais_formados)
}

for _, val := range slice {
        if slice[-i] > 0 {
            slice[-i]--
            casais++
        } else {
            slice[i]++
        }
    }