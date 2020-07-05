package main

import (
	"fmt"
)

// 1.获取一个表中最新的id
// 	table := "acl_order"
//	has, _ := Orm.Limit(1).Table(&table).Desc("id").Get(&lastId)

// 2.gin的ctx.JSON()是直接返回还是执行到函数结束

// 3, 按id更新一个interface
//func Update(id int64, beanPtr interface{}) (int64, error) {
//		return Orm.ID(id).Update(beanPtr)

// 4.按id获取一个数据
//	table := "acl_order"
//		has, _ := Orm.Table(&table).Where("id=?", id).Get(&res)

func main() {
	fmt.Println("vim-go")
}
