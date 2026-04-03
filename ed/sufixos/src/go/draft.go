package main
import "fmt"

func palavPedaço(palavra string) {
    if len(palavra) == 0 {
        return
    }

    palavPedaço(palavra[1:])
    fmt.Println(palavra)
}

func main() {
    var palavra string
    fmt.Scan(&palavra)

    palavPedaço(palavra)
}
