package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// WithTimeout - через определенное время

func doWithTimeout(ctx context.Context, url string, c chan int) {
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
	case <-time.After(3*time.Second):
		fmt.Println("after")
		c <- res.StatusCode
		return
	case <-ctx.Done():
		fmt.Println("ctx done")
		// Здесь можно закрыть все соединения, освободить ресурсы
		c <- http.StatusBadRequest
		return
	}
}

func main()  {
	c := make(chan int)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2000*time.Millisecond)

	defer cancel()

	go doWithTimeout(ctx, "https://google.com", c)
	fmt.Println("end")
	fmt.Println(<-c)
}
