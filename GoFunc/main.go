package main

func main() {
	p1 := Person{
		first: "James",
	}
	p2 := Person{
		first: "Jenny",
	}
	p3 := Person{
		first: "Scuba",
	}

	p1.speak()
	p2.speak()
	p3.speak()
}
