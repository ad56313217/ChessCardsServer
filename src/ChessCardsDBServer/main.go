/**
*FileName: mysql
*Create on 2018/7/17 下午4:57
*Create by mok
*golang中mysql的用法
 */

package main

import (
	GoTest "ChessCardsDBServer/DataTable"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   int            `db:"id"`
	Name sql.NullString `db:"name"`
	Age  int            `db:"age"`
}

const (
	USERNAME = "root"
	PASSWORD = "123456"
	NETWORK  = "tcp"
	SERVER   = "101.132.149.251"
	PORT     = 3306
	DATABASE = "ChessCards"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Open mysql failed,err:%v\n", err)
		return
	}

	CreateTest(DB)
}

//建表
func CreateTest(poDb *sql.DB) {
	GoTest.CreateTable("ChessCards", "zzh3", poDb)
}
