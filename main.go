package main

import "fmt"

func main() {
	problemList := []string{
		"Day-1: Arrival of the General (https://codeforces.com/problemset/problem/144/A)",
		"Day-2: Xenia and Ringroad (https://codeforces.com/problemset/problem/339/B)",
		"Day-3: Vanya and Lanterns (https://codeforces.com/problemset/problem/492/B)",
		"Day-4: Sum of Round Numbers (https://codeforces.com/problemset/problem/1352/A)",
		"Day-5: Theatre Square (https://codeforces.com/problemset/problem/1/A)",
		"Day-6: Fair Division (https://codeforces.com/problemset/problem/1472/B)",
		"Day-7: Multiplication Dilemma (https://codeforces.com/problemset/gymProblem/101972/A)",
		"Day-8: Cyclic Shift (https://codeforces.com/gym/101972/problem/K)",
		"Day-9: Marathon (https://codeforces.com/problemset/problem/1692/A)",
		"Day-10: Lucky? (https://codeforces.com/problemset/problem/1676/A)",
		"Day-11: Quintomania (https://codeforces.com/problemset/problem/2036/A)",
	}

	fmt.Println("Hello! Welcome to my solutions of DIU CPC 99 Days of Problem Solving")
	fmt.Printf("\nDays completed: %d\n", len(problemList))

	fmt.Println("Problem List:")
	for _, problem := range problemList {
		fmt.Println(problem)
	}
}
