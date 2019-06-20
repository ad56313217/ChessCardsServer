// test_005_tcp_svr project main.go
package main

import (
	"fmt"
	"net"
)

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
		fmt.Printf("服务器等待客户端 %s 发送信息\n", con.RemoteAddr().String())
		n, err := con.Read(buf)
		if err != nil {
			fmt.Println("客户端已退出,err:", err)
			return
		} else {
			//3.服务器显示客户端信息
			//fmt.Printf("收到了客户端（IP：%v）%d 个字节数据",con.RemoteAddr().String(),n)
			fmt.Printf("收到了客户端 %s 数据:%s ", con.RemoteAddr().String(), string(buf[:n]))

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
	fmt.Println("服务器开始监听...")
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
		fmt.Println("等待客户端连接")
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
