package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type Acl struct {
	Type	string
	Area	string
}

func main() {
	r := gin.Default()
	r.POST("/form", func(c *gin.Context) {
		var data Acl
	
		fmt.Println(c.Request.Body)
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Println(err)
		}

		err = json.Unmarshal([]byte(string(body)), &data)
		if err != nil {
			fmt.Println(err)
		}

		c.JSON(200, data)

		fmt.Println("after c.Json", data)
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}

//测试接口
//[GIN] 2020/07/05 - 20:28:38 | 200 |     105.538µs |       127.0.0.1 | POST     "/form"
//&{0xc0001f2e20 <nil> <nil> false true {0 0} false false false 0x6dcdd0}
//after c.Json map[acl类型:c3 业务类型:c4]


//curl 127.0.0.1:8080/form -d '{"acl类型":"c3", "业务类型":"c4"}'
//{"acl类型":"c3","业务类型":"c4"}[
