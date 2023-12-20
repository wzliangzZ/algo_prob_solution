package main

import (
	"bufio"
	"fmt"
)

func CD1914C(in *bufio.Reader, out *bufio.Writer) {
	var n, k int
	fmt.Fscan(in, &n, &k)
	arra := make([]int, n)
	arrb := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arra[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arrb[i])
	}
	var res, cur, mx, as int

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 0; i < n; i++ {
		as += arra[i]
		mx = max(mx, arrb[i])
		if k-(i+1) >= 0 {
			cur = as + (k-i-1)*mx
			res = max(res, cur)
		} else {
			break
		}
	}
	fmt.Fprintln(out, res)

}
