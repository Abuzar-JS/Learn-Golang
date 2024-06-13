package main

import "fmt"

// FunctionX is exported to main.go
func FunctionX() {
	//Function Body
	fmt.Println("Welcome To Tech Solutions Inc")

	var option int

	fmt.Println(`
	Please Select Option:
	1- Client List
	2- Employee List
	`)

	fmt.Scanf("%d", &option)

	p1 := Employees{

		eID:         1,
		name:        "John Doe",
		departments: "Engineering",
		skills:      []string{"Pyhton", "Javascript", "Machine Learning"},
		projects: []Projects{{
			nameP:    "Project X",
			duration: 6,
			team: []Team{{
				tID:  1,
				role: "Lead Developer",
			},
				{
					tID:  2,
					role: "Data Scientist",
				},
				{
					tID:  3,
					role: "UX Designer",
				},
			},
		}, {
			nameP:    "Project Y",
			duration: 9,
			team: []Team{{
				tID:  4,
				role: "Senior Developer",
			}, {
				tID:  5,
				role: "Database Admin",
			}},
		},
		},
	}

	p2 := Employees{
		eID:         2,
		name:        "Jane Smith",
		departments: "Marketing",
		skills:      []string{"Digital Marketing", "Content Creation", "SEO"},
		projects: []Projects{{
			nameP:    "Campaign A",
			duration: 4,
			team: []Team{{
				tID:  5,
				role: "Marketing Manager"}, {
				tID:  6,
				role: "Content Writer",
			},
			},
		},
		},
	}

	c1 := Clients{
		cID:      1,
		name:     "ABC Corporation",
		industry: "Finance",
		contact: []Contact{{
			cName: "Alice Johnson",
			cRole: "CTO",
		}, {
			cName: "Bob Williams",
			cRole: "CFO",
		},
		},
	}

	c2 := Clients{
		cID:      2,
		name:     "XYZ Enterprises",
		industry: "Retail",
		contact: []Contact{{
			cName: "Emily Davis",
			cRole: "Marketing Director",
		}, {
			cName: "Jack Brown",
			cRole: "CEO",
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

			fmt.Println("ID: ", c1.cID)
			fmt.Println("Company Name ", c1.name)
			fmt.Println("Industry: ", c1.industry)

			for _, v := range c1.contact {
				fmt.Println("---------")
				fmt.Println("Client Name: ", v.cName)
				fmt.Println("Role: ", v.cRole)
			}
		} else if option2 == 2 {
			fmt.Println("ID: ", c2.cID)
			fmt.Println("Company Name ", c2.name)
			fmt.Println("Industry: ", c2.industry)

			for _, v := range c2.contact {
				fmt.Println("---------")
				fmt.Println("Client Name: ", v.cName)
				fmt.Println("Role: ", v.cRole)
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
			fmt.Println("ID: ", p1.eID)
			fmt.Println("Name ", p1.name)
			fmt.Println("Department: ", p1.departments)
			for _, v := range p1.skills {
				fmt.Printf("Skill : %v \n", v)
			}

			for _, v := range p1.projects {
				fmt.Println("---------")
				fmt.Println("Project name: ", v.nameP)
				fmt.Println("Project Duration: ", v.duration)
				for _, v2 := range v.team {
					fmt.Println("---------")
					fmt.Println("Team ID: ", v2.tID)
					fmt.Println("Team Role: ", v2.role)

				}
			}
		} else if option3 == 2 {
			fmt.Println("ID: ", p2.eID)
			fmt.Println("Name ", p2.name)
			fmt.Println("Department: ", p2.departments)
			for _, v := range p2.skills {
				fmt.Printf("Skill : %v \n", v)
			}

			for _, v := range p2.projects {
				fmt.Println("---------")
				fmt.Println("Project name: ", v.nameP)
				fmt.Println("Project Duration: ", v.duration)
				for _, v2 := range v.team {
					fmt.Println("---------")
					fmt.Println("Team ID: ", v2.tID)
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
