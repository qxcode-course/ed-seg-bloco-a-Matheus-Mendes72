package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MagicSearch(slice []int, value int) int {
	low := 0
	high := len(slice)
	pos := -1

	for low < high {
		mid := low + (high - low) / 2

		if value == slice[mid] {
			pos = mid
			break
		} else if value < slice[mid] {
			high = mid
		} else {
			low = mid + 1
		}
	}

	if pos != -1 {
		for i := 0; i < len(slice); i++ {
			if slice[i] == slice[pos] {
				pos = i
			}
		}
		
		return pos
	}

	return low
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	slice := make([]int, 0, 1)
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}

	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	result := MagicSearch(slice, value)
	fmt.Println(result)
}
