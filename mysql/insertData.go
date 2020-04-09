package main

// 今天有个项目需要加管理员，是登录数据库手动加了
// 完成这个代码，下次直接go run xx.go username ... ...
// 同时go操作mysql相关初始连接以后也看参考此代码

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strconv"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func openDB() (success bool, db *sql.DB) {
	var isopen bool
	host := "10.x.x.x"
	user := "xxx"
	pwd := "xxx"
	database := "xxx"

	db, err := sql.Open("mysql", user+":"+pwd+"@tcp("+host+":port)/"+database+"?charset=utf8")
	if err != nil {
		isopen = false
	} else {
		isopen = true
	}
	checkErr(err)
	return isopen, db
}

func insertUserinfo(name, chinese_name, label string, level int, db *sql.DB) bool {
	_, err := db.Exec("INSERT INTO userinfo (username, chinese_name, label, level) VALUES (?,?,?,?)", name, chinese_name, label, level)
	if err != nil {
		return false
	}
	return true
}

func main() {
	opend, db := openDB()
	if opend {
		fmt.Println("connect db ok")
	} else {
		fmt.Println("connect db fail")
	}

	args := os.Args
	name, chinese_name, label := args[1], args[2], args[3]

	level, _ := strconv.Atoi(args[4])

	if insertUserinfo(name, chinese_name, label, level, db) {
		fmt.Println("insert ok")
	} else {
		fmt.Println("insert fail")
	}

}
