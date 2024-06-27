package main

//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strings"
//)
//
//// Define a Score struct to hold the score details
//type torn struct {
//	Team string
//	MP   int // Matches Played
//	W    int // Wins
//	D    int // Draws
//	L    int // Losses
//	P    int // Points
//}
//
//func main() {
//	file, err := os.Open("ReadFile.txt")
//	if err != nil {
//		fmt.Println("Error opening file:", err)
//
//		return
//	}
//	defer file.Close()
//
//	scanner := bufio.NewScanner(file)
//	scores := make(map[string]*Score)
//
//	for scanner.Scan() {
//		line := scanner.Text()
//		parts := strings.Split(line, ";")
//		if len(parts) < 3 {
//			continue
//		}
//
//		team1 := parts[0]
//		team2 := parts[1]
//		result := parts[2]
//
//		// Initialize or get the score struct for each team
//		if _, exists := scores[team1]; !exists {
//			scores[team1] = &Score{Team: team1}
//		}
//		if _, exists := scores[team2]; !exists {
//			scores[team2] = &Score{Team: team2}
//		}
//
//		// Update matches played
//		scores[team1].MP++
//		scores[team2].MP++
//
//		// Update wins, draws, losses, and points
//		switch result {
//		case "win":
//			scores[team1].W++
//			scores[team1].P += 3 // 3 points for a win
//			scores[team2].L++
//		case "loss":
//			scores[team1].L++
//			scores[team2].W++
//			scores[team2].P += 3 // 3 points for a win
//		case "draw":
//			scores[team1].D++
//			scores[team2].D++
//			scores[team1].P++
//			scores[team2].P++
//		}
//	}
//
//	if err := scanner.Err(); err != nil {
//		fmt.Println("Error reading file:", err)
//	}
//
//	// Print the results
//	for _, score := range scores {
//		fmt.Printf("Team: %s, MP: %d, W: %d, D: %d, L: %d, P: %d\n",
//			score.Team, score.MP, score.W, score.D, score.L, score.P)
//	}
//}
