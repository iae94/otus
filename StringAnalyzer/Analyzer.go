package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type Word struct {
	word string
	frequency int
}

func FreqAnalyzer(s string) []Word {
	max_count := 10
	pattern := regexp.MustCompile("[^a-zA-Zа-яА-Я0-9]+")
	s = pattern.ReplaceAllString(strings.ToLower(s), " ")
	frequency_map := make(map[string]int, 0)
	result_words := make([]Word, 0)
	words := strings.Fields(s)

	for _, word := range words {
		frequency_map[word] += 1
	}
	for word, freq := range frequency_map {
		result_words = append(result_words, Word{word, freq})
	}
	sort.Slice(result_words, func(i, j int) bool {
		return result_words[i].frequency > result_words[j].frequency
	})
	if len(result_words) > max_count{
		return result_words[:max_count]
	} else {
		return result_words
	}
}

func main()  {

	text :=
		`
		Ежели вы не жили возле ежевичника,
		но ежели вы жили возле земляничника,
		то значит земляничное варенье вам привычное
		и вовсе не привычное варенье ежевичное.
		Ежели вы жили возле ежевичника, 
		то значит, ежевичное варенье вам привычное,
		и вовсе не привычное варенье земляничное. 
		Но ежели вы жили возле ежевичника,
		и ежели вы жили возле земляничника
		и ежели вы времени на лес не пожалели, 
		то значит, преотличное варенье ежевичное,
		варенье земляничное вы ежедневно ели.
		`
	words := FreqAnalyzer(text)
	for _, word := range words {
		fmt.Printf("Word: '%v' Frequency: %v\n", word.word, word.frequency)
	}
}
