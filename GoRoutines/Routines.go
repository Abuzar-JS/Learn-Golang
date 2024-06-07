package main

import (
	"fmt"
	"time"
)

func TalkwithFriend() {
	for i := 1; i <= 5; i++ {
		fmt.Println("I'm Talking with my friend", i)
		time.Sleep(1 * time.Second)
	}

}

func WriteAPage() {
	for i := 1; i <= 5; i++ {
		fmt.Println("I'm a Writing a Page", i)
		time.Sleep(1 * time.Second)
	}
}
