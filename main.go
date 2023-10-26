package main

import (
	"fmt"
)

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	head *Node
	len  int
}

// CreateNode создание нового узла
func CreateNode(val int) *Node {
	node := Node{val, nil}
	return &node
}

// New создание нового связного списка длины q
func New(q int) *LinkedList {
	list := LinkedList{CreateNode(0), q}
	if q > 1 {
		temp := list.head
		for i := 1; i < q; i++ {
			temp.next = CreateNode(0)
			temp = temp.next
		}
	}
	list.len = q
	return &list

}

// Add добавление нового узла в конец списка
func (l *LinkedList) Add(val int) {
	if l.len == 0 { //если список еще пустой
		head := CreateNode(val)
		l.head = head
		l.len = 1
	} else {
		temp := l.head
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = CreateNode(val)
		l.len++
	}
}

// Pop удаление последнего элемента списка
func (l *LinkedList) Pop() {
	if l.len == 0 {
		fmt.Println("Empty list")
	} else {
		if l.len == 1 {
			l.head = nil
			l.len = 0
		} else {
			temp := l.head
			for temp.next.next != nil {
				temp = temp.next
			}
			temp.next = nil
			l.len--
		}
	}
}

// At вывод числа соответсвующего номеру pos в списке
func (l *LinkedList) At(val int) int {
	temp := l.head
	for i := 1; i < val; i++ {
		temp = temp.next
	}
	return temp.val
}

// Size размер списка
func (l *LinkedList) Size() int {
	return l.len
}

// DeleteFrom удаление элемента с места pos
func (l *LinkedList) DeleteFrom(pos int) {
	temp := l.head
	switch true {
	case pos > l.len || pos < 1:
		fmt.Println("wrong enter")

	case pos == 1: //если нужно удалить первый узел списка
		l.head = l.head.next

	case pos == l.len: //если удаляется последний узел списка
		for temp.next.next != nil {
			temp = temp.next
		}
		temp.next = nil

	default:
		for i := 1; i < pos-1; i++ {
			temp = temp.next
		}
		temp.next = temp.next.next
	}
	l.len--
}

// UpdateAt обновление значения узла по его номеру в списке
func (l *LinkedList) UpdateAt(pos, val int) {
	temp := l.head
	i := 1
	for i != pos {
		temp = temp.next
		i++
	}
	temp.val = val

}

// NewFromSlice создание нового списка из среза
func NewFromSlice(s []int) *LinkedList {
	list := New(len(s))
	for i := 0; i < len(s); i++ {
		list.UpdateAt(i+1, s[i])
	}
	return list
}

// PrintList вывод листа на экран
func (l *LinkedList) PrintList() {
	temp := l.head
	for temp != nil {
		fmt.Println(temp.val)
		temp = temp.next
	}
	fmt.Println("len = ", l.len)
}
