package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

/**
1. Создать программу печатающую точное время с использованием NTP -библиотеки. Инициализировать как go module.
Использовать библиотеку github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.
*/

func main() {
	res, err := ntp.Query("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
		return
	}
	// Validate метод выполняет дополнительные проверки работоспособности, чтобы определить, подходит ли ответ для целей синхронизации времени
	err = res.Validate()
	if err == nil {
		// ClockOffset: смещение часов локальной системы относительно часов сервера (+3 MSK)
		t := time.Now().Add(res.ClockOffset)
		fmt.Printf("%d:%d:%d", t.Hour(), t.Minute(), t.Second())
	}

}
