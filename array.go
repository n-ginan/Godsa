package main

import "fmt"

type List struct {
	data string
	size uint16
}

func (l *List) Add(datum string) string {
	l.size++
	if l.data == "" || l.data == "[]" {
		l.data = "[" + datum + "]"
		return l.data
	}
	l.data = l.data[:len(l.data) - 1] + ", " + datum + "]"
	return l.data
}

func (l *List) Display() {
	fmt.Println(l.data)
}

func (l *List) Get(index int16) string {
	if l.size == 0 {
		return "There are no existing elements inside the list"
	}
	previous, current := l.position(index)
	if index == 0 {
		return l.data[previous + 1 : current]
	} else {
		return l.data[previous + 2 : current]
	}
}

func (l *List) Replace(new_datum string, index int16) string {
	if l.size == 0 {
		return "There are no existing elements inside the list"
	}
	if int16(l.size) <= index {
		return "You have exceeded the size limit of the list"
	}
	previous, current := l.position(index)
	if index == 0 {
		l.data = "[" + new_datum + ", " + l.data[current + 2 : len(l.data)]
	} else {
		l.data = l.data[0 : previous] + ", " + new_datum + l.data[current : len(l.data)]
	}
	return l.data
}

func (l *List) Remove(index int16) string {
	if l.size == 0 {
		return "There are no existing elements inside the list"
	}
	previous, current := l.position(index)
	if index == 0 {
		l.data = "[" + l.data[current + 2 : len(l.data)]
	} else {
		l.data = l.data[0 : previous] + l.data[current : len(l.data)]
	}
	return l.data
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

func (l *List) Contains(datum string) bool {
	var commaPosition int = 0
	var currentPosition int = 0
	for position, letter := range l.data {
		if string(letter) == "," || string(letter) == "]" {
			currentPosition = position
			if l.data[commaPosition + 1 : currentPosition] == datum ||
			   l.data[commaPosition + 2 : currentPosition] == datum {
				   return true
			   }
		}
		commaPosition = currentPosition
	}
	return false
}

func (l *List) IndexOf(datum string) int {
	var commaCount int = -1
	var commaPosition int = 0
	var currentPosition int = 0
	for position, letter := range l.data {
		if string(letter) == "," || string(letter) == "]" {
			currentPosition = position
			commaCount++
			if l.data[commaPosition + 1 : currentPosition] == datum ||
			   l.data[commaPosition + 2 : currentPosition] == datum {
				   return commaCount
			   }
		}
		commaPosition = currentPosition
	}
	return 0
}

func (l * List) Clear() string {
	l.data = "[]"
	return l.data
}

func (l *List) position(index int16) (int, int) {
	var commaCount int16 = -1
	var commaPosition int = 0
	var currentPosition int = 0
	for position, letter := range l.data {
		if string(letter) == "," || string(letter) == "]" {
			currentPosition = position
			commaCount++
		}
		if commaCount == index {
			break
		}
		commaPosition = currentPosition
	}
	return commaPosition, currentPosition
}

func main() {
	var list *List = &List{}
	list.Add("This")
	list.Add("Is")
	list.Add("A")
	list.Add("Data")
	list.Add("Adatu")
	fmt.Println(list.Replace("Updated Data", 1))
	fmt.Println(list.Replace("Hotdog", 0))
	fmt.Println(list.Replace("Sinugba", 3))
	fmt.Println(list.Size())
}