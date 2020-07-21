package main

import "fmt"

func main() {
	s := "asdf"
	for _, k := range []byte(s) {
		fmt.Println(string(k))
	}
}
