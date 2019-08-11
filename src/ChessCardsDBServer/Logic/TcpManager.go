package tcpManager

import (
	ccg_loginLogic "ChessCardsDBServer/Logic/Login"
	"fmt"
	"strings"
)

func ImplementMessage(Mes string) (bool, string) {

	//fmt.Println(strings.Split("a,b,c,d,e", ",")) //[a b c d e]
	MesInfo := strings.Split(Mes, ",") //[a b c d e]
	var target string                  // 用来获取最终需要的os.Args[1]
	for k, v := range MesInfo {
		if k == 0 { // 假设需要获取os.Args[k], k = 1
			target = v
		}
	}

	if target != "" { // 不为空则表示MesInfo[1]存在
		if target == "login" {
			return ccg_loginLogic.ImplementLogic(Mes)
		}
	} else {
		fmt.Printf("tcpManager | Clinet Empty Message")
		return false, "nologic"
	}
	//fmt.Printf("tcpManager | Clinet Empty Message")
	return false, "nothing happend"
}
