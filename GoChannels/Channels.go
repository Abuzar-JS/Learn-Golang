package main

import "fmt"

func Cha() {
	c := make(chan int, 4)
	c <- 42
	c <- 56
	c <- 20
	c <- 98

	fmt.Print("\n", <-c)
	fmt.Print("\n", <-c)
	fmt.Print("\n", <-c)
	fmt.Print("\n", <-c)

}

// Directional Channels

func DirectionalChannels() {
	c := make(chan int)
	cr := make(<-chan int) //Receive Only Channel
	cs := make(chan<- int) // Send Only Channel

	//Printing Type
	fmt.Println("---------Start---------")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)
	fmt.Println("---------End---------")

}

//Send Only Channel

func SendOnlyChannel() {
	// Send Only Channel can
	c := make(chan<- int, 2)
	c <- 22
	// move C before the Channel symbol syntax to fix and Execute
	fmt.Println(c)

}

func ReceiveOnlyChannel() {
	// Receive Only Channel can
	c := make(<-chan int, 2)

	fmt.Println(<-c)

}

func ExperimentChannel() {
	c := make(chan int)

	//send
	go foo(c)

	// receive
	go bar(c)

	fmt.Println("About to Exit")
}
func foo(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- 42
	}

} //Send
func bar(c <-chan int) {
	fmt.Println(<-c)
} //receive
