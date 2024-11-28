package main

import (
	"bufio"
	"fmt"
	"os"
)

// A. Sliding (https://codeforces.com/problemset/problem/2035/A)
// Time: 61ms
// Memory: 600KB

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tests int
	fmt.Fscan(reader, &tests)

	for i := 0; i < tests; i++ {
		var n, m, r, c int

		fmt.Fscan(reader, &n, &m, &r, &c)

		fmt.Println((n-r)*(m-1) + (n * m) - (r-1)*m - c)
	}
}
