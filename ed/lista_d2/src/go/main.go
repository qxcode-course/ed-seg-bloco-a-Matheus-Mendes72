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
	root *Node
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
	// ll.root.root = ll.root
	ll.size = 0

	return ll
}

func (ll *LList) toString() {
	if ll.size == 0 {
		fmt.Println("[]")
		return
	}

	saida := "["
	current := ll.root.next
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

func (n *Node) Next() *Node {
	if n.next == n.root {
		return nil
	}

	return n.next
}

func (n *Node) Prev() *Node {
	if n.prev == n.root {
		return nil
	}

	return n.prev
}

func (ll *LList) Front() *Node {
	if ll.size == 0 {
		return nil
	}

	return ll.root.next
}

func (ll *LList) Back() *Node {
	if ll.size == 0 {
		return nil
	}

	return ll.root.prev
}

func (ll *LList) Clear() {
	ll.root.next = ll.root
    ll.root.prev = ll.root
	ll.size = 0
}

func (ll *LList) PushBack(value int) {
	newNode := &Node{
		value: value,
		root: ll.root,
	}

	newNode.prev = ll.root.prev
	newNode.next = ll.root

	ll.root.prev.next = newNode
	ll.root.prev = newNode

	ll.size++
}

func (ll *LList) PushFront(value int) {
	newNode := &Node{
		value: value,
		root: ll.root,
	}

	newNode.next = ll.root.next
	newNode.prev = ll.root

	ll.root.next.prev = newNode
	ll.root.next = newNode

	ll.size++
}

func (ll *LList) Search(value int) *Node {
	for node := ll.Front(); node != nil; node = node.Next() {
		if node.value == value {
			return node
		}
	}

	return nil
}

func (ll *LList) Insert(node *Node, value int) {
	newNode := &Node{
		value: value,
		root: ll.root,
	}

	newNode.prev = node.prev
	newNode.next = node

	node.prev.next = newNode
	node.prev = newNode

	ll.size++
}

func (ll *LList) Remove(node *Node) {
	if node == ll.root {
		return
	}

	node.prev.next = node.next
	node.next.prev = node.prev
	node.next = nil
	node.prev = nil
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
			ll.toString()
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
			// ll.PopBack()
		case "pop_front":
			// ll.PopFront()
		case "clear":
			ll.Clear()
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next() {
				fmt.Printf("%v ", node.value)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev() {
				fmt.Printf("%v ", node.value)
			}
			fmt.Println("]")
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				node.value = newvalue
			} else {
				fmt.Println("fail: not found")
			}
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Insert(node, newvalue)
			} else {
				fmt.Println("fail: not found")
			}
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Remove(node)
			} else {
				fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
