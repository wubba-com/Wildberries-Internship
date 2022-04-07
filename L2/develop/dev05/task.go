package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"
)

/**
5. Реализовать утилиту фильтрации по аналогии с консольной утилитой (man grep — смотрим описание и основные параметры).
*/

type Greper struct {
	output     io.Writer
	wg         *sync.WaitGroup
	lock       *sync.Mutex
	c          chan int
	lines      []string
	q          string
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	line       bool
	i          int
	amount     int
}

// contextN - высчитывает индексы для строк ДО и ПОСЛЕ найденной строки
func (g *Greper) contextN(i int) string {
	counter := g.context
	opposite := -counter
	for counter >= opposite {
		if i+counter < len(g.lines) && i-counter > -1 {
			g.addPreviousQ(g.wg, counter)
		}
		counter--
	}
	return g.q
}

// afterN - высчитывает индексы для строк ПОСЛЕ найденной строки
func (g *Greper) afterN(i int) string {
	var j = 0
	counter := g.after
	for counter >= 0 {
		if i+j < len(g.lines) {
			g.addNextQ(g.wg, j)
		}
		counter--
		j++
	}

	return g.q
}

// beforeN - высчитывает индексы для строк ДО найденной строки
func (g *Greper) beforeN(i int) string {
	counter := g.before
	for counter >= 0 {
		if i-counter > -1 {
			g.addPreviousQ(g.wg, counter)
		}
		counter--
	}

	return g.q
}

// add - Добавляет строку к итоговой строке поиска
func (g *Greper) add(index int) {
	if g.line {
		g.q += fmt.Sprintf("%d %s\n", index+1, g.lines[index])
	} else {
		g.q += fmt.Sprintf("%s\n", g.lines[index])
	}
}

// addCurrentQ - Добавляет текущую строку
func (g *Greper) addCurrentQ(index int) {
	g.add(index)
}

// addPreviousQ - добавляет предыдущую строку по N к итоговой строке поиска
func (g *Greper) addPreviousQ(wg *sync.WaitGroup, index int) {
	g.lock.Lock()
	defer g.lock.Unlock()
	index = g.i - index
	g.add(index)

}

// addNextQ - добавляет след строку по N к итоговой строке поиска
func (g *Greper) addNextQ(wg *sync.WaitGroup, index int) {
	g.lock.Lock()
	defer g.lock.Unlock()
	index = g.i + index
	g.add(index)
}

// worker - проверяет флаги и передает дальше функциям исполнителям
func (g *Greper) worker(i int) {
	g.amount += 1
	g.i = i // current string

	if g.before+g.after+g.context > 0 {
		if g.after > 0 {
			g.afterN(i)
		}

		if g.before > 0 {
			g.beforeN(i)
		}

		if g.context > 0 {
			g.contextN(i)
		}
	} else {
		g.addCurrentQ(i)
	}

}

// Grep - запускает алгоритм поиска
func (g *Greper) Grep(sub string) {
	fmt.Println(g.after, g.before, g.context)
	var reg *regexp.Regexp
	var err error
	if g.ignoreCase {
		reg, err = regexp.Compile("(?i)" + sub)
		sub = strings.ToLower(sub)
	} else {
		reg, err = regexp.Compile(sub)
	}

	if err != nil {
		log.Println(err.Error())
		return
	}

	for i, l := range g.lines {
		if g.ignoreCase {
			l = strings.ToLower(l)
		}
		if g.fixed {
			// Если g.invert = true, попадаем во второй if
			if l == sub && !g.invert {
				g.worker(i)
			}

			// противоположна пред-му if
			if l != sub && g.invert {
				g.worker(i)
			}
		} else {
			// Попадут подстроки
			if reg.MatchString(l) && !g.invert {
				g.worker(i)
			}

			// Попадут все кроме подстроки
			if !reg.MatchString(l) && g.invert {
				g.worker(i)
			}
		}
	}

	g.q = strings.Trim(g.q, "\n")

	if g.count {
		fmt.Fprintf(g.output, "%d", g.amountQ())
	} else {
		fmt.Fprintf(g.output, "%s", g.query())
	}
}

// Amount - Получаем сколько строк найдено
func (g *Greper) amountQ() int {
	return g.amount
}

// Query - Получаем итоговую строку поиска
func (g *Greper) query() string {
	return g.q
}

func writeFile(name string, b []byte) {
	err := ioutil.WriteFile(name, b, fs.ModePerm)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[err] %s", err.Error())
		return
	}
}

func readFile(scan *bufio.Scanner) []string {
	s := make([]string, 0)

	for scan.Scan() {
		s = append(s, scan.Text())
	}

	return s
}

var afterF = flag.Int("A", 0, "печатать +N строк после совпадения")
var beforeF = flag.Int("B", 0, "печатать +N строк до совпадения")
var contextF = flag.Int("C", 0, "(A+B) печатать ±N строк вокруг совпадения")
var countF = flag.Bool("c", false, "количество строк")
var ignoreCaseF = flag.Bool("i", false, "игнорировать регистр")
var invertF = flag.Bool("v", false, "вместо совпадения, исключать")
var fixedF = flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
var lineF = flag.Bool("n", false, "напечатать номер строки")

var fileName string
var sl []string
var q string

func main() {
	flag.Parse()
	fileName = flag.Arg(0)
	q = flag.Arg(1)

	// Открываем файл
	r, err := os.Open(fileName)
	defer r.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[err] %s\n", err.Error())
		return
	}

	// Создаем сканера, который считает данные по строчно и добавит их в слайс
	sc := bufio.NewScanner(r)
	sl = readFile(sc)

	// создаем структуру из переданных параметров и флагов пользователя
	g := Greper{
		output:     os.Stdout,
		lock:       &sync.Mutex{},
		wg:         &sync.WaitGroup{},
		lines:      sl,
		ignoreCase: *ignoreCaseF,
		after:      *afterF,
		before:     *beforeF,
		context:    *contextF,
		count:      *countF,
		invert:     *invertF,
		fixed:      *fixedF,
		line:       *lineF,
	}

	g.Grep(q)
}
