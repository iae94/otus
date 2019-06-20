package main

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestLogOtusEvent_Stdout(t *testing.T) {
	homeworkSubmit := HwSubmitted{Date: time.Now(), Id: 47, Code: `func main(){ fmt.Println("Hello world") }`, Comment: "Please take a look at my homework"}
	homeworkAccept := HwAccepted{Date: time.Now(), Id: 47, Grade: 8}

	//Writer Stdout test
	writer := os.Stdout

	num_bytes, err := LogOtusEvent(homeworkSubmit, writer)
	expected_bytes := len(homeworkSubmit.PrintEvent())

	if err != nil {
		t.Errorf("Write HwSubmitted in Stdout | Actual error: %v Expected error: %v\n", err, nil)
	}
	if num_bytes != expected_bytes {
		t.Errorf("Write HwSubmitted in Stdout | Actual written bytes: %v Expected written bytes: %v\n", num_bytes, expected_bytes)
	}

	num_bytes, err = LogOtusEvent(homeworkAccept, writer)
	expected_bytes = len(homeworkAccept.PrintEvent())
	if err != nil {
		t.Errorf("Write HwAccepted in Stdout | Actual error: %v Expected error: %v\n", err, nil)
	}
	if num_bytes != expected_bytes {
		t.Errorf("Write HwAccepted in Stdout | Actual written bytes: %v Expected written bytes: %v\n", num_bytes, expected_bytes)
	}
}

func TestLogOtusEvent_Stderr(t *testing.T) {
	homeworkSubmit := HwSubmitted{Date: time.Now(), Id: 47, Code: `func main(){ fmt.Println("Hello world") }`, Comment: "Please take a look at my homework"}
	homeworkAccept := HwAccepted{Date: time.Now(), Id: 47, Grade: 8}

	//Writer Stderr test
	writer := os.Stderr

	num_bytes, err := LogOtusEvent(homeworkSubmit, writer)
	expected_bytes := len(homeworkSubmit.PrintEvent())
	if err != nil {
		t.Errorf("Write HwSubmitted in Stderr | Actual error: %v Expected error: %v\n", err, nil)
	}
	if num_bytes != expected_bytes {
		t.Errorf("Write HwSubmitted in Stderr | Actual written bytes: %v Expected written bytes: %v\n", num_bytes, expected_bytes)
	}

	num_bytes, err = LogOtusEvent(homeworkAccept, writer)
	expected_bytes = len(homeworkAccept.PrintEvent())
	if err != nil {
		t.Errorf("Write HwAccepted in Stderr | Actual error: %v Expected error: %v\n", err, nil)
	}
	if num_bytes != expected_bytes {
		t.Errorf("Write HwAccepted in Stderr | Actual written bytes: %v Expected written bytes: %v\n", num_bytes, expected_bytes)
	}
}

func TestLogOtusEvent_FileWriter(t *testing.T) {
	homeworkSubmit := HwSubmitted{Date: time.Now(), Id: 47, Code: `func main(){ fmt.Println("Hello world") }`, Comment: "Please take a look at my homework"}
	homeworkAccept := HwAccepted{Date: time.Now(), Id: 47, Grade: 8}

	//Writer file test
	fileName := "FileWriter.txt"
	writer, err := os.Create(fileName)
	if err != nil {
		t.Errorf("Create file with name: %v | Actual error: %v Expected error: %v", fileName, err, nil)
	}
	defer writer.Close()

	num_bytes, err := LogOtusEvent(homeworkSubmit, writer)
	expected_bytes := len(homeworkSubmit.PrintEvent())
	if err != nil {
		t.Errorf("Write HwSubmitted in SliceWriter | Actual error: %v Expected error: %v\n", err, nil)
	}
	if num_bytes != expected_bytes {
		t.Errorf("Write HwSubmitted in SliceWriter | Actual written bytes: %v Expected written bytes: %v\n", num_bytes, expected_bytes)
	}

	num_bytes, err = LogOtusEvent(homeworkAccept, writer)
	expected_bytes = len(homeworkAccept.PrintEvent())
	if err != nil {
		t.Errorf("Write HwAccepted in SliceWriter | Actual error: %v Expected error: %v\n", err, nil)
	}
	if num_bytes != expected_bytes {
		t.Errorf("Write HwAccepted in SliceWriter | Actual written bytes: %v Expected written bytes: %v\n", num_bytes, expected_bytes)
	}

	//Read file
	file, err := os.Open(fileName)
	if err != nil {
		t.Errorf("Open file with name: %v | Actual error: %v Expected error: %v", fileName, err, nil)
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)

	if err != nil {
		t.Errorf("Reading file %v | Actual error: %v Expected error: %v", fileName, err, nil)
	} else {
		t.Logf("File content:\n%v", string(bytes))
	}
}
