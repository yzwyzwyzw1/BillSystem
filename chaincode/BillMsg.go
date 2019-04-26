package main

import (
	"encoding/json"
	"fmt"
)

type ChaincodeResponseMsg struct {
	Code int  `json:"Code"`// 0: successful 2 : failed
	Dec  string `json:"Dec"`//description of the message
}

func GetMsgString(code int,dec string) string{
	b, err := getMsg(code, dec)
    if err != nil {
    	return "Marshal failture"
	}
	
	return string(b)
}


func getMsg(code int,dec string) ([]byte,error) {
	var crm ChaincodeResponseMsg
	crm.Code =  code
	crm.Dec = dec

	b ,err := json.Marshal(crm)//序列化之后是
	if err != nil {
		return nil,fmt.Errorf(err.Error())
	}
	return b,nil

}