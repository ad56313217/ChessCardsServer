/**
*FileName: mysql
*Create on 2018/7/17 下午4:57
*Create by mok
*golang中mysql的用法
 */

package main

import (
	CCG_Login "ChessCardsDBServer/DataTable/CCG_Login"
	"database/sql"
	"fmt"
	"net"

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

/*
func main() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Open mysql failed,err:%v\n", err)
		return
	}
	//GoTest.CreateDb("chesscards", DB)

	if CCG_Login.IsHasTable("ChessCards", "CCG_Login", DB) == false {
		if CCG_Login.CreateTable("ChessCards", "CCG_Login", DB) == false {
			fmt.Printf("CCG_Login  create table faile!")
			return
		}
		if CCG_LoginLog.CreateTable("ChessCards", "CCG_LoginLog", DB) == false {
			fmt.Printf("CCG_LoginLog  create table faile!")
			return
		}
		if CCG_LogoutLog.CreateTable("ChessCards", "CCG_LogoutLog", DB) == false {
			fmt.Printf("CCG_LogoutLog  create table faile!")
			return
		}
	}

	CreateTest(DB)
}

//建表
func CreateTest(poDb *sql.DB) {

}
*/

var ConnMap map[string]*net.TCPConn

func process(con net.Conn) {
	//循环接收客户端发送的数据
	defer con.Close() //关闭con

	for {
		//创建一个新的切片
		buf := make([]byte, 1024)

		//con.Read(buf)
		//1.等待客户端通过con发送信息
		//2.如果客户端没有write[发送]，协程就会阻塞于此
		fmt.Printf("Server wait for client %s send mes\n", con.RemoteAddr().String())
		n, err := con.Read(buf)
		if err != nil {
			fmt.Println("Client exit,err:", err)
			return
		} else {
			//3.服务器显示客户端信息
			//fmt.Printf("收到了客户端（IP：%v）%d 个字节数据",con.RemoteAddr().String(),n)
			fmt.Printf("From Client %s Data:%s ", con.RemoteAddr().String(), string(buf[:n]))

			dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
			DB, err := sql.Open("mysql", dsn)
			if err != nil {
				fmt.Printf("Open mysql failed,err:%v\n", err)
				return
			}
			//对表，进行相关的处理
			oIn := CCG_Login.CCG_LoginDb{
				Id:      2,
				UsrName: string(buf[:n]),
			}
			_b := CCG_Login.FInsToTbl("CCG_Login", DB, &oIn)
			if _b == false {
				//return
				fmt.Printf("write db faile!")
			}

			con.Write([]byte("456"))
		}

		// flag := checkErr(err)
		// if flag == 0 {
		// 	break
		// }
		for _, conn := range ConnMap {
			if conn.RemoteAddr().String() == con.RemoteAddr().String() {
				continue
			}
			conn.Write(buf[:n])
		}

	}

}

func main() {
	fmt.Println("Server Start Listion...")
	//1.tcp表示使用网络协议是tcp
	//2.0.0.0.0:8888表示在本地监听8888端口
	//lister, err := net.Listen("tcp","0.0.0.0:8888")
	tcpAddr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:8888")
	lister, err := net.ListenTCP("tcp", tcpAddr)

	if err != nil {
		fmt.Println("监听失败...err: ", err)
		return
	}

	defer lister.Close() //延时关闭listen

	ConnMap = make(map[string]*net.TCPConn)

	//循环等待客户端连接
	for {
		//等待客户端连接
		fmt.Println("Waiting for Client")
		//tcpConn, err := lister.Accept()
		tcpConn, err := lister.AcceptTCP()

		if err != nil {
			fmt.Printf("连接Accept() 失败，err: ", err)
		} else {
			fmt.Printf("Accept() suc conn=%v,客户端IP=%v\n", tcpConn, tcpConn.RemoteAddr().String())
		}

		ConnMap[tcpConn.RemoteAddr().String()] = tcpConn

		go process(tcpConn)
	}
	//fmt.Printf("lister=%v\n",lister)
}
