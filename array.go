package main

import "fmt"

type List struct {
	data string
	size uint16
}

func (l *List) Add(datum string) string {
	l.size += 1
	if l.data == "" {
		l.data = "[" + datum + "]"
		return l.data
	}
	l.data = l.data[:len(l.data) - 1] + ", " + datum + "]"
	return l.data
}

func (l *List) Display() {
	fmt.Println(l.data)
}

func (l *List) Get(index uint16) {
	var count uint16 = 0
}

func (l *List) Size() uint16 {
	return l.size
}

func (l *List) IsEmpty() bool {
	if l.data == "" {
		return true
	} else {
		return false
	}
}

func main() {
	var list *List = &List{}
	fmt.Println(list.IsEmpty())
}