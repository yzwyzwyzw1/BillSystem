package main

import (
	"fmt"
	"github.com/BillSystem/blockchain"
	"github.com/BillSystem/service"
	"github.com/BillSystem/web"
	"github.com/BillSystem/web/controller"
)

const (
	CONFIGFILE="config.yaml"
	ORGADMIN="Admin"
	ORGNAME="Org1"
	CHANNELID="mychannel"
	CHANNELCONFIGPATH= "/home/yzw/GoSpace/gopath/src/github.com/BillSystem/fixtures/channel-artifacts/mychannel.tx"
	ORDERERORGNAME="orderer.example.com"
	CHAINCODEPATH="github.com/BillSystem/chaincode"
	CHAINCODEGOPATH="/home/yzw/GoSpace/gopath"
	CHAINCODEID="mycc"
	USERNAME="User1"
)


func FabricSetupInit() *blockchain.FabricSetup{
	f := blockchain.FabricSetup{
		ConfigFile:CONFIGFILE,
		Instantiated:false,
		OrgAdmin:ORGADMIN,
		OrgName:ORGNAME,
		ChannelID:CHANNELID,
		ChannelConfigPath:CHANNELCONFIGPATH,
		ChaincodePath:CHAINCODEPATH,
		ChaincodeGoPath:CHAINCODEGOPATH,
		ChaincodeID:CHAINCODEID,
		UserName:USERNAME,
		OrdererOrgName:ORDERERORGNAME,

	}
	return &f
}


func main() {

	//1. Initialize  fabric objection
	fab := FabricSetupInit()

	//2. Instantiate SDK  objection
	sdk,err := blockchain.InstantiateSdk(fab.ConfigFile,fab.Instantiated)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer sdk.Close()

	//3. Create Application Channel
	err = blockchain.CreateChannel(sdk,fab)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//4. Install and Instantiate chaincode
	channelClient, err := blockchain.InstallAndInstantiateCC(sdk, fab)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(channelClient)

	//5. Transmit channelClient object  to the  Service Layer
	serviceSetup:=&service.FabricSetupService{Fab:fab,ChannelClient:channelClient}

/*
	//==========================  业务层测试开始  ======================================//

	// 发布票据
	bill := service.Bill{
		BillInfoID:"BOC1001",
		BillInfoAmt:"20000",
		BillInfoType:"111",
		BillInfoIsseDate:"20180101",
		BillInfoDueDate:"201801011",

		DrwrAcct:"111",
		DrwrCmID:"111",
		AccptrAcct:"111",
		AccptrCmID:"111",

		PyeeAcct:"111",
		PyeeCmID:"111",
		HoldrAcct:"jack",
		HoldrCmID:"jackID",
	}

	bill2 := service.Bill{
		BillInfoID:"BOC1002",
		BillInfoAmt:"10000",
		BillInfoType:"111",
		BillInfoIsseDate:"20180101",
		BillInfoDueDate:"201801011",

		DrwrAcct:"111",
		DrwrCmID:"111",
		AccptrAcct:"111",
		AccptrCmID:"111",

		PyeeAcct:"111",
		PyeeCmID:"111",
		HoldrAcct:"jack",
		HoldrCmID:"jackID",
	}

	msg, err := serviceSetup.SaveBill(bill)
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println("票据发布成功, 交易编号为: " + msg)
	}

	msg, err = serviceSetup.SaveBill(bill2)
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println("票据发布成功, 交易编号为: " + msg)
	}

	// 查询当前持票人的票据列表
	result, err := serviceSetup.FindBills("jackID")
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println("根据持票人证件号码查询票据列表成功")
		var bills = []service.Bill{}
		json.Unmarshal(result, &bills)
		for _, obj := range bills{
			fmt.Println(obj)
		}
	}

	// 发起背书请求
	msg, err = serviceSetup.Endorse("BOC1001", "aliceID", "alice")
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println(msg)
	}

	// 查询待背书票据列表
	result, err = serviceSetup.FindWaitBills("aliceID")
	if err != nil{
		fmt.Println(err.Error())
	}else {
		fmt.Println("查询待背书票据列表成功")
		var bills = []service.Bill{}
		json.Unmarshal(result, &bills)
		for _, obj := range bills {
			fmt.Println(obj)
		}
	}

	// 签收票据
	msg, err = serviceSetup.Accept("BOC1001", "aliceID", "alice")
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println(msg)
	}

	// 根据票据号码查询票据详情
	result, err = serviceSetup.FindBillByNo("BOC1001")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var bill service.Bill
		json.Unmarshal(result, &bill)
		fmt.Println(bill)
	}

	//以下是 拒签票据测试
	// 发起背书
	msg, err = serviceSetup.Endorse("BOC1002", "aliceID", "alice")
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println(msg)
	}

	// 票据拒签
	msg, err = serviceSetup.Reject("BOC1002", "aliceID", "alice")
	if err != nil {
		fmt.Println(err.Error())
	}else{
		fmt.Println(msg)
	}

	result, err = serviceSetup.FindBillByNo("BOC1002")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var bill service.Bill
		json.Unmarshal(result, &bill)
		fmt.Println(bill)
	}

	//==========================  业务层测试完毕  ======================================//
*/

   // I advice to close comments when you test setup 6 ,beacuse we need not to test setup 5 when we test web layer.



	//6 Invoke WebServer to startup Web service
	app := controller.Application{
		Setup:serviceSetup,
	}
	web.WebStart(app)

}