package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

var Orm *xorm.Engine

func InitMysql() error {

	db_config := config.Config().Db
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		db_config.User, db_config.Passwd, db_config.Addr, db_config.Port, db_config.DBName)

	orm, err := xorm.NewEngine("mysql", conn)
	if err != nil {
		return err
	}


	err = orm.Ping()
	if err != nil {
		log.Fatal("连接数据库失败...")
		fmt.Println("连接数据库失败...")
		return err
	}
	fmt.Println("[Connecting MySQL Success!]")
	Orm = orm

	return nil
}

func Insert(beans ...interface{}) (int64, error) {
	return Orm.Insert(beans...)
}
