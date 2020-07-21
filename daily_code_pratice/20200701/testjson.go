package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("vim-go")

	//var str string = "{'ss':'123', 'xx':'231'}"
	//fmt.Println(str)

	jsonStr := `{"host":"10.1.1.1", "port":[8848,8888], "suborder":[{"acl_type":"IDC-互联网", "business_area":"c3"}, {"acl_type":"安全业务", "business_area":"c4"}]}`
	d := make(map[string]interface{})

	json.Unmarshal([]byte(jsonStr), &d)
	fmt.Println(d)
}
