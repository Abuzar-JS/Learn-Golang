package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func Unmarshal() {

	file, err := os.Open("datax.json")
	if err != nil {
		fmt.Println("Error occurs", err)
		return
	}

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("error x", err)
		return
	}

	var compi []Company

	err = json.Unmarshal(data, &compi)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	for _, x := range compi {
		fmt.Println("Company Name: ", x.Companies)
		fmt.Println("---------")

		fmt.Println("Address Street: ", x.Address.Street)
		fmt.Println("Address City: ", x.Address.City)
		fmt.Println("Address State: ", x.Address.State)
		fmt.Println("Address Zipcode: ", x.Address.Zipcode)
		for _, x := range x.Departments {
			fmt.Println("-----------------")
			fmt.Println("Departments: ", x.D_name)
			fmt.Println("Head Name: ", x.Head.H_name)
			fmt.Println("Head Position: ", x.Head.H_position)

			for _, x := range x.Employees {
				fmt.Println("Employee ID: ", x.E_id)
				fmt.Println("Employee Name: ", x.E_name)
				fmt.Println("Employee Position: ", x.E_position)
				fmt.Println(x.E_skills)

				fmt.Println("-----------------")
			}
			for _, x := range x.Projects {
				fmt.Println("Project ID: ", x.Projectid)
				fmt.Println("Project Name: ", x.P_name)
				fmt.Println("Project Description: ", x.Description)
				fmt.Println("Project Team: ", x.Team)
				fmt.Println("Project Status: ", x.Status)
			}

		}
	}
}
