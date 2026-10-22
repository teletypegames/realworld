package main

import (
	"errors"
	"fmt"

	inkwell "git.teletypegames.org/engines/inkwell"
)

var errUnsolvable = errors.New("the game cannot be finished")

func report(solution inkwell.Solution) error {
	fmt.Printf("moments reached: %d\n", solution.States)
	if solution.Truncated {
		fmt.Println("the search stopped at its limit")
	}
	for _, path := range solution.Runaway {
		fmt.Println("never settles:")
		printPath(path)
	}
	if len(solution.Unvisited) > 0 {
		fmt.Printf("scenes never reached: %v\n", solution.Unvisited)
	}
	if solution.DeadEnds > 0 {
		fmt.Printf("dead ends: %d, the first one after %d moves:\n", solution.DeadEnds, len(solution.DeadEnd))
		printPath(solution.DeadEnd)
		fmt.Println("what still changes the world there:")
		printPath(solution.DeadEndFrontier)
	}
	fmt.Printf("milestones: %v\n", solution.Milestones)
	if !solution.Solved {
		return errUnsolvable
	}
	fmt.Printf("solved in %d moves:\n", len(solution.Path))
	printPath(solution.Path)
	return nil
}

func printPath(path []inkwell.Move) {
	for index, move := range path {
		fmt.Printf("  %3d  %s\n", index+1, move)
	}
}
