package main

import (
	"fmt"
)

func combineIps(sips, dips, dports []string) [][3]string {
	var res [][3]string
	for _, sip := range sips {
		for _, dip := range dips {
			for _, dport := range dports {
				var tmp [3]string
				tmp[0], tmp[1], tmp[2] = sip, dip, dport
				res = append(res, tmp)
			}
		}
	}
	return res
}

func main() {
	sips := []string{"10.1.1.1"}
	dips := []string{"10.2.2.2", "10.3.3.3"}
	dports := []string{"8888", "5556"}

	res := combineIps(sips, dips, dports)

	fmt.Println(res)
}
