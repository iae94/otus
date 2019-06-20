package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

//Функция логирования Otus
//Задача: написать функцию логирования LogOtusEvent,
//на вход которой приходят события типа HwAccepted (домашняя работа принята) и HwSubmitted (студент отправил дз)
//для этого - спроектировать и реализовать интерфейс OtusEvent.
//Для события HwAccepted мы хотим логирровать дату, айди и грейд,
//для HwSubmitter - дату, id и комментарий, например:
//2019-01-01 submitted 3456 "please take a look at my homework"
//2019-01-01 accepted 3456 4

type HwAccepted struct {
	Date time.Time
	Id int
	Grade int
}
func (s HwAccepted) PrintEvent() []byte{
	return []byte(fmt.Sprintf("%v Accepted homework ID: %v Grade: %v\n", s.Date.Format("2006-01-02 15:04:05"), s.Id, s.Grade))
}

type HwSubmitted struct {
	Date time.Time
	Id int
	Code string
	Comment string
}
func (s HwSubmitted) PrintEvent() []byte{
	return []byte(fmt.Sprintf("%v Submitted homework ID: %v Comment: %v\n", s.Date.Format("2006-01-02 15:04:05"), s.Id, s.Comment))
}

type OtusEvent interface {
	PrintEvent() []byte
}

func LogOtusEvent(e OtusEvent, w io.Writer) (int, error) {
	num, err := w.Write(e.PrintEvent())
	return num, err
}


func main()  {
	writer1 := os.Stdout
	writer2 := os.Stderr

	fileName := "FileWriter.txt"
	writer3, err := os.Create(fileName)
	if err != nil {
		log.Printf("Create file with name: %v | Actual error: %v Expected error: %v", fileName, err, nil)
	}
	defer writer3.Close()

	submit1 := HwSubmitted{Date: time.Now(), Id: 5, Code: `func main(){ fmt.Println("Hello world") }`, Comment: "Please take a look at my homework"}
	submit2 := HwSubmitted{Date: time.Now(), Id: 934, Code: `func main(){ fmt.Println("Hello world") }`, Comment: "Please take a look at my homework"}
	submit3 := HwSubmitted{Date: time.Now(), Id: 69, Code: `func main(){ fmt.Println("Hello world") }`, Comment: "Please take a look at my homework"}

	accept1 := HwAccepted{Date: time.Now(), Id: 5, Grade:8}
	accept2 := HwAccepted{Date: time.Now(), Id: 934, Grade:4}
	accept3 := HwAccepted{Date: time.Now(), Id: 69, Grade:7}

	_, err = LogOtusEvent(submit1, writer1)
	if err != nil{
		log.Printf("Cannot write entry with ID: %v and type: 'submit' in writer %v | Error: %v", submit1.Id, writer1.Name(), err)
	}
	_, err = LogOtusEvent(accept1, writer1)
	if err != nil{
		log.Printf("Cannot write entry with ID: %v and type: 'accept' in writer %v | Error: %v", accept1.Id, writer1.Name(), err)
	}

	_, err = LogOtusEvent(submit2, writer2)
	if err != nil{
		fmt.Printf("Cannot write entry with ID: %v and type: 'submit' in writer %v | Error: %v\n", submit2.Id, writer2.Name(), err)
	}
	_, err = LogOtusEvent(accept2, writer2)
	if err != nil{
		fmt.Printf("Cannot write entry with ID: %v and type: 'accept' in writer %v | Error: %v\n", accept2.Id, writer2.Name(), err)
	}

	_, err = LogOtusEvent(submit3, writer3)
	if err != nil{
		log.Printf("Cannot write entry with ID: %v and type: 'submit' in writer %v | Error: %v", submit3.Id, writer3.Name(), err)
	}
	_, err = LogOtusEvent(accept3, writer3)
	if err != nil{
		log.Printf("Cannot write entry with ID: %v and type: 'accept' in writer %v | Error: %v", accept3.Id, writer3.Name(), err)
	}
}
