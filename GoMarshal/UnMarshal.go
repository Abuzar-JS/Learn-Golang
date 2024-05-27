package main

import (
	"encoding/json"
	"fmt"
)

func Unmarshal() {

	s := `[{
		"company": "Tech Solutions Inc.",
		"employees": [
		  {
			"id": 1,
			"name": "John Doe",
			"department": "Engineering",
			"skills": ["Python", "JavaScript", "Machine Learning"],
			"projects": [
			  {
				"name": "Project X",
				"duration": 6,
				"team": [
				  {"id": 1, "role": "Lead Developer"},
				  {"id": 2, "role": "Data Scientist"},
				  {"id": 3, "role": "UX Designer"}
				]
			  },
			  {
				"name": "Project Y",
				"duration": 9,
				"team": [
				  {"id": 1, "role": "Senior Developer"},
				  {"id": 4, "role": "Database Administrator"}
				]
			  }
			]
		  },
		  {
			"id": 2,
			"name": "Jane Smith",
			"department": "Marketing",
			"skills": ["Digital Marketing", "Content Creation", "SEO"],
			"projects": [
			  {
				"name": "Campaign A",
				"duration": 4,
				"team": [
				  {"id": 5, "role": "Marketing Manager"},
				  {"id": 6, "role": "Content Writer"}
				]
			  }
			]
		  }
		],
		"clients": [
		  {
			"id": 1,
			"name": "ABC Corporation",
			"industry": "Finance",
			"contacts": [
			  {"name": "Alice Johnson", "role": "CTO"},
			  {"name": "Bob Williams", "role": "CFO"}
			]
		  },
		  {
			"id": 2,
			"name": "XYZ Enterprises",
			"industry": "Retail",
			"contacts": [
			  {"name": "Emily Davis", "role": "Marketing Director"},
			  {"name": "Jack Brown", "role": "CEO"}
			]
		  }
		]
	  }]`

	bs := []byte(s)
	var comp []Company
	err := json.Unmarshal(bs, &comp)
	if err != nil {
		fmt.Println(err)

	}

	// fmt.Println("All of the Data", comp)
	for _, v := range comp {
		fmt.Println("Company: ", v.Companies)
		fmt.Println("---------")
		for _, v2 := range v.Employees {
			fmt.Printf("ID: %v \nName: %v \nDepartment: %v \nSkills: %v  \n", v2.Id, v2.Name, v2.Department, v2.Skills)
			fmt.Println("---------")
			for _, v3 := range v2.Projects {
				fmt.Println("---------")
				fmt.Printf("Project Name: %v Project Duration %v \n", v3.P_name, v3.P_duration)
				fmt.Println("---------")
				for _, v4 := range v3.P_team {
					fmt.Printf("Team ID: '%v' Role: %v \n", v4.T_id, v4.T_role)
				}
			}

			fmt.Println("------------------------------------------------")
		}
		for _, v := range v.Clients {
			fmt.Println("-----------------CLIENTS---------------------")
			fmt.Printf("ID: %v \nName: %v \nIndustry: %v \n", v.C_id, v.C_name, v.C_industry)
			for _, v := range v.C_Conatct {
				fmt.Println("---Client Contact---")
				fmt.Printf("Name: %v \nRole: %v \n", v.Contacts_name, v.Contacts_role)
			}
		}

	}
}
