package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestDoublyLinkedList_First(t *testing.T) {
	DLL := DoublyLinkedList{}
	firstElem, err := DLL.First()
	if firstElem != nil {
		t.Errorf("DoublyLinkedList.First() for empty list | Actual value = %v Expected value = %v", firstElem, nil)
	}
	if err == nil {
		t.Errorf("DoublyLinkedList.First() for empty list | Actual error = %v Expected error = %v", err,  errors.New("list is empty"))
	}

	DLL.PushBack(45)
	firstElem, err = DLL.First()
	if firstElem == nil {
		t.Errorf("DoublyLinkedList.First() for list with 1 element | Actual value = %v Expected value = %v", firstElem, Item{value:45})
	}
	if err != nil {
		t.Errorf("DoublyLinkedList.First() for list with 1 element | Actual error = %v Expected error = %v", err,  nil)
	}

	DLL.PushBack(54)
	firstElem, err = DLL.First()
	if firstElem == nil {
		t.Errorf("DoublyLinkedList.First() for list with 2 element | Actual value = %v Expected value = %v", firstElem, Item{value:45})
	}
	if err != nil {
		t.Errorf("DoublyLinkedList.First() for list with 2 element | Actual error = %v Expected error = %v", err,  nil)
	}
}
func TestDoublyLinkedList_Last(t *testing.T) {
	DLL := DoublyLinkedList{}
	lastElem, err := DLL.Last()
	if lastElem != nil {
		t.Errorf("DoublyLinkedList.Last() for empty list | Actual value = %v Expected value = %v", lastElem, nil)
	}
	if err == nil {
		t.Errorf("DoublyLinkedList.Last() for empty list | Actual error = %v Expected error = %v", err,  errors.New("list is empty"))
	}

	DLL.PushBack(45)
	lastElem, err = DLL.Last()
	if lastElem == nil {
		t.Errorf("DoublyLinkedList.Last() for list with 1 element | Actual value = %v Expected value = %v", lastElem, Item{value:45})
	}
	if err != nil {
		t.Errorf("DoublyLinkedList.Last() for list with 1 element | Actual error = %v Expected error = %v", err,  nil)
	}

	DLL.PushBack(54)
	lastElem, err = DLL.Last()
	if lastElem == nil {
		t.Errorf("DoublyLinkedList.Last() for list with 2 element | Actual value = %v Expected value = %v", lastElem, Item{value:54})
	}
	if err != nil {
		t.Errorf("DoublyLinkedList.Last() for list with 2 element | Actual error = %v Expected error = %v", err,  nil)
	}
}
func TestDoublyLinkedList_Len(t *testing.T) {
	DLL := DoublyLinkedList{}
	Len := DLL.Len()
	if DLL.Len() != 0{
		t.Errorf("DoublyLinkedList.Len() for empty list | Actual = %v Expected = %v", Len, 0)
	}
	DLL.PushBack(45)
	Len = DLL.Len()
	if DLL.Len() != 1{
		t.Errorf("DoublyLinkedList.Len() for list with 1 element | Actual = %v Expected = %v", Len, 1)
	}
	DLL.PushBack(5)
	DLL.PushBack(6)
	DLL.PushBack(7)
	DLL.PushBack(8)
	DLL.PushBack(9)
	Len = DLL.Len()
	if DLL.Len() != 6{
		t.Errorf("DoublyLinkedList.Len() for list with 6 element | Actual = %v Expected = %v", Len, 6)
	}
}
func TestDoublyLinkedList_PushBack(t *testing.T) {
	DLL := DoublyLinkedList{}
	elem1 := 12345
	elem2 := 54321

	DLL.PushBack(elem1)
	if DLL.tail.value != elem1 {
		t.Errorf("DoublyLinkedList.PushBack(%v) | Actual Tail = %v Expected Tail = %v", elem1, DLL.tail.value, elem1)
	}
	DLL.PushBack(elem2)

	if DLL.tail.value != elem2 {
		t.Errorf("DoublyLinkedList.PushBack(%v) | Actual Tail = %v Expected Tail = %v", elem2, DLL.tail.value, elem2)
	}
}
func TestDoublyLinkedList_PushFront(t *testing.T) {
	DLL := DoublyLinkedList{}
	elem1 := 56789
	elem2 := 98765

	DLL.PushFront(elem1)
	if DLL.head.value != elem1 {
		t.Errorf("DoublyLinkedList.PushFront(%v) | Actual Head = %v Expected Head = %v", elem1, DLL.head.value, elem1)
	}
	DLL.PushFront(elem2)
	if DLL.head.value != elem2 {
		t.Errorf("DoublyLinkedList.PushFront(%v) | Actual Head = %v Expected Head = %v", elem2, DLL.head.value, elem2)
	}
}
func TestDoublyLinkedList_GetByIndex(t *testing.T) {
	DLL := DoublyLinkedList{}
	index := 0
	elem, err := DLL.GetByIndex(index)
	if elem != nil {
		t.Errorf("DoublyLinkedList.GetByIndex(%v) for empty list | Actual value = %v Expected value = %v",index, elem, nil)
	}
	if err == nil{
		t.Errorf("DoublyLinkedList.GetByIndex(%v) for empty list | Actual error = %v Expected error = %v",index, err,  fmt.Errorf("no element with index %v", index).Error())
	}

	DLL.PushBack(25)
	DLL.PushBack(30)
	DLL.PushBack(35)
	DLL.PushBack(40)

	index = 2
	elem, err = DLL.GetByIndex(index)
	if elem == nil {
		t.Errorf("DoublyLinkedList.GetByIndex(%v) for list with 4 elements | Actual value = %v Expected value = %v", index, elem, 35)
	}
	if err != nil{
		t.Errorf("DoublyLinkedList.GetByIndex(%v) for list with 4 elements | Actual error = %v Expected error = %v", index, err.Error(), nil)
	}

	index = 999999
	elem, err = DLL.GetByIndex(index)
	if elem != nil {
		t.Errorf("DoublyLinkedList.GetByIndex(%v) for list with 4 elements | Actual value = %v Expected value = %v",index, elem, nil)
	}
	if err == nil{
		t.Errorf("DoublyLinkedList.GetByIndex(%v) for list with 4 elements | Actual error = %v Expected error = %v",index, err,  fmt.Errorf("no element with index %v", index).Error())
	}
}
func TestDoublyLinkedList_GetByValue(t *testing.T) {
	DLL := DoublyLinkedList{}
	value := 123
	elem, err := DLL.GetByValue(value)
	if elem != nil {
		t.Errorf("DoublyLinkedList.GetByValue(%v) for empty list | Actual value = %v Expected value = %v",value, elem, nil)
	}
	if err == nil{
		t.Errorf("DoublyLinkedList.GetByValue(%v) for empty list | Actual error = %v Expected error = %v",value, err,  fmt.Errorf("no element with index %v", value).Error())
	}

	DLL.PushBack(25)
	DLL.PushBack(30)
	DLL.PushBack(35)
	DLL.PushBack(40)

	value = 30
	elem, err = DLL.GetByValue(value)
	if elem == nil {
		t.Errorf("DoublyLinkedList.GetByValue(%v) for list with elements=[25, 30, 35, 40] | Actual value = %v Expected value = %v", value, elem, 30)
	}
	if err != nil{
		t.Errorf("DoublyLinkedList.GetByValue(%v) for list with elements=[25, 30, 35, 40] | Actual error = %v Expected error = %v", value, err.Error(), nil)
	}

	value = -42
	elem, err = DLL.GetByValue(value)
	if elem != nil {
		t.Errorf("DoublyLinkedList.GetByValue(%v) for list with elements=[25, 30, 35, 40] | Actual value = %v Expected value = %v",value, elem, nil)
	}
	if err == nil{
		t.Errorf("DoublyLinkedList.GetByValue(%v) for list with elements=[25, 30, 35, 40] | Actual error = %v Expected error = %v",value, err,  fmt.Errorf("no element with index %v", value).Error())
	}
}
func TestItem_Remove(t *testing.T) {
	DLL := DoublyLinkedList{}

	DLL.PushBack("first")
	DLL.PushBack("second")
	DLL.PushBack("third")
	DLL.PushBack("fourth")
	DLL.PushBack("fifth")

	//Removing head element
	value := "third"
	originalLen := DLL.Len()
	elem, err1 := DLL.GetByValue(value)
	if elem == nil {
		t.Errorf("DoublyLinkedList.GetByValue(%v) for list with elements=['first', 'second', 'third', 'fourth', 'fifth'] | Actual value = %v Expected value = %v",elem, elem, value)
	} else {
		elem.Remove()
		if DLL.Len() == originalLen {
			t.Errorf("Item.Remove() for item %v | Actual list len = %v Expected list len = %v", elem, DLL.Len(), originalLen - 1)
		}
		second, err2 := DLL.GetByValue("second")
		fourth, err3 := DLL.GetByValue("fourth")
		if second != nil && fourth != nil {
			if second.next != fourth {
				t.Errorf("Item.Remove() for item %v | Pointers of second and fourth elements | Actual second Next pointer/addr = %v/%p Expected second Next pointer/addr = %v/%p", elem, second.next, second.next, fourth, fourth)
			}
			if fourth.prev != second {
				t.Errorf("Item.Remove() for item %v | Pointers of second and fourth elements after deleting third element | Actual fourth Prev pointer/addr = %v/%p Expected fourth Prev pointer/addr = %v/%p", elem, fourth.prev, fourth.prev, second, second)
			}
		} else {
			t.Errorf("DoublyLinkedList.GetByValue(%v) and DoublyLinkedList.GetByValue(%v) | Actual errors: %v %v Expected errors: %v %v", "second", "fourth", err2, err3, nil, nil)
		}
	}
	if err1 != nil{
		t.Errorf("DoublyLinkedList.GetByValue(%v) for list with elements=['first', 'second', 'third', 'fourth', 'fifth'] | Actual error = %v Expected error = %v",value, err1,  nil)
	}

	//Removing first element
	value = "first"
	originalLen = DLL.Len()
	elem, err1 = DLL.GetByValue(value)
	if elem == nil {
		t.Errorf("DoublyLinkedList.GetByValue(%v) for list with elements=['first', 'second', 'fourth', 'fifth'] | Actual value = %v Expected value = %v",elem, elem, value)
	} else {
		elem.Remove()
		if DLL.Len() == originalLen {
			t.Errorf("Item.Remove() for item %v | Actual list len = %v Expected list len = %v", elem, DLL.Len(), originalLen - 1)
		}
		newHead := DLL.head
		newHeadPrev := newHead.prev
		second, err2 := DLL.GetByValue("second")
		if second != nil{
			if second != newHead {
				t.Errorf("Item.Remove() for first item %v | Expected Head == second | Actual second pointer/addr = %v/%p Actual Head pointer/addr = %v/%p", elem, second, second, newHead, newHead)
			} else {
				if newHeadPrev != nil {
					t.Errorf("Item.Remove() for item %v | Expected Head == second and second Prev == nil | Actual New Head Prev pointer/addr = %v/%p Expected New Head Prev pointer/addr = %v", elem, newHeadPrev, newHeadPrev, nil)
				}
			}
		} else {
			t.Errorf("DoublyLinkedList.GetByValue(%v) | Actual errors: %v Expected errors: %v", "second", err2, nil)
		}

	}
	if err1 != nil{
		t.Errorf("DoublyLinkedList.GetByValue(%v) for list with elements=['first', 'second', 'fourth', 'fifth'] | Actual error = %v Expected error = %v",value, err1,  nil)
	}

	//Removing tail element
	value = "fifth"
	originalLen = DLL.Len()
	elem, err1 = DLL.GetByValue(value)
	if elem == nil {
		t.Errorf("DoublyLinkedList.GetByValue(%v) for list with elements=['second', 'fourth', 'fifth'] | Actual value = %v Expected value = %v",elem, elem, value)
	} else {
		elem.Remove()
		if DLL.Len() == originalLen {
			t.Errorf("Item.Remove() for item %v | Actual list len = %v Expected list len = %v", elem, DLL.Len(), originalLen - 1)
		}
		newTail := DLL.tail
		newTailNext := newTail.next
		fourth, err3 := DLL.GetByValue("fourth")
		if fourth != nil{
			if fourth != newTail {
				t.Errorf("Item.Remove() for first item %v | Expected Tail == fourth | Actual fourth pointer/addr = %v/%p Actual Tail pointer/addr = %v/%p", elem, fourth, fourth, newTail, newTail)
			} else {
				if newTailNext != nil {
					t.Errorf("Item.Remove() for item %v | Expected Head == fourth and fourth Next == nil | Actual New Tail Next pointer/addr = %v/%p Expected New Tail Next pointer/addr = %v", elem, newTailNext, newTailNext, nil)
				}
			}
		} else {
			t.Errorf("DoublyLinkedList.GetByValue(%v) | Actual errors: %v Expected errors: %v", "foutth", err3, nil)
		}
	}
	if err1 != nil{
		t.Errorf("DoublyLinkedList.GetByValue(%v) for list with elements=['first', 'second', 'fourth', 'fifth'] | Actual error = %v Expected error = %v",value, err1,  nil)
	}
}
