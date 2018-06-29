package main

import (
	"fmt"
	"log"
)

func Prompt(projects []string) (int, int) {
	// give user options for testing
	for i, p := range projects {
		fmt.Printf("%d - %s\n", i, p)
	}

	// declare the list of projects that are supported
	fmt.Printf("Which project would you like to test?: ")
	var t int
	_, err := fmt.Scanf("%d", &t)
	if err != nil {
		log.Fatal("Must select a valid target to test")
	}

	// determine the number of agents to produce
	fmt.Printf("How many agents would you like to create? ")
	var a int
	_, err = fmt.Scanf("%d", &a)
	if err != nil {
		log.Fatal("Must have a valid number of agents")
	}

	return t, a
}
