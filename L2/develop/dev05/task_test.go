package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"sync"
	"testing"
)

func TestGrep(t *testing.T) {
	inputs := []struct {
		afterN      int
		beforeN     int
		contextN    int
		count       bool
		ignoreCaseF bool
		invertF     bool
		fixedF      bool
		lineF       bool
		q           string
		res         []string
		file        string
	}{
		{
			afterN:      0,
			beforeN:     2,
			contextN:    0,
			count:       false,
			ignoreCaseF: false,
			invertF:     false,
			fixedF:      false,
			lineF:       false,
			file:        "input.txt",
			res: []string{
				"",
				"$templateData = array(",
				"        'TEMPLATE_THEME' => $arParams['TEMPLATE_THEME'],",
				"        '~COMPARE_URL_TEMPLATE' => $arResult['~COMPARE_URL_TEMPLATE'],",
				"        '~COMPARE_DELETE_URL_TEMPLATE' => $arResult['~COMPARE_DELETE_URL_TEMPLATE'],",
				"        'TEMPLATE_THEME' => $arParams['TEMPLATE_THEME'],",
			},
		},
		{
			afterN:      2,
			beforeN:     0,
			contextN:    0,
			count:       false,
			ignoreCaseF: false,
			invertF:     false,
			fixedF:      false,
			lineF:       false,
			file:        "input.txt",
			q:           "TEMPLATE_THEME",
			res: []string{
				"        'TEMPLATE_THEME' => $arParams['TEMPLATE_THEME'],",
				"        'TEMPLATE_LIBRARY' => $templateLibrary,",
				"        'CURRENCIES' => $currencyList",
				"        'TEMPLATE_THEME' => $arParams['TEMPLATE_THEME'],",
				"        'USE_ENHANCED_ECOMMERCE' => $arParams['USE_ENHANCED_ECOMMERCE'],",
				"        'DATA_LAYER_NAME' => $arParams['DATA_LAYER_NAME'],",
			},
		},
		{
			afterN:      0,
			beforeN:     0,
			contextN:    2,
			count:       false,
			ignoreCaseF: false,
			invertF:     false,
			fixedF:      false,
			lineF:       false,
			file:        "input.txt",
			q:           "TEMPLATE_THEME",
			res: []string{
				"$templateData = array(",
				"        'TEMPLATE_THEME' => $arParams['TEMPLATE_THEME'],",
				"        'TEMPLATE_LIBRARY' => $templateLibrary,",
				"        'CURRENCIES' => $currencyList",
				"        '~COMPARE_URL_TEMPLATE' => $arResult['~COMPARE_URL_TEMPLATE'],",
				"        '~COMPARE_DELETE_URL_TEMPLATE' => $arResult['~COMPARE_DELETE_URL_TEMPLATE'],",
				"        'TEMPLATE_THEME' => $arParams['TEMPLATE_THEME'],",
				"        'TEMPLATE_THEME' => $arParams['TEMPLATE_THEME'],",
				"        'DATA_LAYER_NAME' => $arParams['DATA_LAYER_NAME'],",
			},
		},
		{
			afterN:      0,
			beforeN:     0,
			contextN:    0,
			count:       false,
			ignoreCaseF: false,
			invertF:     true,
			fixedF:      false,
			lineF:       false,
			file:        "input2.txt",
			q:           "item-label-top",
			res: []string{
				"$positionClassMap = array(",
				"        'left' => 'product-item-label-left',",
				"        'center' => 'product-item-label-center',",
				"        'right' => 'product-item-label-right',",
				"        'bottom' => 'product-item-label-bottom',",
				"        'middle' => 'product-item-label-middle',",
				");",
			},
		},
	}

	for _, input := range inputs {
		r, err := os.Open(input.file)
		if err != nil {
			log.Fatalln(err)
		}

		// Создаем сканера, который считает данные по строчно и добавит их в слайс
		sc := bufio.NewScanner(r)
		sl = readFile(sc)

		g := &Greper{
			after:      input.afterN,
			before:     input.beforeN,
			context:    input.contextN,
			count:      input.count,
			ignoreCase: input.ignoreCaseF,
			invert:     input.invertF,
			fixed:      input.fixedF,
			line:       input.lineF,
			lines:      sl,

			output: os.Stdout,
			lock:   &sync.Mutex{},
			wg:     &sync.WaitGroup{},
		}
		g.Grep(input.q)

		query := strings.Split(g.query(), "\n")

		for i, v := range input.res {
			if query[i] == v {
				t.Error("result != value")
			}
		}
	}
}
