package main

import (
	"fmt"
	"strconv"
)

/**
2. Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например
*/

func RepeatS(s string) string {
	if len(s) == 0 {
		return ""
	}

	var repeatCh string
	isDigit := func(c rune) bool {
		return c >= 48 && c <= 57
	}

	for i, v := range s {
		// Если строка является численной строкой
		if isDigit(v) {
			// И берем предыдущий элемент который вскоре будем умножать
			if i == 0 {
				continue
			}

			c := s[i-1]

			// проверяем, что он не является численной строкой
			if isDigit(rune(c)) {
				continue
			}

			if string(c) == "\\" { // \
				c = s[i]
			}

			var n string
			// В цикле начинаем проверять является ли следующий символ числовой строки, да бы проверить ее в разряде десятков, тысяч итд.
			for _, v2 := range s[i:] {
				if isDigit(v2) {
					// создаем будущее число из строки
					n += string(v2)
					continue
				}

				break
			}
			// переводим строку в число
			x, err := strconv.Atoi(n)
			if err != nil {
				return ""
			}
			// повторяем букву на x
			for (x - 1) > 0 {
				repeatCh += string(c)
				x--
			}
		} else {
			// если это не число, то прибавляем к будущей новой строке
			repeatCh += string(v)
		}
	}
	return repeatCh
}

func main() {
	s := `a4bc2d5e`
	fmt.Println(RepeatS(s))
}
