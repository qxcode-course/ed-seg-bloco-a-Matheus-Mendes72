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
	head *Node
	size int
}

func NewLList() *LList {
	ll := &LList{}
	ll.head = &Node{}
	ll.head.next = ll.head
	ll.head.prev = ll.head
	ll.size = 0

	return ll
}

func (ll *LList) toString() {
	if ll.size == 0 {
		fmt.Println("[]")
		return
	}

	saida := ""
	current := ll.head.next
	for i := 0; i < ll.size; i++ {
		if i > 0 {
			saida += ", "
		}

		saida += strconv.Itoa(current.value)
		current = current.next
	}
	saida += "]"

	fmt.Println(saida)
}

func (ll *LList) Size() int {
	return ll.size
}

func (ll *LList) PushBack(value int) {
	newNode := &Node{value: value}
	
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
			ll.toString()
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			// for _, v := range args[1:] {
			// 	num, _ := strconv.Atoi(v)
			// 	ll.PushFront(num)
			// }
		case "pop_back":
			// ll.PopBack()
		case "pop_front":
			// ll.PopFront()
		case "clear":
			// ll.Clear()
		case "walk":
			// fmt.Print("[ ")
			// for node := ll.Front(); node != nil; node = node.Next() {
			// 	fmt.Printf("%v ", node.Value)
			// }
			// fmt.Print("]\n[ ")
			// for node := ll.Back(); node != nil; node = node.Prev() {
			// 	fmt.Printf("%v ", node.Value)
			// }
			// fmt.Println("]")
		case "replace":
			// oldvalue, _ := strconv.Atoi(args[1])
			// newvalue, _ := strconv.Atoi(args[2])
			// node := ll.Search(oldvalue)
			// if node != nil {
			// 	node.Value = newvalue
			// } else {
			// 	fmt.Println("fail: not found")
			// }
		case "insert":
			// oldvalue, _ := strconv.Atoi(args[1])
			// newvalue, _ := strconv.Atoi(args[2])
			// node := ll.Search(oldvalue)
			// if node != nil {
			// 	ll.Insert(node, newvalue)
			// } else {
			// 	fmt.Println("fail: not found")
			// }
		case "remove":
			// oldvalue, _ := strconv.Atoi(args[1])
			// node := ll.Search(oldvalue)
			// if node != nil {
			// 	ll.Remove(node)
			// } else {
			// 	fmt.Println("fail: not found")
			// }
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
