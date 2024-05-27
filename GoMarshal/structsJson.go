package main

type Company struct {
	Companies string      `json:"company"`
	Employees []Employees `json:"employees"`
	Clients   []Clients   `json:"clients"`
}

type Employees struct {
	Id         int        `json:"id"`
	Name       string     `json:"name_new"`
	Department string     `json:"department"`
	Skills     []string   `json:"skills"`
	Projects   []Projects `json:"projects"`
}

type Projects struct {
	P_name     string `json:"name_new"`
	P_duration int    `json:"duration"`
	P_team     []Team `json:"team"`
}

type Team struct {
	T_id   int    `json:"id"`
	T_role string `json:"role_new"`
}

type Clients struct {
	C_id       int        `json:"id"`
	C_name     string     `json:"name_new"`
	C_industry string     `json:"industry"`
	C_Conatct  []Contacts `json:"contacts"`
}

type Contacts struct {
	Contacts_name string `json:"name_new"`
	Contacts_role string `json:"role_new"`
}
