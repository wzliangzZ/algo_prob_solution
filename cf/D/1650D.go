package main

import (
	"bufio"
	"fmt"
	// "os"
)

func CF1650D(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}
	find := func(arr []int, t int) int {
		for i, v := range arr {
			if v == t {
				return i
			}
		}
		return -1
	}
	var res []int
	for i := n; i > 0; i-- {
		idx := find(arr, i)
		res = append(res, (idx+1)%i)
		arr = append(arr[idx+1:], arr[:idx]...)
	}
	for i := n - 1; i >= 0; i-- {
		fmt.Fprint(out, res[i], " ")
	}
	fmt.Fprintln(out)
}

/*
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
 
	var tot int = 1
	fmt.Fscan(in, &tot)
	for i := 0; i < tot; i++ {
		CF1650D(in, out)
	}
}
*/
