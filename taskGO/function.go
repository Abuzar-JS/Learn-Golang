package main

import "fmt"

func FunctionX() {

	fmt.Println("Welcome To Tech Solutions Inc")

	var option int

	fmt.Println(`
	Please Select Option:
	1- Client List
	2- Employee List
	`)

	fmt.Scanf("%d", &option)

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

	p2 := Employees{
		e_id:        2,
		name:        "Jane Smith",
		departments: "Marketing",
		skills:      []string{"Digital Marketing", "Content Creation", "SEO"},
		projects: []Projects{{
			name_p:   "Campaign A",
			duration: 4,
			team: []Team{{
				t_id: 5,
				role: "Marketing Manager"}, {
				t_id: 6,
				role: "Content Writer",
			},
			},
		},
		},
	}

	c1 := Clients{
		c_id:     1,
		name:     "ABC Corporation",
		industry: "Finance",
		contact: []Contact{{
			c_name: "Alice Johnson",
			c_role: "CTO",
		}, {
			c_name: "Bob Williams",
			c_role: "CFO",
		},
		},
	}

	c2 := Clients{
		c_id:     2,
		name:     "XYZ Enterprises",
		industry: "Retail",
		contact: []Contact{{
			c_name: "Emily Davis",
			c_role: "Marketing Director",
		}, {
			c_name: "Jack Brown",
			c_role: "CEO",
		}},
	}

	if option == 1 {
		var option2 int

		fmt.Println(`
		Please Select Client:
		1- ABC Corporation
		2- XYZ Enterprises
		`)

		fmt.Scanf("%d", &option2)
		if option2 == 1 {

			fmt.Println("ID: ", c1.c_id)
			fmt.Println("Company Name ", c1.name)
			fmt.Println("Industry: ", c1.industry)

			for _, v := range c1.contact {
				fmt.Println("---------")
				fmt.Println("Client Name: ", v.c_name)
				fmt.Println("Role: ", v.c_role)
			}
		} else if option2 == 2 {
			fmt.Println("ID: ", c2.c_id)
			fmt.Println("Company Name ", c2.name)
			fmt.Println("Industry: ", c2.industry)

			for _, v := range c2.contact {
				fmt.Println("---------")
				fmt.Println("Client Name: ", v.c_name)
				fmt.Println("Role: ", v.c_role)
			}
		}

	} else if option == 2 {
		var option3 int

		fmt.Println(`
		Please Select Client:
		1- John Doe
		2- Jane Smith
		`)

		fmt.Scanf("%d", &option3)
		if option3 == 1 {
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
			}
		} else if option3 == 2 {
			fmt.Println("ID: ", p2.e_id)
			fmt.Println("Name ", p2.name)
			fmt.Println("Department: ", p2.departments)
			for _, v := range p2.skills {
				fmt.Printf("Skill : %v \n", v)
			}

			for _, v := range p2.projects {
				fmt.Println("---------")
				fmt.Println("Project name: ", v.name_p)
				fmt.Println("Project Duration: ", v.duration)
				for _, v2 := range v.team {
					fmt.Println("---------")
					fmt.Println("Team ID: ", v2.t_id)
					fmt.Println("Team Role: ", v2.role)

				}
			}
		} else {
			fmt.Println("You've Entered Wrong Option")
		}
	} else {
		fmt.Println("You've Entered Wrong Option")
	}
}
