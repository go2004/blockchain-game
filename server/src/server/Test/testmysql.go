package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, _ := sql.Open("mysql", "root:123@tcp(localhost:3306)/blockchain?parseTime=true")

	sqlStr := "SELECT id,`name`,style,app_coins,(SELECT count(DISTINCT app_coins) FROM accounts AS b WHERE a.app_coins<b.app_coins)+1 AS rank	FROM accounts AS a ORDER BY app_coins "
	rows, _ := db.Query(sqlStr)


	for rows.Next() {
		var playerID int
		var name string
		var style uint32
		var value float32
		var noID uint32
		_ = rows.Scan(&playerID, &name, &style, &value,&noID)

		fmt.Println(playerID,name,style,value,noID)
	}
}
