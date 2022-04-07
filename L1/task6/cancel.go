package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

/**
Реализовать все возможные способы остановки выполнения горутины
*/

func doWithCancel(ctx context.Context, w *sync.WaitGroup, url string, c chan int) {
	defer w.Done()
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req = req.WithContext(ctx)
	if err != nil {
		fmt.Println("1:", err.Error())
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("2:", err.Error())
	}

	select {
	case <-time.After(800 * time.Millisecond):
		fmt.Println("after")
		c <- res.StatusCode
		return
	case <-ctx.Done():
		c <- http.StatusBadRequest
		fmt.Println("done")
		return
	}

}

func main() {
	c := make(chan int)
	w := &sync.WaitGroup{}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	urls := []string{"https://google.com", "https://yandex.ru", "https://seasonkrasoty.ru"}

	for _, url := range urls {
		w.Add(1)
		go doWithCancel(ctx, w, url, c)
	}

	time.Sleep(100 * time.Millisecond)

	go func() {
		defer w.Done()
		for status := range c {
			if status == http.StatusOK {
				fmt.Println("canceled")
				cancel()
			}
			fmt.Println(status)
		}
	}()
	w.Wait()
}

//func do(ctx context.Context) {
//	for {
//		time.Sleep(time.Second)
//		select {
//		case <-ctx.Done():
//			fmt.Println("done")
//			return
//		default:
//			fmt.Println("some work")
//		}
//	}
//}

//func f() {
//	err := errors.New("cancel")
//	ctx := context.Background()
//	ctxWithCancel, cancel := context.WithCancel(ctx)
//	go do(ctxWithCancel)
//
//	time.Sleep(5*time.Second)
//	fmt.Println(err)
//	cancel()
//	time.Sleep(10*time.Millisecond)
//}
