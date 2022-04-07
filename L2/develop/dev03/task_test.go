package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"testing"
)

func TestUnixSort(t *testing.T) {
	tableTest := []struct {
		column  int
		reverse bool
		unique  bool
		byNum   bool
		file    string
		res     []string
	}{
		{
			column:  2,
			reverse: true,
			unique:  true,
			byNum:   true,
			file:    "text.txt",
			res: []string{
				"drwxr-xr-x 12 user user 4096 янв 14 21:49 Documents",
				"drwxr-xr-x 7 user user 4096 янв 13 11:42 Lightworks",
				"drwxr-xr-x 6 user user 4096 дек 6 14:29 Android",
				"drwx------ 5 user user 12288 янв 15 14:59 Downloads",
				"drwx------ 3 user user 4096 янв 14 22:18 Desktop",
			},
		},
		{
			column:  9,
			reverse: true,
			unique:  true,
			byNum:   false,
			file:    "text.txt",
			res: []string{
				"drwxr-xr-x 7 user user 4096 янв 13 11:42 Lightworks",
				"drwx------ 5 user user 12288 янв 15 14:59 Downloads",
				"drwxr-xr-x 12 user user 4096 янв 14 21:49 Documents",
				"drwx------ 3 user user 4096 янв 14 22:18 Desktop",
				"drwxr-xr-x 6 user user 4096 дек 6 14:29 Android",
			},
		},
	}

	for _, input := range tableTest {

		fl := &FlagsSort{
			column:  input.column,
			reverse: input.reverse,
			unique:  input.unique,
			byName:  input.byNum,
		}

		f, err := os.Open(input.file)
		if err != nil {
			log.Fatalln(err)
		}

		fscan = bufio.NewScanner(f)
		sl = readScan(fscan)

		sl = strings.Split(string(unixSort(sl, fl)), "\n")

		for i, v := range input.res {
			if sl[i] != v {
				t.Error("result != value")
			}
		}
	}

}
