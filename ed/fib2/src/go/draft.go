package main
import "fmt"

func Coelhos(n int) int {
    if n == 1 || n == 2 {
        return 1
    }

    if n == 3 {
        return 2
    }
    /*
    A - 1
    A - 1
    A, B - 2
    B, C - 2
    B, C, D - 3
    B, C, D, E - 4
    C, D, E, F - 5

    F(4) = F(2) + F(1) = 2
    F(5) = F(3) + F(2) = 3
    F(6) = F(4) + F(3) = 4
    F(7) = F(5) + F(4) = 5
    F(8) = F(6) + F(5) = 7
    */

    return Coelhos(n-2) + Coelhos(n-3)
}

func main() {
    var n int
    fmt.Scan(&n)

    fmt.Println(Coelhos(n))
}