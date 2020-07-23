package main

import (
	"fmt"
)


type Teacher struct {
	Name	string
}

type Student struct {
	Age	int
}

func main() {
	tmp := make(map[*Teacher][]*Student)

	t1 := &Teacher{Name:"xiaohong"}


	s := &Student{Age:12}
	s1 := &Student{Age:11}
	tmp[t1] = append(tmp[t1], s)
	tmp[t1] = append(tmp[t1], s1)

	fmt.Println(tmp)
}
