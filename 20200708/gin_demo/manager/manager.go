package manager

import (
	"fmt"
	"gin_demo/manager/model"
)

func init() {

}

func Start() {
	err := model.InitMysql()
	if err != nil {
		fmt.Printf("can't connect to db, get err:%s", err.Error())
	}
}

