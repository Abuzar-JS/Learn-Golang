package main

import (
	"fmt"
	"time"
)

func E73() {
	timeFunc(DoWork)

}

func DoWork() {
	for i := 0; i < 2_000; i++ {
		fmt.Println(i)
	}
}

func timeFunc(f func()) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
