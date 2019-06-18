package main

import "fmt"
import "errors"

type Item struct {
	Value interface{}
	Next  *Item
	Prev  *Item
	List  *DoublyLinkedList
}

func (item *Item) Remove() *Item {
	switch {
	case item.List.Len() == 1:
		item.List.Head = nil
		item.List.Tail = nil
	case item == item.List.Head:
		item.List.Head = item.Next
		item.Next.Prev = nil
	case item == item.List.Tail:
		item.List.Tail = item.Prev
		item.Prev.Next = nil
	default:
		item.Prev.Next = item.Next
		item.Next.Prev = item.Prev
	}

	return item
}

type DoublyLinkedList struct {
	Head *Item
	Tail *Item
}

func (d *DoublyLinkedList) Print() {

	elem := d.Head
	if elem == nil {
		fmt.Printf("List is empty")
	} else {
		var index int
		for elem != nil {

			fmt.Printf("i = '%v' Value = '%v' This addr = '%p' Next addr = '%p' Prev addr = '%p'\n", index, elem.Value, elem, elem.Next, elem.Prev)
			index++
			elem = elem.Next
		}
	}
}
func (d *DoublyLinkedList) Len() int {
	elem := d.Head
	if elem == nil {
		return 0
	} else {
		var size int
		for elem != nil {
			size++
			elem = elem.Next
		}
		return size
	}

}
func (d *DoublyLinkedList) First() (*Item, error) {
	if d.Len() > 0 {
		return d.Head, nil
	} else {
		return nil, errors.New("list is empty")
	}
}
func (d *DoublyLinkedList) Last() (*Item, error) {
	if d.Len() > 0 {
		return d.Tail, nil
	} else {
		return nil, errors.New("list is empty")
	}
}

func (d *DoublyLinkedList) PushFront(elem interface{}) Item {
	head, err := d.First()
	var item Item
	if err != nil {
		item = Item{
			Value: elem,
			Next:  nil,
			Prev:  nil,
			List:  d,
		}
		d.Head = &item
		d.Tail = &item

	} else {
		item = Item{
			Value: elem,
			Next:  head,
			Prev:  nil,
			List:  d,
		}
		head.Prev = &item
		d.Head = &item
	}
	return item
}
func (d *DoublyLinkedList) PushBack(elem interface{}) Item {
	tail, err := d.Last()
	var item Item
	if err != nil {
		item = Item{
			Value: elem,
			Next:  nil,
			Prev:  nil,
			List:  d,
		}
		d.Head = &item
		d.Tail = &item

	} else {
		item = Item{
			Value: elem,
			Next:  nil,
			Prev:  tail,
			List:  d,
		}
		tail.Next = &item
		d.Tail = &item
	}
	return item
}
func (d *DoublyLinkedList) GetByIndex(i int) (*Item, error) {
	if i >= d.Len() {
		return nil, fmt.Errorf("no element with index %v", i)
	} else {
		elem := d.Head
		var index int

		for index != i {
			index++
			elem = elem.Next
		}
		return elem, nil
	}

}
func (d *DoublyLinkedList) GetByValue(value interface{}) (*Item, error) {

	if d.Len() != 0 {
		elem := d.Head
		for elem != nil {
			if elem.Value == value {
				return elem, nil
			}
			elem = elem.Next
		}
	}
	return nil, fmt.Errorf("no element with value %v", value)

}

func main() {

	slicel := []int{4, 6, 7, 3, 9, 1}
	slice2 := []string{"str1", "str2", "str3", "str4", "str5"}

	DLL := DoublyLinkedList{}

	//Push back slice1
	for _, s := range slicel {
		DLL.PushBack(s)
	}
	//Push front slice2
	for _, s := range slice2 {
		DLL.PushFront(s)
	}
	//Print list
	DLL.Print()

	//Removing element and print list
	delElem, err := DLL.GetByValue("str2")
	if err == nil {
		delElem.Remove()
		fmt.Printf("Deleting element %v\n", delElem)
		DLL.Print()
	} else {
		fmt.Printf(err.Error())
	}
}
