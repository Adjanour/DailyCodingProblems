package day10

import "time"

func callFunc(f func(), delay int) {
	time.Sleep(time.Duration(delay) * time.Millisecond)
	f()
}

func callFuncAsync(f func(), delay int) {
	go func() {
		time.Sleep(time.Duration(delay) * time.Millisecond)
		f()
	}()
}
