package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getMen(vet []int) []int {
	vetor := make([]int, 0)
	for i := 0; i < len(vet); i++ {
		if vet[i] < 0 {
			continue
		}

		vetor = append(vetor, vet[i])
	}
	
	return vetor
}

func getCalmWomen(vet []int) []int {
	vetor := make([]int, 0)

	for i := 0; i < len(vet); i++ {
		
		if vet[i] < 0 {
			aux := -vet[i]
		
			if aux < 10 {
				vetor = append(vetor, vet[i])
			}
		}
	} 

	return vetor
}

func sortVet(vet []int) []int {
	for i := 0; i < len(vet); i++ {
		min := i

		for j := i +1; j < len(vet); j++ {
			if vet[j] < vet[min] {
				min = j
			}

			vet[i], vet[min] = vet[min], vet[i]
		}
	}
	return vet
}

func sortStress(vet []int) []int {
	for i := 0; i < len(vet); i++ {
		min := i

		for j := i+1; j < len(vet); j++ {
			aux := vet[j]
			aux2 := vet[min]

			if vet[j] < 0 {
				aux = -vet[j]
			}
			if vet[min] < 0 {
				aux2 = -vet[min]
			}

			if aux < aux2 {
				min = j
			}

		}
		vet[i], vet[min] = vet[min], vet[i]
	}

	return vet
}

func reverse(vet []int) []int {
	vetor := make([]int, 0)

	for i := 0; i < len(vet); i++ {
		vetor = append(vetor, vet[i])
	}

	for i := 0; i < len(vetor)/2; i++ {
		j := len(vetor) -1 -i
		vetor[i], vetor[j] = vetor[j], vetor[i]
	}

	return vetor
}

func unique(vet []int) []int {
	mapa := make(map[int]bool)
	sozinhos := make([]int, 0)

	for _, valor := range vet {
		if !mapa[valor] {
			mapa[valor] = true
			sozinhos = append(sozinhos, valor)
		}
	} 

	return sozinhos
}

func repeated(vet []int) []int {
	mapa := make(map[int]int)
	repetidos := make([]int, 0)

	for _, valor := range vet {
		if mapa[valor] >= 0 {
			mapa[valor] += 1

			if mapa[valor] == 2 {
				repetidos = append(repetidos, valor)
				mapa[valor] = 1
			}
		}
	} 

	return repetidos
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}

