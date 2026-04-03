package main
import "fmt"

// func tostr(palavra string) string {
//     if len(palavra) == 0 {
//         return ""
//     }

//     if len(palavra) == 1 {
//         return string(palavra[0])
//     }

//     return string(palavra[0]) + "\n " + tostr(palavra[1:])
// }

func diagonal(palavra string, k int) {
    if len(palavra) == 0 {
        return
    }

    for i := 1; i <= k; i++ {
        fmt.Print(" ")
    }
    fmt.Println(string(palavra[0]))

    diagonal(palavra[1:], k+1)
}

func main() {
    var palavra string
    fmt.Scan(&palavra)

    diagonal(palavra, 0)
}
