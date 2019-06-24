package main

import "fmt"
import "errors"

type Item struct {
	value interface{}
	next  *Item
	prev  *Item
	list  *DoublyLinkedList
}

func (item *Item) Remove() *Item {
	switch {
	case item.list.Len() == 1:
		item.list.head = nil
		item.list.tail = nil
	case item == item.list.head:
		item.list.head = item.next
		item.next.prev = nil
	case item == item.list.tail:
		item.list.tail = item.prev
		item.prev.next = nil
	default:
		item.prev.next = item.next
		item.next.prev = item.prev
	}
	item.list.length--
	return item
}

type DoublyLinkedList struct {
	head *Item
	tail *Item
	length int
}

func (d *DoublyLinkedList) Print() {
	elem := d.head
	if elem == nil {
		fmt.Printf("List is empty")
	} else {
		var index int
		for elem != nil {
			fmt.Printf("i = '%v' Value = '%v' This addr = '%p' Next addr = '%p' Prev addr = '%p'\n", index, elem.value, elem, elem.next, elem.prev)
			index++
			elem = elem.next
		}
	}
}
func (d *DoublyLinkedList) Len() int {	return d.length }

func (d *DoublyLinkedList) First() (*Item, error) {
	if d.Len() > 0 {
		return d.head, nil
	} else {
		return nil, errors.New("list is empty")
	}
}
func (d *DoublyLinkedList) Last() (*Item, error) {
	if d.Len() > 0 {
		return d.tail, nil
	} else {
		return nil, errors.New("list is empty")
	}
}

func (d *DoublyLinkedList) PushFront(elem interface{}) Item {
	head, err := d.First()
	var item Item
	if err != nil {
		item = Item{
			value: elem,
			next:  nil,
			prev:  nil,
			list:  d,
		}
		d.head = &item
		d.tail = &item
	} else {
		item = Item{
			value: elem,
			next:  head,
			prev:  nil,
			list:  d,
		}
		head.prev = &item
		d.head = &item
	}
	d.length++
	return item
}
func (d *DoublyLinkedList) PushBack(elem interface{}) Item {
	tail, err := d.Last()
	var item Item
	if err != nil {
		item = Item{
			value: elem,
			next:  nil,
			prev:  nil,
			list:  d,
		}
		d.head = &item
		d.tail = &item
	} else {
		item = Item{
			value: elem,
			next:  nil,
			prev:  tail,
			list:  d,
		}
		tail.next = &item
		d.tail = &item
	}
	d.length++
	return item
}
func (d *DoublyLinkedList) GetByIndex(i int) (*Item, error) {
	if i >= d.Len() {
		return nil, fmt.Errorf("no element with index %v", i)
	} else {
		elem := d.head
		var index int

		for index != i {
			index++
			elem = elem.next
		}
		return elem, nil
	}
}
func (d *DoublyLinkedList) GetByValue(value interface{}) (*Item, error) {
	if d.Len() != 0 {
		elem := d.head
		for elem != nil {
			if elem.value == value {
				return elem, nil
			}
			elem = elem.next
		}
	}
	return nil, fmt.Errorf("no element with value %v", value)
}

func main() {
	DLL := DoublyLinkedList{}
	slicel := []int{4, 6, 7, 3, 9, 1}
	slice2 := []string{"str1", "str2", "str3", "str4", "str5"}

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
