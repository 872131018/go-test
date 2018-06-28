package main

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
	"time"
)

func Handle(attempts chan attempt, o *map[int]string) {
	for a := range attempts {
		(*o)[a.a] = fmt.Sprintf("Agent %d - successes: %d - failures: %d - duration: %v", a.a, a.s, a.e, a.l)
	}
}

func Display(o map[int]string) {
	// create timer to refresh display
	t := time.NewTicker(100 * time.Millisecond)

	// listen to the timer
	for _ = range t.C {

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
}
