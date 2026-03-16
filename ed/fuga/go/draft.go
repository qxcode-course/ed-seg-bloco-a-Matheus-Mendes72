package main
import "fmt"
func main() {
    var h, p, f, d int
    fmt.Scan(&h, &p, &f, &d)

    for {
        f = (f + d + 16) % 16

        if (f == h) {
            fmt.Println("S")
            break
        }
        if (f == p) {
            fmt.Println("N")
            break
        }
    }
}
