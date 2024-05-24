package main

import "fmt"

func mainMulti() {
	jb := []string{"Abx", "abz", "xbz", "bzx", "bxz"}
	jc := []string{"Abk", "kba", "bka", "Akb", "kab"}
	jbc := [][]string{jb, jc}
	fmt.Println(jbc)
}
