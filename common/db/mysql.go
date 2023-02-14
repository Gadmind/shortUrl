package db

import (
	"github.com/jmoiron/sqlx"
	"log"
	"sync"
)

var (
	once     sync.Once
	dataBase *sqlx.DB
)

// MySQLDatabase 创建数据库连接
func MySQLDatabase() *sqlx.DB {
	once.Do(func() {
		dataBase = MysqlConn("root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
	})
	return dataBase
}

// MysqlConn 数据库链接创建
func MysqlConn(dataSource string) *sqlx.DB {
	var err error
	db, err := sqlx.Connect("mysql", dataSource)
	if err != nil {
		log.Fatalln(err, dataSource)
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		log.Fatalln(err, dataSource)
	}
	return db
}
