package main

import "fmt"

func main() {
	problemList := []string{
		"Day-1: Arrival of the General (https://codeforces.com/problemset/problem/144/A)",
		"Day-2: Xenia and Ringroad (https://codeforces.com/problemset/problem/339/B)",
	}

	fmt.Println("Hello! Welcome to my solutions of DIU CPC 99 Days of Problem Solving")
	fmt.Printf("\nDays completed: %d\n", len(problemList))

	fmt.Println("Problem List:")
	for _, problem := range problemList {
		fmt.Println(problem)
	}

	day1()
}
