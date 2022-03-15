package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

var precision int = 4294967295
var inside int = 0
var size int = 4294967295
var threads int = runtime.NumCPU()

func getPi(total int) float64 {
	return (float64(inside) / float64(total)) * 4
}

func addpoints(times int, out chan int) {
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	var x, y, ret int
	ret = 0
	for i := 0; i < times; i++ {
		x = r1.Intn(size)
		y = r1.Intn(size)
		if ((x * x) + (y * y)) < (size * size) {
			ret += 1
		}
	}
	out <- ret
}

func main() {
	fmt.Println("running on " + fmt.Sprint(threads) + " threads")
	start := time.Now()
	results := make(chan int, threads)
	for i := 0; i < threads; i++ {
		go addpoints(precision/threads, results)
	}
	for i := 0; i < threads; i++ {
		inside += <-results
	}
	duration := time.Since(start)
	fmt.Println(getPi(precision))
	fmt.Println("compute time: " + fmt.Sprint(duration))
}
