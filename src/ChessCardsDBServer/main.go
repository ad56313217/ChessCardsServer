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
	"time"

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
	DB.SetConnMaxLifetime(100 * time.Second)
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(16)
	queryOne(DB)
	queryMulti(DB)
	insertData(DB)
	updateData(DB)
	deleteData(DB)

	if GoTest.IsHasTable("ChessCards", "test1", DB) == true {
		//GoTest.DeleteTable("test1", "test2", DB)
		fmt.Printf("Has Table")
	} else {
		GoTest.CreateTable("ChessCards", "test2", DB)
		fmt.Printf("Has no Table")
	}
}

//查询单行
func queryOne(DB *sql.DB) {
	user := new(User)
	row := DB.QueryRow("select * from users where id=?", 1)
	if err := row.Scan(&user.ID, &user.Name, &user.Age); err != nil {
		fmt.Printf("scan failed, err:%v", err)
		return
	}
	fmt.Println(*user)
}

//查询多行
func queryMulti(DB *sql.DB) {
	user := new(User)
	rows, err := DB.Query("select * from users where id > ?", 1)
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err != nil {
		fmt.Printf("Query failed,err:%v", err)
		return
	}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			fmt.Printf("Scan failed,err:%v", err)
			return
		}
		fmt.Print(*user)
	}

}

//插入数据
func insertData(DB *sql.DB) {
	result, err := DB.Exec("insert INTO users(name,age) values(?,?)", "hqx", 23)
	if err != nil {
		fmt.Printf("Insert failed,err:%v", err)
		return
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("Get lastInsertID failed,err:%v", err)
		return
	}
	fmt.Println("LastInsertID:", lastInsertID)
	rowsaffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v", err)
		return
	}
	fmt.Println("RowsAffected:", rowsaffected)
}

//更新数据
func updateData(DB *sql.DB) {
	result, err := DB.Exec("UPDATE users set age=? where id=?", "30", 3)
	if err != nil {
		fmt.Printf("Insert failed,err:%v", err)
		return
	}
	rowsaffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v", err)
		return
	}
	fmt.Println("RowsAffected:", rowsaffected)
}

//删除数据
func deleteData(DB *sql.DB) {
	result, err := DB.Exec("delete from users where id=?", 1)
	if err != nil {
		fmt.Printf("Insert failed,err:%v", err)
		return
	}
	rowsaffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v", err)
		return
	}
	fmt.Println("RowsAffected:", rowsaffected)
}
