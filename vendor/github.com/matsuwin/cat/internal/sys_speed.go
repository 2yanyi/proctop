package internal

import (
	"time"
)

func fibonacci(n int) int {
	if n > 1 {
		return fibonacci(n-1) + fibonacci(n-2)
	}
	return 1
}

func eating() {
	for i := 2; i <= 39; i++ {
		fibonacci(i)
	}
}

func processorSpeed() int {
	start := time.Now()
	eating()
	return int(time.Now().Sub(start).Milliseconds())
}
