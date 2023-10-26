package linkedList

import (
	"fmt"
)

type node struct {
	val  int
	next *node
}

type LinkedList struct {
	head *node
	len  int
}

// CreateNode создание нового узла
func CreateNode(val int) *node {
	node := node{val, nil}
	return &node
}

// New создание нового связного списка длины q
func New(q int) *LinkedList {
	list := LinkedList{CreateNode(0), q}
	temp := list.head
	for i := 1; i < q; i++ {
		temp.next = CreateNode(0)
		temp = temp.next
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
		return
	}
	temp := l.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = CreateNode(val)
	l.len++
}

// Pop удаление последнего элемента списка
func (l *LinkedList) Pop() {
	switch {
	case l.len == 0:
		fmt.Println("Empty list")
		return
	case l.len == 1:
		l.head = nil
		l.len = 0
		return
	}
	temp := l.head
	for temp.next.next != nil {
		temp = temp.next
	}
	temp.next = nil
	l.len--
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
	if pos > l.len || pos < 1 {
		fmt.Println("wrong enter")
		return
	}
	l.len--
	if pos == 1 {
		l.head = l.head.next
		return
	}
	for i := 1; i < pos-1; i++ {
		temp = temp.next
	}
	temp.next = temp.next.next

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
	list := LinkedList{CreateNode(s[0]), len(s)}
	temp := list.head
	for i := 1; i < len(s); i++ {
		temp.next = CreateNode(s[i])
		temp = temp.next
	}
	return &list
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
