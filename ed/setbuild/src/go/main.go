package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	//"math"
)

type Vector struct {
	data []int
	size int
	capacity int
}

func (v *Vector) toString() {
	saida := ""
	if v.size == 0 {
		fmt.Println("[]")
		return
	}

	saida += "["
	for i := 0; i < v.size; i++ {
		if i > 0 {
			saida += ", "
		}

		saida += strconv.Itoa(v.data[i])
	}
	saida += "]"
	
	fmt.Println(saida)
}

func (v *Vector) Reserve() {
	newCap := 2 * v.capacity

	if newCap <= v.capacity {
		return
	}

	if newCap == 0 {
		newCap = 1
	}

	newData := make([]int, newCap)
	for i := 0; i < v.size; i++ {
		newData[i] = v.data[i]
	}

	v.data = newData
	v.capacity = newCap
}

func NewSet(capacity int) *Vector {
	return &Vector{
		data: make([]int, capacity),
		size: 0,
		capacity: capacity,
	}
}

func (v *Vector) insert(value int, index int) error {
	if index < 0 || index > v.size {
		return fmt.Errorf("index out of range")
	}

	for i := v.size; i > index; i-- {
		v.data[i] = v.data[i-1]
	}

	v.data[index] = value
	v.size++
	return nil
}

func (v *Vector) Insert(value int) {
	if v.size == v.capacity {
		v.Reserve()
	}

	//menor := math.MaxInt
	index := 0
	for i := 0; i < v.size; i++ {
		if value == v.data[i] {
			return
		}

		if value < v.data[i] {
			index = i
			v.insert(value, index)
			return
		}
	}

	v.data[v.size] = value
	v.size++
}

func (v *Vector) Contains(value int) bool {
	for i := 0; i < v.size; i++ {
		if v.data[i] == value {
			return true
		}
	}

	return false
}

func (v *Vector) Erase(value int) bool {
	index := -1
	for i := 0; i < v.size; i++ {
		if value == v.data[i] {
			index = i
			break
		}
	}
	
	if index == -1 {
		return false
	}

	for i := index; i < v.size; i++ {
		if i == v.size -1 {
			v.data[i] = 0
			break
		}
		v.data[i] = v.data[i+1]
	}

	v.size--
	return true
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.Insert(value)
			}
		case "show":
			v.toString()
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			if !v.Erase(value) {
				fmt.Println("value not found")
			}
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
