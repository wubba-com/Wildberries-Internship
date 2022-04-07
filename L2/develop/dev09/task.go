package main

import (
	"flag"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

/**
9. Реализовать утилиту wget с возможностью скачивать сайты целиком.
*/

const (
	// Dir - имя директории, куда будут сохраняться html-файлы
	Dir = "pages"
)

// GetWithClient - функция-замыкание которая замыкает в себе http.Client
func GetWithClient(client *http.Client, fileName string) func(string) error {
	return func(url string) error {
		// создаем объект запроса
		r, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		// Отправляем запрос
		w, err := client.Do(r)
		defer func(Body io.ReadCloser) {
			err = Body.Close()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}(w.Body)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		fmt.Printf("status: %d", w.StatusCode)
		p, err := ioutil.ReadAll(w.Body)
		if err != nil {
			return err
		}

		// создаем имя файла из полного пути url

		err = write(fileName, p)
		if err != nil {
			return err
		}

		return nil
	}
}

// Write - пишет файл html, который получили из запроса
func write(fileName string, p []byte) error {
	// Getwd - возвращает полный путь до директории в которой находимся
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	// os.OpenFile с os.O_CREATE|os.O_WRONLY - откроет файл и запишет в него если такой имеется
	//или создаст новый файл. fs.ModePerm - задает права для файла
	f, err := os.OpenFile(filepath.Join(dir, Dir, fileName), os.O_CREATE|os.O_WRONLY, fs.ModePerm)
	defer f.Close()
	if err != nil {
		return err
	}

	return ioutil.WriteFile(f.Name(), p, fs.ModePerm)
}

var fileName = flag.String("O", "index.html", "файл в который будут сохранены полученные данные")

// https://seasonkrasoty.ru/product/legkiy_omolazhivayushchiy_krem/
func main() {

	flag.Parse()
	path := flag.Arg(0)
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}

	DownPage := GetWithClient(client, *fileName)
	err := DownPage(path)
	if err != nil {
		fmt.Println(err.Error())
	}
}
