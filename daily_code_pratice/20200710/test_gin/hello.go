package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type Area struct {
	Name	string
	BigCity string
}

type Acl struct {
	Type		string
	AreaName	[]Area
	AgeNow		int64
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

		fmt.Println(data)
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}

//测试接口
