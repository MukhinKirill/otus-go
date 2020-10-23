package main

import "fmt"

type MyList struct {
	First  *Item
	Last   *Item
	Length int
}

type Item struct {
	Value int
	Next  *Item
	Prev  *Item
}

func NewMyList() *MyList {
	p := new(MyList)
	p.Length = 0
	return p
}

func (list *MyList) PushFront(value int) {
	item := new(Item)
	item.Value = value
	if list.First == nil {
		list.First = item
		list.Last = item
	} else {
		item.Next = list.First
		list.First.Prev = item
		list.First = item
	}
	list.Length++
}

func (list *MyList) PushBack(value int) {
	item := new(Item)
	item.Value = value
	if list.First == nil {
		list.First = item
		list.Last = item
	} else {
		item.Prev = list.Last
		list.Last.Next = item
		list.Last = item
	}
	list.Length++
}

func (list *MyList) Remove(item *Item) {
	if item == list.First {
		list.First = list.First.Next
		list.First.Prev = nil
		list.First.Next.Prev = list.First
		list.Length--
		return
	}
	if item == list.Last {
		list.Last = list.Last.Prev
		list.Last.Next = nil
		list.Last.Prev.Next = list.Last
		list.Length--
		return
	}
	item.Next.Prev = item.Prev
	item.Prev.Next = item.Next
	list.Length--
}
func (list *MyList) Print() {
	if list.First != nil {
		current := list.First
		for current != nil {
			fmt.Printf("%d\n", current.Value)
			current = current.Next
		}
	}
}
