package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"server/conf"
	"github.com/name5566/leaf/log"
)

var (
	db *sql.DB
)

func OpenDB() {
	log.Release("mysqldb->open db")
	db1, err := sql.Open("mysql", conf.Server.MySqlUrl)
	if err != nil {
		db.Close()
		panic("connect db error")
	}
	db = db1
}

func MysqlDB()  *sql.DB {
	return db
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func PingDB(err error) {
	db :=MysqlDB()
	db.Ping()
}
