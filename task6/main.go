package main

import "fmt"

func main() {

	list := NewMyList()
	list.PushFront(5)
	list.PushFront(4)
	list.PushFront(3)
	list.PushFront(2)
	list.PushFront(1)
	list.PushFront(0)
	list.PushBack(6)
	list.PushBack(7)
	list.Print()
	fmt.Println("-----------------------------")
	list.Remove(list.First)
	list.Remove(list.Last)
	list.Remove(list.Last.Prev.Prev)
	list.Print()

}

//Двусвязный список
//Цель: https://en.wikipedia.org/wiki/Doubly_linked_list
//List // тип контейнер
//Len() // длинна списка
//First() // первый Item
//Last() // последний Item
//PushFront(v interface{}) // добавить значение в начало
//PushBack(v interface{}) // добавить значение в конец
//Remove(i Item) // удалить элемент

//Item // элемент списка
//Value() interface{} // возвращает значение
//Next() *Item // следующий Item
//Prev() *Item // предыдущий ```
//Реализовать двусвязанный список на языке Go
