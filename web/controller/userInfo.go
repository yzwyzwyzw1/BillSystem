package controller

import "github.com/BillSystem/service"

type Application struct {
	Setup *service.FabricSetupService
}

type User struct { //定义user数组
	LoginName	string
	Password	string
	CmID	string
	Acct	string
}


var users []User


func init() {

	admin := User{LoginName:"Admin", Password:"123456", CmID:"AdminID", Acct:"管理员"}//给驻足传入参数
	alice := User{LoginName:"alice", Password:"123456", CmID:"AliceID", Acct:"Alice"}
	bob := User{LoginName:"bob", Password:"123456", CmID:"BobID", Acct:"Bob"}
	jack := User{LoginName:"jack", Password:"123456", CmID:"JackID", Acct:"Jack"}

	users = append(users, admin)
	users = append(users, alice)
	users = append(users, bob)
	users = append(users, jack)

}
