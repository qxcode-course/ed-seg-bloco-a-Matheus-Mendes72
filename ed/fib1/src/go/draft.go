package main
import "fmt"

func Fibonacci(n, k int) int {
    if n == 1 || n == 2 {
        return 1
    }

    return Fibonacci(n-1, k) + k*Fibonacci(n-2, k)
}

func main() {
    var n, k int
    fmt.Scan(&n, &k)

    fmt.Println(Fibonacci(n, k))
}