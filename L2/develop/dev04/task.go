package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"
)

/**
4. Написать функцию поиска всех множеств анаграмм по словарю.

Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому
*/

func Set(sl []string) []string {
	m := make(map[string]int)
	set := make([]string, 0)
	for _, v := range sl {
		m[v] += 1
	}

	for k := range m {
		if m[k] < 2 {
			set = append(set, k)
		}
	}

	return set
}
func IsAnagram(self, other string) bool {
	if utf8.RuneCountInString(self) != utf8.RuneCountInString(other) {
		return false
	}
	m := make(map[string]int)

	for _, v := range strings.Trim(strings.ToLower(self), " ") {
		m[string(v)]++
	}

	for _, v := range strings.Trim(strings.ToLower(other), " ") {
		m[string(v)] = m[string(v)] - 1
		if m[string(v)] < 0 {
			return false
		}
	}

	return true
}

func SetAnagram(anagrams []string) map[string][]string {
	m := make(map[string][]string)

	// Проверяем каждую строку с другими строками, является ли она анаграммой если да, она сохраняться в мапе
	for _, anagram := range anagrams {
		s := make([]string, 0)

		for _, word := range anagrams {
			if IsAnagram(anagram, word) {
				s = append(s, strings.ToLower(word))
				if len(s) > 1 {
					m[s[0]] = s[1:]
				}
			}
		}
	}

	for k := range m {
		sort.Strings(m[k])
	}

	return m
}

func main() {
	anagrams := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "Лунь", "нуль", "горечь"}
	fmt.Println(SetAnagram(anagrams))
}
