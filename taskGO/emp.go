package main

import "fmt"

func PEmp() {
	p9 := Employees{
		e_id:        1,
		name:        "John Doe",
		departments: "Engineering",
		skills:      []string{"Pyhton", "Javascript", "Machine Learning"},
		projects: []Projects{{
			name_p:   "Project X",
			duration: 6,
			team: []Team{{
				t_id: 1,
				role: "Lead Developer",
			},
				{
					t_id: 2,
					role: "Data Scientist",
				},
				{
					t_id: 3,
					role: "UX Designer",
				},
			},
		}, {
			name_p:   "Project Y",
			duration: 9,
			
			team: []Team{{
				t_id: 4,
				role: "Senior Developer",
			}, {
				t_id: 5,
				role: "Database Admin",
			}},
		},
		},
	}

	for _, v := range p9.projects {
		fmt.Println(v)
		
	}