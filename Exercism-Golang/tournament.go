package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Score struct {
	MP int
	W  int
	D  int
	L  int
	P  int
}

var Team map[string]Score

func mainData() {

	Team = make(map[string]Score)
	file, err := os.Open("ReadFile.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		details := strings.Split(line, ";")

		team1 := details[0] // D
		team2 := details[1] // A
		status := details[2]

		if _, exists := Team[team1]; !exists {
			Team[team1] = GetStatus(status)
		} else {
			Team[team1] = UpdateScore(team1, status)
		}

		if _, exists := Team[team2]; !exists {
			if status == "win" {
				status = "loss"
			} else if status == "loss" {
				status = "win"
			}
			Team[team2] = GetStatus(status)
		} else {
			if status == "win" {
				status = "loss"
			} else if status == "loss" {
				status = "win"
			}
			Team[team2] = UpdateScore(team2, status)
		}

	}
	for team, score := range Team {
		fmt.Printf("%-30s | %-3d | %-3d | %-3d | %-3d | %-3d\n", team, score.MP, score.W, score.D, score.L, score.P)
	}
}

func GetStatus(status string) Score {

	p1 := Score{
		MP: 0,
		W:  0,
		D:  0,
		L:  0,
		P:  0,
	}
	p1.MP = 1
	if status == "win" {
		p1.W = 1
		p1.P = 3
	} else if status == "loss" {
		p1.L = 1
	} else {
		p1.D = 1
		p1.P = 1
	}

	return p1
}

func UpdateScore(team, status string) Score {
	p1 := Team[team]

	p1.MP += 1

	if status == "win" {
		p1.W += 1
		p1.P += 3
	} else if status == "loss" {
		p1.L += 1
	} else if status == "draw" {
		p1.D += 1
		p1.P += 1
	} else {
		fmt.Println("Invalid Status")
	}

	return p1
}
