package main

import (
	"fmt"
	"math"
	"time"
)

func zexa() {
	fmt.Println("I'm function Z")
	fmt.Println(math.Sqrt(20))
	fmt.Println(time.Now())

}

func init() { //init function executes before the main func
	fmt.Println("Abuzar Here")
}
