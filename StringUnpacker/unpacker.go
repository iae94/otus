package main

import (
	"fmt"
	"log"
	"strconv"
	"unicode"
)

func Unpack(s string) string {
	var slashRune rune = 92   //Руна слеша
	result := make([]rune, 0) //Сюда будем складывать результат
	runes := []rune(s)        //Работаем со входящей строкой как со слайсом рун
	isEscaped := false        //Экранирован ли текущий символ в цикле

	for i, r := range runes {
		number := int(r - '0')
		isDigit := unicode.IsDigit(r) //Проверяем что текущая руна это цифра
		isSlash := r == slashRune     //Проверка на слеш
		isPrevDigit := false
		isNextDigit := false
		isPrevResultDigit := false

		if len(result) > 0 {
			isPrevResultDigit = unicode.IsDigit(result[len(result)-1]) //Цифра ли последняя руна в результирующем слайсе
		}
		if i > 0 && i+1 < len(runes) {
			isPrevDigit = unicode.IsDigit(runes[i-1]) //Цифра ли предыдущая руна
			isNextDigit = unicode.IsDigit(runes[i+1]) //Цифра ли следующая руна
		}
		switch {

		case i == 0 && isDigit: //Если первая руна - цифра = выход
			return ""
		case isSlash && !isEscaped: // Если текущая руна - неэкранируемый слеш - значит следующая руна - экранируемый символ
			isEscaped = true
			continue
		case isSlash && isEscaped: // Если текущая руна - экранируемый слеш - добавляем в результирующий слайс
			isEscaped = false
			result = append(result, r)
		case isDigit && !isEscaped: // Если текущая руна - неэкранируемая цифра - нужно повторить предыдущую руну
			if isPrevDigit && !isPrevResultDigit {
				continue
			}
			offset := 0
			if isNextDigit { // Если следующая руна тоже цифра - собираем все последующие цифры в одну и конвертим в инт
				fullNum := ""
				for ind, sr := range runes[i:] {
					isDigit := unicode.IsDigit(sr)
					if isDigit {
						fullNum += string(sr)
					}
					if !isDigit || ind == len(runes[i:])-1 {
						offset = len(fullNum) - 1
						break
					}
				}
				var convertError error
				number, convertError = strconv.Atoi(fullNum)
				if convertError != nil {
					log.Fatalf("Error during to convert string '%v' to integer '%v'", fullNum, convertError)
				}
			}
			last := result[len(result)-1] // Берем последнюю руну и дублируем num раз
			for j := 0; j < number-1; j++ {
				result = append(result, last)
			}
			return string(append(result, []rune(Unpack(string(runes[i+offset+1:])))...)) // После дублирования, рекурсивно обрабатываем остальную часть рун

		default:
			result = append(result, r) // По дефолту добавляем руну в результирующий слайс
			isEscaped = false
		}
	}
	return string(result)
}

func main() {
	type Tests struct {
		incoming string
		expected string
	}
	tests := []Tests{
		{"", ""},
		{"a4bc2d5e", "aaaabccddddde"},
		{"a15b11", "aaaaaaaaaaaaaaabbbbbbbbbbb"},
		{"abcd", "abcd"},
		{"a10b20", "aaaaaaaaaabbbbbbbbbbbbbbbbbbbb"},
		{"45", ""},
		{"012", ""},

		{`qwe\415a2`, `qwe444444444444444aa`},
		{`qwe\415\310\\\\\\3`, `qwe4444444444444443333333333\\\\\`},
		{`qwe\4\5`, `qwe45`},
		{`qwe\45`, `qwe44444`},
		{`qwe\\5`, `qwe\\\\\`},
		{`qwe\\\5`, `qwe\5`},
		{`qwe\\2\3\\2`, `qwe\\3\\`},
		{`\\`, `\`},
		{`\\\3\4\\`, `\34\`},
		{`\45q2w3e10`, `44444qqwwweeeeeeeeee`},
		{`\417\310`, `444444444444444443333333333`},
	}
	for _, test := range tests {
		unpackedString := Unpack(test.incoming)
		if test.expected == unpackedString {
			fmt.Printf("Incoming string: '%v' Expected string: '%v' Result: '%v' - %s\n", test.incoming, test.expected, unpackedString, "PASS")
		} else {
			fmt.Printf("Incoming string: '%v' Expected string: '%v' Result: '%v'\n - %s\n", test.incoming, test.expected, unpackedString, "FAILED")
		}
	}
}
