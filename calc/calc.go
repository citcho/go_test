package calc

import "time"

func Add(a, b int) int {
	sum := a + b
	time.Sleep(3 * time.Second)
	return sum
}
