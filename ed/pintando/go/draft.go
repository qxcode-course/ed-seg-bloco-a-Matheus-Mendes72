package main
import "fmt"
import "math"
func main() {
    var l1, l2, l3 float64
    fmt.Scan(&l1, &l2, &l3)

    p := (l1+l2+l3) / 2 
    area := math.Sqrt(p * (p-l1) * (p-l2) * (p-l3))

    fmt.Printf("%.2f\n", area)
}
