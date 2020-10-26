package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"
)

// Домашнее задание
// Параллельное исполнение
// Написать функцию для параллельного выполнения N заданий (т.е. в N параллельных горутинах).

// Функция принимает на вход:
// - слайс с заданиями `[]func() error`;
// - число заданий которые можно выполнять параллельно (`N`);
// - максимальное число ошибок после которого нужно приостановить обработку.

func main() {
	funcs := make([]func() error, 0)
	for i := 0; i < 100; i++ {
		f := func() error {
			num := rand.Intn(10)
			fmt.Printf("Task num=%d\n", num)
			if num > 5 {
				return errors.New("Error of task")
			}
			time.Sleep(time.Duration(num) * time.Second)
			return nil
		}
		funcs = append(funcs, f)
	}
	goWork(funcs, 5, 2)
}

func goWork(funcs []func() error, n int, errorMax int) {
	runtime.GOMAXPROCS(n)
	ch := make(chan bool, errorMax)
	for _, f := range funcs {
		go func(ch chan<- bool, fu func() error) {
			err := fu()
			if err != nil {
				ch <- true
			}
		}(ch, f)
	}
	result := false
	for i := 0; i < errorMax; i++ {
		result = <-ch
		if result {
			fmt.Printf("Get task with error %d\n", i)
			result = false
		}
	}
	close(ch)
	os.Exit(0)
}
