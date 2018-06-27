package main

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
	"time"
)

func Display(output map[int]string) {
	// create timer to refresh display
	ticker := time.NewTicker(100 * time.Millisecond)

	// listen to the timer
	for _ = range ticker.C {

		//clear the screen
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()

		//hack to fix output
		fmt.Println("")

		// get the ids of the agents that have responded for sorting
		var keys []int
		for k := range output {
			keys = append(keys, k)
		}
		sort.Ints(keys)

		// print sorted list of agent responses
		for key := range keys {
			fmt.Println(output[key])
		}
	}
}
