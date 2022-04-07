package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Middleware Тип функции, которая оборачивает обработчик
type Middleware func(f http.HandlerFunc) http.HandlerFunc

// Logging Ведет логи всех запросов с их путем и временем, которое потребовалось для обработки
func Logging() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			// middleware
			start := time.Now()
			defer func() {
				log.Println("http://"+r.Host+r.URL.Path, time.Since(start))
			}()

			// Вызвать следующее промежуточное ПО/обработчик в цепочке
			next(w, r)
		}
	}
}

// Method middleware проверяет метод запроса
func Method(method string) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if method != r.Method {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			fmt.Printf("[%s] ", r.Method)
			next(w, r)
		}
	}
}

func SetHeaderText() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "text/html")
			next(w, r)
		}
	}
}

func SetCORS() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			next(w, r)
		}
	}
}

// Chain Устанавливает цепочку обязанностей для обработчика
func Chain(handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		handler = m(handler)
	}

	return handler
}

func main() {
	r := http.NewServeMux()

	r.HandleFunc("/welcome/",
		Chain(
			welcome,
			Logging(),
			Method("GET"),
			SetCORS(),
			SetHeaderText(),
		),
	)

	log.Fatal(http.ListenAndServe(":3000", r))
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hello server</h1>"))
}