package main

import (
	"bufio"
	"fmt"
	"os"
)

// A. Shohag Loves Mod (https://codeforces.com/problemset/problem/2039/A)
// Time: 61ms
// Memory: 300KB

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tests int
	fmt.Fscan(reader, &tests)

	for i := 0; i < tests; i++ {
		var n int
		fmt.Fscan(reader, &n)

		for j := 1; j <= n; j++ {
			fmt.Fprintf(writer, "%d ", 2*j-1)
		}
		fmt.Fprintln(writer)
	}
}
