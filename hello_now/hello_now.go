package main

import (
	"fmt"
	"github.com/beevik/ntp"
)

func PrintTime() {
	time, _ := ntp.Time("ntp3.stratum2.ru")
	fmt.Println("Time: ", time)
}

func main() {
	PrintTime()
}
