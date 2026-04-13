package main
import "fmt"
func main() {
    var qtd int
    fmt.Scan(&qtd)

    animais := make([]int, qtd)
    for i := 0; i < qtd; i++ {
        fmt.Scan(&animais[i])
    }

    varAuxiliar := make([]int, len(animais)-1)
    contador := 0

    for i := 0; i < len(animais) -1; i++ {
        varAuxiliar[i] = animais[i+1]

        if (varAuxiliar[i] == -animais[i]) {
            contador++
        }
    }

    fmt.Println(contador)
}
