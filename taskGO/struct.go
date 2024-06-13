package main

// Employees is exported to functions.go file
type Employees struct {
	eID         int
	name        string
	departments string
	skills      []string
	projects    []Projects
}

// Projects is exported to functions.go file
type Projects struct {
	nameP    string
	duration int
	team     []Team
}

// Team is exported to functions.go file
type Team struct {
	tID  int
	role string
}

// Clients is exported to functions.go file
type Clients struct {
	cID      int
	name     string
	industry string
	contact  []Contact
}

// Contact is exported to functions.go file
type Contact struct {
	cName string
	cRole string
}
