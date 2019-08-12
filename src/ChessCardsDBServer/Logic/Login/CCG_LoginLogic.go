package ccg_loginLogic

import (
	"ChessCardsDBServer/DataTable/CCG_Login"
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

const (
	USERNAME = "root"
	PASSWORD = "123456"
	NETWORK  = "tcp"
	SERVER   = "101.132.149.251"
	PORT     = 3306
	DATABASE = "ChessCards"
)

func ImplementLogic(Mes string) (bool, string) {

	var target string // 用来获取最终需要的os.Args[1]
	target = getArrayValueByStringMes(1, Mes)

	if target != "" { // 不为空则表示MesInfo[1]存在
		if target == "login" {
			return loginLogic(Mes)
		}
		if target == "register" {
			return registerLogic(Mes)
		}
		if target == "test" {
			return test(Mes), ""
		}
	} else {
		fmt.Printf("ccg_loginLogic | Clinet Empty Message")
		return false, "nologic"
	}
	//fmt.Printf("ccg_loginLogic | Clinet Empty Message")
	return false, "nothing happend"
}

func getArrayValueByStringMes(Index int, Mes string) string {

	MesInfo := strings.Split(Mes, ",") //[a b c d e]
	var target string                  // 用来获取最终需要的os.Args[1]
	for k, v := range MesInfo {
		if k == Index { // 假设需要获取os.Args[k], k = 1
			target = v
		}
	}

	return target
}

func loginLogic(Mes string) (bool, string) {

	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Open mysql failed,err:%v\n", err)
		return false, "mysql is failed"
	}

	//检测用户名是否存在
	info, _checkName := CCG_Login.FGetByUsrName("CCG_Login", DB, getArrayValueByStringMes(2, Mes))
	if _checkName == true {
		fmt.Sprintf("ccg_loginLogic | UserNameId: %d", info.Id)
		//return false, "UserName is exist!"
	} else {
		fmt.Printf("ccg_loginLogic | no userName")
		return false, "userName is not exist!"
	}

	if info.PassWord == getArrayValueByStringMes(3, Mes) {
		return true, "login success!"
	}
	return false, "ccg_loginLogic | password not true"
}

func registerLogic(Mes string) (bool, string) {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Open mysql failed,err:%v\n", err)
		return false, "mysql is failed"
	}

	//检测用户名是否存在
	info, _checkName := CCG_Login.FGetByUsrName("CCG_Login", DB, getArrayValueByStringMes(2, Mes))
	if _checkName == true {
		fmt.Sprintf("ccg_loginLogic | UserNameId: %d", info.Id)
		return false, "UserName is exist!"
	} else {
		fmt.Printf("ccg_loginLogic | no userName\n")
	}

	tableAll := CCG_Login.FGetAll("CCG_Login", DB)
	_id := tableAll[len(tableAll)-1].Id + 1
	//对表，进行相关的处理
	oIn := CCG_Login.CCG_LoginDb{
		Id:       _id,
		UsrName:  getArrayValueByStringMes(2, Mes),
		PassWord: getArrayValueByStringMes(3, Mes),
	}
	_b := CCG_Login.FInsToTbl("CCG_Login", DB, &oIn)
	if _b == false {
		//return
		fmt.Printf("write db faile!\n")
		return false, "write db faile!"
	}
	return true, "register success!"
}

func test(Mes string) bool {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Open mysql failed,err:%v\n", err)
		return false
	}

	info, _b := CCG_Login.FGetByUsrName("CCG_Login", DB, "11")
	if _b == true {
		fmt.Sprintf("ccg_loginLogic | UserNameId: %d\n", info.Id)
	} else {
		fmt.Printf("ccg_loginLogic | no userName\n")
	}
	return true
}
