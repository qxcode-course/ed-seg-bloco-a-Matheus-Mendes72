package main
import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    input := scanner.Text()

    stack := make([]rune, 0)
    for _, ch := range input {
        if ch == '(' || ch == '[' {
            stack = append(stack, ch)
        } else if ch == ')' || ch == ']' {
            if len(stack) == 0 {
                fmt.Println("nao balanceado")
                return
            }
        }

        topo := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if (topo == '(' && ch != ')') || (topo == '[' && ch != ']') {
            fmt.Println("nao balanceado")
            return
        }
    }

    if len(stack) == 0 {
        fmt.Println("balanceado")
    } else {
        fmt.Println("nao balanceado")
    }
}
