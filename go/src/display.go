package main

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
)

func Display(o map[int]string) {
	//clear the screen
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()

	//hack to fix output
	fmt.Println("")

	// get the ids of the agents that have responded for sorting
	var keys []int
	for k := range o {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// print sorted list of agent responses
	for k := range keys {
		fmt.Println(o[k])
	}
}
