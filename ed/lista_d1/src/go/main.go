package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Node struct {
	value int
	next *Node
	prev *Node
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	ll := &LList{}
	ll.root = &Node{}
	ll.root.next = ll.root
	ll.root.prev = ll.root
	ll.size = 0

	return ll
}

func (ll *LList) String() {
	saida := ""
	if ll.size == 0 {
		fmt.Println("[]")
		return
	}

	saida += "["
	aux := ll.root.next
	for i := 0; i < ll.size; i++ {
		if i > 0 {
			saida += ", "
		}

		saida += strconv.Itoa(aux.value)
		aux = aux.next
	}
	saida += "]"

	fmt.Println(saida)

	return
}

func (ll *LList) Size() int {
	return ll.size
}

func (ll *LList) PushFront(value int) {
	newNode := &Node{value: value}

	first := ll.root.next

	newNode.next = first
	newNode.prev = ll.root

	first.prev = newNode

	ll.root.next = newNode

	ll.size++
}

func (ll *LList) PushBack(value int) {
	newNode := &Node{value: value}

	last := ll.root.prev

	newNode.prev = last
	newNode.next = ll.root

	last.next = newNode

	ll.root.prev = newNode

	ll.size++
}

func (ll *LList) Clear() {
	ll.root.next = ll.root
	ll.root.prev = ll.root
	ll.size = 0
}

func (ll *LList) PopFront() {
	if ll.size == 0 {
		return
	}

	first := ll.root.next
	second := first.next

	ll.root.next = second
	second.prev = ll.root

	ll.size--
}

func (ll *LList) PopBack() {
	if ll.size == 0 {
		return
	}

	last := ll.root.prev
	antLast := last.prev

	ll.root.prev = antLast
	antLast.next = ll.root

	ll.size--
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
 	ll := NewLList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			ll.String()
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
