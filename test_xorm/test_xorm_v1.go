package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"fmt"
)

type Article struct {
	Id 	int64
	Title	string
	Words	string
}

// 如果结构体有Tablename方法，则orm使用时默认使用return名的数据库表，否则使用.Table()指定的
func(this *Article) TableName() string {
	return "article"
}

var orm *xorm.Engine
 
//创建orm引擎
func init() {
	var err error
	orm, err = xorm.NewEngine("mysql", "root:pwd@tcp(127.0.0.1:3306)/blog?charset=utf8")
	if err != nil {
		fmt.Println("数据库连接失败:", err)
	}
	fmt.Println("[Connecting MySQL Success!]")

}

func main() {
	// 新增一条数据
	var article Article
	article.Title = "测试数据，敬请期待，10.23.23.23"
	article.Words = "测试数据，敬请期待，10.32.32.32"
	res, err := orm.Insert(article)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("xorm insert 一条数据成功:", res)

	// 按id获取一个数据
	var article1 Article
	id := 29
	res1, err1 := orm.Where("id=?", id).Get(&article1)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println("xorm 按id获取一条数据成功:", res1, "data:", article1)

	// 按id更新一个数据
	res2, err2 := orm.ID(22).Update(&Article{Title:"更新数据，敬请期待，10.1.1.1"}) 
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("xorm 按id更新一个数据成功:", res2)

	// 按列获取数据　QueryString()
	title := "zhaoyi"
	res3, err3 := orm.Table("article").Where("title=?", title).QueryString()
	if err3 != nil {
		fmt.Println(err3)	
	}
	fmt.Println("xorm 按列获取数据成功", title, ":", res3)

	// 按列模糊获取数据　QueryString()
	title = "zhao"
	res4, err4 := orm.Table("article").Where("title like ?", "%"+title+"%").QueryString()
	if err4 != nil {
		fmt.Println(err4)
	}
	fmt.Println("xorm 模糊查询成功", title, ":", res4)

	// 获取最新一条数据id
	var lastid int64
	res5, err5 := orm.Limit(1).Table("article").Desc("id").Cols("id").Get(&lastid)
	if err5 != nil {
		fmt.Println(err5)
	}
	fmt.Println("xorm 获取最新一条数据id成功:", res5, "lastid:", lastid)
}

