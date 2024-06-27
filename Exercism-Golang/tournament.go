package main

//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strings"
//)
//
//package main
//
//import (
//"bufio"
//"fmt"
//"os"
//"strings"
//)
//
//// type Team struct {
//// 	Team1  string
//// 	Team2  string
//// 	Status string
//// }
//
//type Team struct {
//	MP int
//	W  int
//	D  int
//	L  int
//	P  int
//}
//
//var scores map[string]Team
//
//func mainData() {
//	file, err := os.Open("ReadFile.txt")
//	if err != nil {
//		fmt.Println("Error: ", err)
//		return
//	}
//	defer file.Close()
//
//	scanner := bufio.NewScanner(file)
//
//	for scanner.Scan() {
//		line := scanner.Text()
//
//		details := strings.Split(line, ";")
//
//		score := Team{
//			Team1:  details[0],
//			Team2:  details[1],
//			Status: details[2],
//		}
//		scores = append(scores, score)
//	}
//
//	for _, score := range scores {
//		fmt.Printf("Team1: %v, Team2: %v, Status: %v\n", score.Team1, score.Team2, score.Status)
//	}
//}
