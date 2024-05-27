package main

type Company struct {
	Companies string `json:"company"`
	Employees []Employees
	Clients   []Clients
}

type Employees struct {
	Id         int      `json:"id"`
	Name       string   `json:"name"`
	Department string   `json:"department"`
	Skills     []string `json:"skills"`
	Projects   []Projects
}

type Projects struct {
	P_name     string `json:"name"`
	P_duration int    `json:"duration"`
	P_team     []Team `json:"team"`
}

type Team struct {
	T_id   int    `json:"id"`
	T_role string `json:"role"`
}

type Clients struct {
	C_id       int        `json:"id"`
	C_name     string     `json:"name"`
	C_industry string     `json:"industry"`
	C_Conatct  []Contacts `json:"contacts"`
}

type Contacts struct {
	Contacts_name string `json:"name"`
	Contacts_role string `json:"role"`
}
