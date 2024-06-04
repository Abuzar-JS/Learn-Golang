package main

type comp struct {
	Company     string    `json:"company"`
	Address     []address `json:"address"`
	Departments []address `json:"departments"`
}

type address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zipcode int    `json:"zipcode"`
}

type departments struct {
	D_name    string      `json:"name"`
	Head      []head      `json:"head"`
	Employees []employees `json:"employees"`
	Projects  []projects  `json:"projects"`
}

type head struct {
	H_name     string `json:"name"`
	H_position string `json:"position"`
}
type employees struct {
	E_id       string   `json:"id"`
	E_name     string   `json:"name"`
	E_position string   `json:"position"`
	E_skills   []string `json:"skills"`
}

type projects struct {
	Projectid   string   `json:"projectId"`
	P_name      string   `json:"name"`
	Description string   `json:"description"`
	Team        []string `json:"team"`
	Status      string   `json:"status"`
}
