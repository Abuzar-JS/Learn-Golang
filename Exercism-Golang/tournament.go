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

		team1, team2, status := details[0], details[1], details[2]

		UpdateScore(team1, status)
		UpdateScore(team2, reverseStatus(status))

	}
	printScores()

}

func UpdateScore(team, status string) Score {
	s := Team[team]
	s.MP++
	switch status {
	case "win":
		s.W++
		s.P += 3
	case "loss":
		s.L++
	case "draw":
		s.D++
		s.P++
	default:
		fmt.Println("Invalid status")
	}
	return s
}

func reverseStatus(status string) string {
	switch status {
	case "win":
		return "loss"
	case "loss":
		return "win"
	default:
		return status
	}
}

func printScores() {
	fmt.Printf("%-30s | %-3s | %-3s | %-3s | %-3s | %-3s\n", "Team", "MP", "W", "D", "L", "P")
	for team, score := range Team {
		fmt.Printf("%-30s | %-3d | %-3d | %-3d | %-3d | %-3d\n", team, score.MP, score.W, score.D, score.L, score.P)
	}
}
