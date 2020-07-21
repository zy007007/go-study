package model

import (
	"fmt"
	"time"
)

type Common struct {
	Id         int64
	Name       string
	Createtime time.Time
	Updatetime time.Time
}

var CommonRepo *Common = new(Common)

func (this *Common) TableName() string {
	return "common"
}


func (this *Common) Insert() (int64, error) {
	return Insert(this)
}
