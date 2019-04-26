package controller

import (
	"net/http"
	"github.com/BillSystem/service"
	"fmt"
	"encoding/json"
)

var cuser User

func (app *Application) LoginView(w http.ResponseWriter, r *http.Request)  {  //

	ShowView(w, r, "login.html", nil)
}

// 用户登录
func (app *Application) Login(w http.ResponseWriter, r *http.Request) {
	loginName := r.FormValue("loginName")  //获取网页请求的表单信息
	password := r.FormValue("password")

	var flag bool
	for _, user := range users {
		if user.LoginName == loginName && user.Password == password {
			cuser = user
			flag = true
			break
		}
	}

	data := &struct {
		CurrentUser User
		Flag bool
	}{
		Flag:false,
	}

	if flag {
		// 登录成功
		r.Form.Set("holdrCmID", cuser.CmID)
		app.FindBills(w, r)
	}else{
		// 登录失败
		data.Flag = true
		data.CurrentUser.LoginName = loginName
		ShowView(w, r, "login.html", data)
	}
}

func (app *Application) LoginOut(w http.ResponseWriter, r *http.Request)  {
	cuser = User{}
	ShowView(w, r, "login.html", nil)
}



// 显示发布票据页面
func (app *Application) IssueShow(w http.ResponseWriter, r *http.Request)  {
	data := &struct {
		CurrentUser User
		Msg string
		Flag bool
	}{
		CurrentUser:cuser,
		Msg:"",//初始值为空
		Flag:false,
	}
	ShowView(w, r, "issue.html", data)
}

// 发布票据
func (app *Application) Issue(w http.ResponseWriter, r *http.Request)  {
	bill := service.Bill{
		BillInfoID:r.FormValue("billInfoID"),
		BillInfoAmt:r.FormValue("billInfoAmt"),
		BillInfoType:r.FormValue("billInfoType"),
		BillInfoIsseDate:r.FormValue("billInfoIsseDate"),
		BillInfoDueDate:r.FormValue("billInfoDueDate"),
		DrwrAcct:r.FormValue("drwrAcct"),
		DrwrCmID:r.FormValue("drwrCmID"),
		AccptrAcct:r.FormValue("accptrAcct"),
		AccptrCmID:r.FormValue("accptrCmID"),
		PyeeAcct:r.FormValue("pyeeAcct"),
		PyeeCmID:r.FormValue("pyeeCmID"),
		HoldrAcct:r.FormValue("holdrAcct"),
		HoldrCmID:r.FormValue("holdrCmID"),
	}

	transactionID, err := app.Setup.SaveBill(bill)

	data := &struct {   //这是什么语法格式啊
		CurrentUser User
		Msg string
		Flag bool
	}{
		CurrentUser:cuser,
		Flag:true,
		Msg:"",
	}

	if err != nil {
		data.Msg = err.Error()
	}else{
		data.Msg = "票据发布成功:" + transactionID
	}

	ShowView(w, r, "issue.html", data)

}

// 查询持票人的票据列表
func (app *Application) FindBills(w http.ResponseWriter, r *http.Request){
	//r.FormValue("holdrCmID")
	result, err := app.Setup.FindBills(cuser.CmID)
	if err != nil {
		fmt.Println(err.Error())
	}

	var bills = []service.Bill{}
	json.Unmarshal(result, &bills)

	data := &struct {
		Bills []service.Bill
		CurrentUser User
	}{
		Bills:bills,
		CurrentUser:cuser,
	}

	ShowView(w, r, "bills.html", data)
}

// 根据票据号码查询票据详情
func (app *Application) BillInfoByNo(w http.ResponseWriter, r *http.Request) {
	billNo := r.FormValue("billNo")
	result, _ := app.Setup.FindBillByNo(billNo)
	var bill = service.Bill{}
	json.Unmarshal(result, &bill)

	fmt.Println(billNo + " -> Controller: " + string(result))

	data := &struct {
		Bill service.Bill
		CurrentUser User
		Msg string
		Flag bool
	}{
		Bill:bill,
		CurrentUser:cuser,
		Msg:"",
		Flag:false,
	}

	if r.FormValue("msg") != "" {
		data.Msg = r.FormValue("msg")
		data.Flag = true
	}

	ShowView(w, r, "billInfo.html", data)
}



// 发起背书请求处理
func (app *Application) Endorse(w http.ResponseWriter, r *http.Request)  {
	billNo := r.FormValue("billNo")
	endorseWaitCmID := r.FormValue("waitEndorseCmID")
	endorseWaitAcct := r.FormValue("waitEndorseAcct")

	_, err := app.Setup.Endorse(billNo, endorseWaitCmID, endorseWaitAcct)

	/*data := &struct {
		Bill service.Bill
		Flag bool
		Msg string
		CurrentUser User
	}{
		Flag:true,
		Msg:result,
		CurrentUser:cuser,
	}
	ShowView(w, r, "billInfo.html", data)*/

	if err != nil {
		r.Form.Set("msg", err.Error())
	}else{
		r.Form.Set("msg", "背书请求成功, 待" + endorseWaitAcct + "签收")
	}

	r.Form.Set("billNo", billNo)


	app.BillInfoByNo(w, r)
}

// 查询当前用户的待背书票据列表
func (app *Application) WaitEndorseBills(w http.ResponseWriter, r *http.Request)  {
	waitEndorseCmID := cuser.CmID
	result, err := app.Setup.FindWaitBills(waitEndorseCmID)
	if err != nil {
		fmt.Println(err.Error())
	}

	var bills = []service.Bill{}
	json.Unmarshal(result, &bills)

	data := &struct {
		Bills []service.Bill
		CurrentUser User
	}{
		Bills:bills,
		CurrentUser:cuser,
	}

	ShowView(w, r, "waitEndorse.html", data)
}

// 根据票据号码查询待签收票据详情
func (app *Application) WaitEndorseInfo(w http.ResponseWriter, r *http.Request)  {

	billNo := r.FormValue("billNo")
	result, _ := app.Setup.FindBillByNo(billNo)

	var bill service.Bill
	json.Unmarshal(result, &bill)

	data := &struct {   //类型定义
		Bill service.Bill
		CurrentUser User
		Flag bool
		Msg string
	}{
		Bill:bill,   //赋值
		CurrentUser:cuser,
		Flag:false,
		Msg:"",
	}

	if r.FormValue("msg") != "" {
		data.Flag = true
		data.Msg = r.FormValue("msg")
	}

	ShowView(w, r, "waitEndorseInfo.html", data)

}

// 票据签收
func (app *Application) Accept(w http.ResponseWriter, r *http.Request) {
	billNo := r.FormValue("billNo")
	signedAcct := cuser.Acct
	signedCmID := cuser.CmID

	result, err := app.Setup.Accept(billNo, signedCmID, signedAcct)

	r.Form.Set("billNo", billNo)
	if err != nil {
		r.Form.Set("msg", err.Error())
	} else {
		r.Form.Set("msg", result)
	}

	app.WaitEndorseInfo(w, r)

}

// 票据背书拒绝
func (app *Application) Reject(w http.ResponseWriter, r *http.Request)  {
	billNo := r.FormValue("billNo")
	rejectAcct := cuser.Acct
	rejectCmID := cuser.CmID

	result, err := app.Setup.Reject(billNo, rejectCmID, rejectAcct)

	r.Form.Set("billNo", billNo)
	if err != nil {
		r.Form.Set("msg", err.Error())
	}else {
		r.Form.Set("msg", result)
	}

	app.WaitEndorseInfo(w, r)

}