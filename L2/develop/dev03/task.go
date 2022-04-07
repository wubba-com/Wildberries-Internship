package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sortWithRegister(s []string) []string {
	sort.SliceStable(s, func(i, j int) bool {
		return strings.ToLower(sl[i]) < strings.ToLower(sl[j])
	})

	return s
}

func index(s string, w []string) int {
	for i, v := range w {
		if s == v {
			return i
		}
	}

	return -1
}

// readScan - возвращает слайс со строками из файла
func readScan(scan *bufio.Scanner) []string {
	s := make([]string, 0)

	for scan.Scan() {
		s = append(s, scan.Text())
	}

	return s
}

// sortUnique - сортирует и удаляет дубли
func sortUnique(sl []string) []string {

	set := make([]string, 0)

	for _, v := range sl {
		if index(v, set) < 0 {
			set = append(set, v)
		}
	}

	// возвращаем уже отсортированный слайс
	return sortWithRegister(set)
}

// sortReverse - сортирует в обратном порядке
func sortReverse(sl []string) []string {

	for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}

	// возвращаем уже отсортированный слайс
	return sl
}

// sortColumn - сортирует по выбранной колонке и по числовому значению
func sortColumn(lines []string, k int, n bool) []string {

	s := make([][]string, 0)

	k = k - 1
	if k < 0 {
		k = 0
	}

	for _, line := range lines {
		s = append(s, strings.Split(line, " "))
	}

	if n {
		sort.SliceStable(s, func(i, j int) bool {
			if len(s[i]) > k && len(s[j]) > k {
				x, err := strconv.Atoi(s[i][k])
				y, err := strconv.Atoi(s[j][k])
				if err != nil {
					fmt.Println(err)
					return false
				}

				return x < y
			}

			return false
		})
	} else {
		sort.SliceStable(s, func(i, j int) bool {
			if len(s[i]) > k && len(s[j]) > k {
				return strings.ToLower(s[i][k]) < strings.ToLower(s[j][k])
			}
			return false
		})
	}

	var str string
	sl = make([]string, 0)
	// строка файла которая была разделена пробелом, джониться обратно пробелом
	for _, line := range s {
		str = strings.Join(line, " ")
		sl = append(sl, str)
	}

	// возвращаем уже отсортированный слайс
	return sl
}

func unixSort(sl []string, flags *FlagsSort) []byte {
	sl = sortWithRegister(sl)

	// сортировка с удалением дублей
	if flags.unique {
		sl = sortUnique(sl)
	}

	// сортировка по колонке
	if flags.column > -1 {
		sl = sortColumn(sl, flags.column, flags.byName)
	}

	// сортировка в обратном порядке
	if flags.reverse {
		sl = sortReverse(sl)
	}

	return []byte(strings.Join(sl, "\n"))
}

const (
	DefaultColumnVal = -1
)

var fscan *bufio.Scanner
var fileName string
var column int
var byNum bool
var unique bool
var reverse bool
var sl []string

type FlagsSort struct {
	column  int
	reverse bool
	unique  bool
	byName  bool
}

func main() {
	flag.IntVar(&column, "k", -1, "указание колонки для сортировки")
	flag.BoolVar(&reverse, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&unique, "u", false, "не выводить повторяющиеся строки")
	flag.BoolVar(&byNum, "n", false, "сортировать по числовому значению")
	flag.Parse()

	fileName = flag.Arg(0)
	f, err := os.Open(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fl := &FlagsSort{unique: unique, column: column, reverse: reverse, byName: byNum}
	fscan = bufio.NewScanner(f)
	sl = readScan(fscan)

	err = ioutil.WriteFile(f.Name(), unixSort(sl, fl), fs.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}
