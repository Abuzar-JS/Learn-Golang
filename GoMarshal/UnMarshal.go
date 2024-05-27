package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func Unmarshal() {

	s, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	bs, err := io.ReadAll(s)
	if err != nil {
		return
	}

	// fmt.Println(string(bs))

	// bs := []byte(s)
	var comp []Company
	err = json.Unmarshal(bs, &comp)
	if err != nil {
		fmt.Println("The Error", err)
		return
	}

	// fmt.Println("All of the Data", comp)
	for _, v := range comp {
		fmt.Println("Company: ", v.Companies)
		fmt.Println("---------")
		if len(v.Companies) >= 1 {
			for _, v2 := range v.Employees {
				fmt.Printf("ID: %v \nName: %v \nDepartment: %v \n", v2.Id, v2.Name, v2.Department)
				if len(v2.Skills) >= 1 {
					fmt.Println("---Skills----")
					for _, v8 := range v2.Skills {

						fmt.Println(v8)
					}
					fmt.Println("---------")
				}
				if len(v2.Projects) >= 1 {
					for _, v3 := range v2.Projects {
						fmt.Println("---------")
						fmt.Printf("Project Name: %v Project Duration %v \n", v3.P_name, v3.P_duration)
						if len(v3.P_team) >= 1 {
							fmt.Println("---------")
							for _, v4 := range v3.P_team {
								fmt.Printf("Team ID: '%v' Role: %v \n", v4.T_id, v4.T_role)
							}
						}

					}
				}

				fmt.Println("------------------------------------------------")
			}

		}
		if len(v.Clients) >= 1 {
			for _, v := range v.Clients {
				fmt.Println("-----------------CLIENTS---------------------")
				fmt.Printf("ID: %v \nName: %v \nIndustry: %v \n", v.C_id, v.C_name, v.C_industry)

				if len(v.C_Conatct) >= 1 {
					for _, v := range v.C_Conatct {
						fmt.Println("---Client Contact---")
						fmt.Printf("Name: %v \nRole: %v \n", v.Contacts_name, v.Contacts_role)
					}
				}
			}
		}

	}
	defer s.Close()
}

//   // role => role_new , name => name_new
