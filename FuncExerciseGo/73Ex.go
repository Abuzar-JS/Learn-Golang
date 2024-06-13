package main

import (
	"fmt"
	"time"
)

// E73 exported
func E73() {
	timeFunc(DoWork)

}

// DoWork exported
func DoWork() {
	for i := 0; i < 3_000; i++ {
		fmt.Println(i)
	}
}

func timeFunc(f func()) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
