package middleware

import (
	"log"
	"net/http"
	"time"
)

// Middleware Тип функции, которая оборачивает обработчик
type Middleware func(next http.HandlerFunc) http.HandlerFunc

// Log Ведет логи всех запросов с их путем и временем, которое потребовалось для обработки
var Log = func(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("http://%s %s [%s]", r.Host+r.URL.Path, time.Since(start).String(), r.Method)

		next.ServeHTTP(w, r)
	}
}

func Use(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}
