package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)





type BillChaincode struct {

}
func (t *BillChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response  {
	return shim.Success(nil)
}

func (t *BillChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response  {

	// 获取用户意图
	fun,args := stub.GetFunctionAndParameters()
	if fun == "issue"{
		return t.issue(stub, args)
	}else if fun == "queryBills" {
		return t.queryBills(stub, args)		// 批量查询当前持票人的票据
	}else if fun == "queryBillByNo" {
		return t.queryBillByNo(stub, args)	// 根据票据号码查询票据详情
	}else if fun == "endorse" {
		return t.endorse(stub, args)		// 发起背书请求
	}else if fun == "queryWaitBills"{
		return t.queryWaitBills(stub, args)	// 查询待背书票据列表
	}else if fun == "accept" {
		return t.accept(stub, args)		    // 背书签收
	}else if fun == "reject" {
		return t.reject(stub, args)		    // 背书拒签
	}

	//如果输入的函数名错了
	respMsg := GetMsgString(1, "the name of function is error!")
	return shim.Error(respMsg)

}


func main() {
	
	err := shim.Start(new(BillChaincode))
	if err != nil {
		fmt.Println("chaincode startup failure:%v",err)
	}
	
	
}

