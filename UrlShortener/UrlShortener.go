package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type Shortener interface {
	Shorten(url string) string
	Resolve(url string) string
}

type MD5Shortener struct {
	Domain  string
	Storage map[string]string
}

func (s *MD5Shortener) Shorten(url string) string {
	hasher := md5.New()
	hasher.Write([]byte(url))
	shorten := s.Domain + hex.EncodeToString(hasher.Sum(nil))[:7]

	s.Storage[shorten] = url
	return shorten

}

func (s *MD5Shortener) Resolve(short_url string) string {
	if i, ok := s.Storage[short_url]; ok {
		return i
	} else {
		return ""
	}

}

func main() {

	url := "otus.ru/some-very-very-long-link"

	var shortener Shortener = &MD5Shortener{Domain: "otus.ru/", Storage: make(map[string]string)}
	short := shortener.Shorten(url)
	resolve := shortener.Resolve(short)

	fmt.Println("Shorten for:", url, "\tResult:", short)
	fmt.Println("Resolve for:", short, "\tResult", resolve)
}
