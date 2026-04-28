package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type MultiSet struct {
	data []int
	size int
	capacity int
}

func NewMultiSet(capacity int) *MultiSet {
	return &MultiSet {
		data: make([]int, capacity),
		size: 0,
		capacity: capacity,
	}
}

func (ms *MultiSet) toString() {
	saida := ""
	if ms.size == 0 {
		fmt.Println("[]")
		return
	}

	saida += "[" 
	for i := 0; i < ms.size; i++ {
		if i > 0 {
			saida += ", "
		}

		saida += strconv.Itoa(ms.data[i])
	}
	saida += "]"

	fmt.Println(saida)
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func (ms *MultiSet) Expand() {
	newCap := ms.capacity * 2

	if newCap < ms.capacity {
		return 
	}

	if newCap == 0 {
		newCap = 1
	}

	newData := make([]int, newCap)
	for i := 0; i < ms.size; i++ {
		newData[i] = ms.data[i]
	}

	ms.data = newData
	ms.capacity = newCap
}

func (ms *MultiSet) Insert(value int) {
	if ms.size == ms.capacity {
		ms.Expand()
	}

	index := 0
	for i := 0; i < ms.size; i++ {
		if value < ms.data[i] {
			index = i
			ms.insert(value, index)
			return 
		}
	}

	ms.data[ms.size] = value
	ms.size++
}

func (ms *MultiSet) insert(value int, index int) error {
	if index < 0 || index > ms.size {
		return fmt.Errorf("index out of range")
	}

	for i := ms.size; i > index; i-- {
		ms.data[i] = ms.data[i-1]
	}

	ms.data[index] = value
	ms.size++

	return nil
}

func (ms *MultiSet) Contains(value int) bool {
	for i := 0; i < ms.size; i++ {
		if ms.data[i] == value {
			return true
		}
	}

	return false
}

func (ms *MultiSet) Erase(value int) bool {
	index := -1
	for i := 0; i <= ms.size; i++ {
		if ms.data[i] == value {
			index = i
			break
		}
	}

	if index == -1 {
		return false
	}

	for i := index; i < ms.size; i++ {
		if i == ms.size -1 {
			ms.data[i] = 0
			break
		}

		ms.data[i] = ms.data[i+1]
	}

	ms.size--

	return true
}

func (ms *MultiSet) Count(value int) int {
	//aux := ms.data[0]
	cont := 0
	for i := 0; i < ms.size; i++ {
		if ms.data[i] == value {
			cont++
		} 
	}

	return cont
}

func (ms *MultiSet) Unique() int {
	cont := 1
	aux := ms.data[0]
	for i := 0; i < ms.size; i++ {
		if aux != ms.data[i] {
			cont++
			aux = ms.data[i]
		}
	}

	if cont == 1 {
		return 0
	}

	return cont
}

func (ms *MultiSet) Clear() {
	ms.size = 0
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			ms.toString()
		case "erase":
			value, _ := strconv.Atoi(args[1])
			if !ms.Erase(value) {
				fmt.Println("value not found")
			}
		case "contains":
			value, _ := strconv.Atoi(args[1])
			if ms.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "count":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Count(value))
		case "unique":
			fmt.Println(ms.Unique())
		case "clear":
			ms.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
