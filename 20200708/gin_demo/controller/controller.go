package controller

import (
	"gin_demo/manager/model"
	"gopkg.in/gin-gonic/gin.v1"

)

var (
	CommonControllerRepo   *CommonController

)

func Init() {
	CommonControllerRepo = new(CommonController)
}


//swagger:route /test
//测试api
//produces post:
//	- application/json
//responses:
//	200: test
func (this *CommonController) Test(ctx *gin.Context) {
	var data map[string]interface{}

	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmashal([]byte(string(body)), &data)
	if err != nil {
		log.Fatalln(err)
	}

	res := make(map[string]string)
	res["add"] = "ok"

	err = manager.AddCommon(data)
	if err != nil {
		res["add"] = error
	}
	
	ctx.JSON(200, res)
}
