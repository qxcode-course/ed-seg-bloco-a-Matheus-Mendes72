package main
import "fmt"

func escada(numDegraus int) int {
    if numDegraus <= 2 {
        return 1
    }

    if numDegraus == 3 {
        return 2
    }

    return escada(numDegraus-1) + escada(numDegraus-3)
}

func main() {
    var numDegraus int 
    fmt.Scan(&numDegraus)

    fmt.Println(escada(numDegraus))
}
