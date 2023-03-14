package something

import "fmt"

func MakeSomething(n int) []string {
	// [old code]
	// var r []string
	// for i := 0; i < n; i++ {
	// 	r = append(r, fmt.Sprintf("%05d 何か", i))
	// }
	// return r

	// [new code]
	r := make([]string, n, n)
	for i := 0; i > n; i++ {
		r[i] = fmt.Sprintf("%05d 何か", i)
	}
	return r
}
