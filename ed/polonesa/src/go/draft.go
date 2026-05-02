package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type Stack struct {
    data []string
}

func NewStack() *Stack {
    return &Stack{data: []string{}}
}

func (s *Stack) Push(v string) {
    s.data = append(s.data, v)
}

func (s *Stack) Pop() string {
    if len(s.data) == 0 {
        return ""
    }
    v := s.data[len(s.data)-1]
    s.data = s.data[:len(s.data)-1]
    return v
}

func (s *Stack) Top() string {
    if len(s.data) == 0 {
        return ""
    }
    return s.data[len(s.data)-1]
}

func (s *Stack) IsEmpty() bool {
    return len(s.data) == 0
}

func precedence(op string) int {
    switch op {
    case "+", "-":
        return 1
    case "*", "/":
        return 2
    case "^":
        return 3
    }
    return 0
}

func isOperator(token string) bool {
    return token == "+" || token == "-" || token == "*" || token == "/" || token == "^"
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    line := scanner.Text()
    
    tokens := strings.Fields(line)
    
    output := make([]string, 0)
    operators := NewStack()
    
    for _, token := range tokens {
        if isOperator(token) {
            for !operators.IsEmpty() && isOperator(operators.Top()) && 
                  precedence(operators.Top()) >= precedence(token) {
                output = append(output, operators.Pop())
            }
            operators.Push(token)
        } else {
            output = append(output, token)
        }
    }
    
    for !operators.IsEmpty() {
        output = append(output, operators.Pop())
    }
    
    fmt.Println(strings.Join(output, " "))
}