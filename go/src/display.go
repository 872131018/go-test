package main

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
	"time"
)

func Handle(attempts chan attempt, o *map[int]string, rps *int) {
	for a := range attempts {
		(*o)[a.a] = fmt.Sprintf("Agent: %d - Successes: %d - Failures: %d - Duration: %v", a.a, a.s, a.e, a.l)
		*rps++
	}
}

func Display(o map[int]string, rps *int) {
	// create timer to refresh display
	t := time.NewTicker(1 * time.Second)

	// listen to the timer
	for _ = range t.C {

		//clear the screen
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()

		//hack to fix output
		fmt.Println("")

		//hack to fix output
		fmt.Printf("Requests per second: %d\n", *rps)

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

		// reset request counter
		*rps = 0
	}
}
