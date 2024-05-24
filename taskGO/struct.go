package main

type Employees struct {
	e_id        int
	name        string
	departments string
	skills      []string
	projects    []Projects
}

type Projects struct {
	name_p   string
	duration int
	team     []Team
}

type Team struct {
	t_id int
	role string
}

type clients struct {
	c_id     int
	name     string
	industry string
	contact  []contact
}

type contact struct {
	c_name string
	c_role string
}
