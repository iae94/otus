package main

import "testing"

func TestFindMax_Int(t *testing.T) {
	intSlice := SliceInt{17, 12, 36, 42, 41, 98, 64, 9, 0}
	intMax, intMaxErr := FindMax(intSlice)
	var max int = 98
	if intMaxErr != nil {
		t.Errorf("FindMax(intSlice) | Actual error: %v Expected error: %v", intMaxErr, nil)
	}
	if intMax != max {
		t.Errorf("FindMax(intSlice) | Actual max value: %v Expected max value: %v", intMax, max)
	}
}
func TestFindMax_Float(t *testing.T) {
	floatSlice := SliceFloat{0.99, 1.32, 2.03, 0.12, 0.99, 0.54, 2.04, 1, 0}
	floatMax, floatMaxErr := FindMax(floatSlice)
	var max float32 = 2.04
	if floatMaxErr != nil {
		t.Errorf("FindMax(floatSlice) | Actual error: %v Expected error: %v", floatMaxErr, nil)
	}
	if floatMax != max {
		t.Errorf("FindMax(floatSlice) | Actual max value: %v Expected max value: %v", floatMax, max)
	}
}
func TestFindMax_String(t *testing.T) {
	stringSlice := SliceString{"1234", "abcd", "ac", "acb", "bacd", "dacd", "a", "c", "db"}
	stringMax, stringMaxErr := FindMax(stringSlice)
	var max string = "dacd"
	if stringMaxErr != nil {
		t.Errorf("FindMax(stringSlice) | Actual error: %v Expected error: %v", stringMaxErr, nil)
	}
	if stringMax != max {
		t.Errorf("FindMax(stringSlice) | Actual max value: %v Expected max value: %v", stringMax, max)
	}
}
func TestFindMax_Struct(t *testing.T) {
	structSlice := SliceStruct{
		Person{"Dmitry", 30},
		Person{"Vasya", 17},
		Person{"Alexey", 21},
		Person{"Alex", 38},
		Person{"Denis", 29},

	}
	structMax, structMaxErr := FindMax(structSlice)
	var max *Person = &structSlice[3]
	if structMaxErr != nil {
		t.Errorf("FindMax(floatSlice) | Actual error: %v Expected error: %v", structMaxErr, nil)
	}
	if structMax != *max {
		t.Errorf("FindMax(floatSlice) | Actual max value: %v Expected max value: %v", structMax, *max)
	}
}
func TestFindMax_Empty(t *testing.T) {
	emptySlice := SliceString{}
	emptyMax, emptyMaxErr := FindMax(emptySlice)
	if emptyMaxErr == nil {
		t.Errorf("FindMax(emptySlice) | Actual error: %v Expected error: %v", emptyMaxErr, nil)
	}
	if emptyMax != nil {
		t.Errorf("FindMax(emptySlice) | Actual max value: %v Expected max value: %v", emptyMax, nil)
	}
}
