package main

import (
	"fmt"
	"log"
)

func Prompt(projects []string) (int, int) {
	// give user options for testing
	for i, project := range projects {
		fmt.Printf("%d - %s\n", i, project)
	}

	// declare the list of projects that are supported
	fmt.Printf("Which project would you like to test?: ")
	var target int
	_, err := fmt.Scanf("%d", &target)
	if err != nil {
		log.Fatal("Must select a valid target to test")
	}

	// determine the number of agents to produce
	fmt.Printf("How many agents would you like to create? ")
	var agents int
	_, err = fmt.Scanf("%d", &agents)
	if err != nil {
		log.Fatal("Must have a valid number of agents")
	}

	return target, agents
}
