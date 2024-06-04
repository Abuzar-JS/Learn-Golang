package main

type Company struct {
	Companies   string        `json:"company"`
	Address     Address       `json:"address"`
	Departments []Departments `json:"departments"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zipcode string `json:"zipcode"`
}

type Departments struct {
	D_name    string      `json:"name"`
	Head      Head        `json:"head"`
	Employees []Employees `json:"employees"`
	Projects  []Projects  `json:"projects"`
}

type Head struct {
	H_name     string `json:"name"`
	H_position string `json:"position"`
}
type Employees struct {
	E_id       int      `json:"id"`
	E_name     string   `json:"name"`
	E_position string   `json:"position"`
	E_skills   []string `json:"skills"`
}

type Projects struct {
	Projectid   string   `json:"projectId"`
	P_name      string   `json:"name"`
	Description string   `json:"description"`
	Team        []string `json:"team"`
	Status      string   `json:"status"`
}
