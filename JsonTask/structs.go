package main

type comp struct {
	company     string    `json:"company"`
	address     []address `json:"address"`
	departments []address `json:"departments"`
}

type address struct {
	street  string `json:"street"`
	city    string `json:"city"`
	state   string `json:"state"`
	zipcode int    `json:"zipcode"`
}

type departments struct {
	name      string      `json:"name"`
	head      []head      `json:"head"`
	employees []employees `json:"employees"`
	projects  []projects  `json:"projects"`
}

type head struct {
	name     string `json:"name"`
	position string `json:"position"`
}
type employees struct {
	id       string   `json:"id"`
	name     string   `json:"name"`
	position string   `json:"position"`
	skills   []string `json:"skills"`
}

type projects struct {
	projectid   string   `json:"projectId"`
	p_name      string   `json:"name"`
	description string   `json:"description"`
	team        []string `json:"team"`
	status      string   `json:"status"`
}
