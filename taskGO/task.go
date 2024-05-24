package main

import (
	"fmt"
)

func Jsons() {

	p1 := Employees{

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


func Person1(){
	fmt.Println("ID: ", p1.e_id)
	fmt.Println("Name ", p1.name)
	fmt.Println("Department: ", p1.departments)
	for _, v := range p1.skills {
		fmt.Printf("Skill : %v \n", v)
	}

	for _, v := range p1.projects {
		fmt.Println("---------")
		fmt.Println("Project name: ", v.name_p)
		fmt.Println("Project Duration: ", v.duration)
		for _, v2 := range v.team {
			fmt.Println("---------")
			fmt.Println("Team ID: ", v2.t_id)
			fmt.Println("Team Role: ", v2.role)

		}
	}}

	var option int

	fmt.Scanln(&option)
	fmt.Println(" Chose your Option")
	// fmt.Println(p1.projects)

	// for _, r:= range Team.t_id {
	// 	fmt.Println()
	// }

	// p2 := Employees{
	// 	e_id:        2,
	// 	name:        "Jane Smith",
	// 	departments: "Marketing",
	// 	skills:      []string{"Digital Marketing", "Content Creation", "SEO"},
	// 	projects: []Projects{{
	// 		name_p:   "Campaign A",
	// 		duration: 4,
	// 		team: []Team{{
	// 			t_id: 5,
	// 			role: "Marketing Manager"}, {
	// 			t_id: 6,
	// 			role: "Content Writer",
	// 		},
	// 		},
	// 	},
	// 	},
	// }

	// c1 := Clients{
	// 	c_id:     1,
	// 	name:     "ABC Corporation",
	// 	industry: "Finance",
	// 	contact: []Contact{{
	// 		c_name: "Alice Johnson",
	// 		c_role: "CTO",
	// 	}, {
	// 		c_name: "Bob Williams",
	// 		c_role: "CFO",
	// 	},
	// 	},
	// }

	// c2 := Clients{
	// 	c_id:     2,
	// 	name:     "XYZ Enterprises",
	// 	industry: "Retail",
	// 	contact: []Contact{{
	// 		c_name: "Emily Davis",
	// 		c_role: "Marketing Director",
	// 	}, {
	// 		c_name: "Jack Brown",
	// 		c_role: "CEO",
	// 	}},
	// }
	// fmt.Println("--------------")
	// fmt.Println(p1, p2, c1, c2)
	// fmt.Println("--------------")

}
