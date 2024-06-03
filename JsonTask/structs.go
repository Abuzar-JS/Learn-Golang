package main

type comp struct {
	company string
	address []address
	departments []address
}


type address struct {
	street  string
	city    string
	state   string
	zipcode int
}

type departments struct {
	name string
	head []head
	employees []employees
	projects []projects
}

type head struct {
	name string
	position string
}
type employees struct{
	id string
	name string
	position string
	skills []string
}

type projects struct{
	projectid string
	p_name string
	description string
	team []string
	status string
}