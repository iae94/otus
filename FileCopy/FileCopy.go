package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var from string
var to string
var offset int64
var limit uint

func init() {
	flag.StringVar(&from, "from", "", "source file to copy")
	flag.StringVar(&to, "to", "", "dest file to copy")
	flag.Int64Var(&offset, "offset", 0, "source file offset")
	flag.UintVar(&limit, "limit", 0, "limit of bytes to write")

}

func FileCopy(from string, to string, offset int64, limit uint) error{
	//Open source file
	reader, err := os.Open(from)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("Source file %v is not exist!\n", from)
			return err
		}
	}
	defer reader.Close()
	
	//Moving with Seeker
	_, err = reader.Seek(offset, io.SeekStart)
	if err != nil {
		log.Printf("Seek offset error %v\n", err)
		return err
	}

	//Open dest file
	writer, err := os.Create(to)
	if err != nil {
		log.Printf("Failed to open dest file %v %v\n", to, err)
		return err
	}
	defer writer.Close()
	
	//Create bufio Reader and Writer
	bufReader := bufio.NewReader(reader)
	bufWriter := bufio.NewWriter(writer)
	var readBytes uint = 0
	for readBytes < limit {

		oneByte, err := bufReader.ReadByte()
		if err == io.EOF {
			if readBytes == 0{
				err = errors.New("Offset is too high\n")
				log.Println(err)
				return err
			} else {
				break
			}
		}
		readBytes++
		err = bufWriter.WriteByte(oneByte)
		if err != nil{
			log.Printf("Cannot write byte %v\n", oneByte)
		} else {
			bufWriter.Flush()
		}
		if readBytes % 10 == 0{
			fmt.Printf("Copyied %v:%v bytes | Reader buf size: %v | Writer buf size: %v\n", readBytes, limit, bufReader.Buffered(), bufWriter.Buffered())
		}
	}
	fmt.Printf("File offset: %v bytes | Limit: %v bytes | Actual copyied: %v bytes\n", offset, limit, readBytes)
	
	return nil
}

func main() {
	flag.Parse()
	err := FileCopy(from, to, offset, limit)
	if err != nil {
		log.Printf("Copying error: %v\n", err)
	}
}
