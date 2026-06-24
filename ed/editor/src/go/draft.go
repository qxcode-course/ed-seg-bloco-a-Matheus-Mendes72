package main
import (
    "fmt"
    "bufio"
    "os"
)
func main() {
    in := bufio.NewReader(os.Stdin)

	var entrada string
	fmt.Fscan(in, &entrada)

    texto := []rune{}
    cursor := 0

    for _, c := range entrada {
        switch c {
        case 'R':
            texto = append(texto[:cursor], append([]rune{'\n'}, texto[cursor:]...)...)
            cursor++

        case '<':
            if cursor > 0 {
                cursor--
            }

        case '>':
            if cursor < len(texto) {
                cursor++
            }

        case 'B':
            if cursor > 0 {
                texto = append(texto[:cursor-1], texto[cursor:]...)
                cursor--
            }

        case 'D':
            if cursor < len(texto) {
                texto = append(texto[:cursor], texto[cursor+1:]...)
            }

        default:
            texto = append(texto[:cursor], append([]rune{c}, texto[cursor:]...)...)
            cursor++
        }
    }

    for i := 0; i <= len(texto); i++ {
        if i == cursor {
            fmt.Print("|")
        }

        if i < len(texto) {
            fmt.Printf("%c", texto[i])
        }
    }

    fmt.Println()
}
