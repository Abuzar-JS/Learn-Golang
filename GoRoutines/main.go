package main

import "time"

func main() {
	go TalkwithFriend()
	go WriteAPage()

	time.Sleep(6 * time.Second)
}
