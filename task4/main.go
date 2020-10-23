package main

import (
	"fmt"
	"strings"
)

// Домашнее задание
// Частотный анализ
// Написать функцию, которая получает на вход текст и возвращает
// 10 самых часто встречающихся слов без учета словоформ
var text string = `test of of of of of test rest cast rest mest fest aest list glist sadist matist sa sa sa ma ma ma ma ma ma ma ma ru ru ru ru ru try try try pasta pasta pasta`

func main() {

	dict := make(map[string]int)
	words := strings.Fields(text)
	for _, w := range words {
		if _, ok := dict[w]; ok {
			dict[w]++
		} else {
			dict[w] = 1
		}
	}
	maxKeys := getTenMaxKeys(dict)
	for _, k := range maxKeys {
		fmt.Printf("key: %s, value: %d\n", k, dict[k])
	}
}

func getTenMaxKeys(dict map[string]int) []string {

	result := make([]string, 0)
	maxDict := make(map[string]int)
	for i := 0; i < 10 && i < len(dict); i++ {
		maxValue := 0
		maxKey := ""
		for k, v := range dict {
			if _, ok := maxDict[k]; !ok && v > maxValue {
				maxValue = v
				maxKey = k
			}
		}
		maxDict[maxKey] = maxValue
		result = append(result, maxKey)
	}
	return result
}
