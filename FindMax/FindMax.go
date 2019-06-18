package main

import (
	"errors"
	"fmt"
)

type Comparable interface {
	Compare(i, j int) bool
	Len() int
	Get(i int) interface{}
}

type Person struct {
	name string
	age  int
}

type SliceInt []int
type SliceFloat []float32
type SliceString []string
type SliceStruct []Person

func (s SliceInt) Compare(i, j int) bool   { return s[i] > s[j] }
func (s SliceFloat) Compare(i, j int) bool { return s[i] > s[j] }
func (s SliceString) Compare(i, j int) bool {
	if len(s[i]) == len(s[j]) {
		return s[i] > s[j]
	} else {
		return len(s[i]) > len(s[j])
	}
}
func (s SliceStruct) Compare(i, j int) bool { return s[i].age > s[j].age }

func (s SliceInt) Len() int    { return len(s) }
func (s SliceFloat) Len() int  { return len(s) }
func (s SliceString) Len() int { return len(s) }
func (s SliceStruct) Len() int { return len(s) }

func (s SliceInt) Get(i int) interface{}    { return s[i] }
func (s SliceFloat) Get(i int) interface{}  { return s[i] }
func (s SliceString) Get(i int) interface{} { return s[i] }
func (s SliceStruct) Get(i int) interface{} { return s[i] }

func FindMax(slice Comparable) (interface{}, error) {
	if slice.Len() == 0 {
		return nil, errors.New("slice is empty")
	}
	maxIndex := 0
	for i := 0; i < slice.Len(); i++ {
		if slice.Compare(i, maxIndex) {
			maxIndex = i
		}
	}
	return slice.Get(maxIndex), nil
}

func main() {

	intSlice := SliceInt{1, 2, 3, 4, 5, 80, 3, 56, 29}
	floatSlice := SliceFloat{0.4, 0.89, 2.1, 2.3, 0.87, 0.99}
	stringSlice := SliceString{"abcd", "bacd", "bcd", "cbd", "bd"}
	structSlice := SliceStruct{Person{"Alex", 15}, Person{"Dmitry", 27}, Person{"Roman", 22}, Person{"David", 28}}
	emptySlice := make(SliceString, 0)

	intMax, intMaxErr := FindMax(intSlice)
	floatMax, floatMaxErr := FindMax(floatSlice)
	stringMax, stringMaxErr := FindMax(stringSlice)
	structMax, structMaxErr := FindMax(structSlice)
	emptyMax, emptyMaxErr := FindMax(emptySlice)

	fmt.Printf("Int slice max: %v Error: %v\n", intMax, intMaxErr)
	fmt.Printf("Float slice max: %v Error: %v\n", floatMax, floatMaxErr)
	fmt.Printf("String slice max: %v Error: %v\n", stringMax, stringMaxErr)
	fmt.Printf("Struct slice max: %v Error: %v\n", structMax, structMaxErr)
	fmt.Printf("Empty slice max: %v Error: %v\n", emptyMax, emptyMaxErr)
}
