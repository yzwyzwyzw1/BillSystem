package web

import (
	"net/http"
	"fmt"
	"github.com/BillSystem/web/controller"
)


//过程：指定一个登录页面，根据指定页面，进行解析，获取页面内容，转换成流，再发送给客户端

// 启动Web服务并指定路由信息
func WebStart(app controller.Application)  {


	fs:= http.FileServer(http.Dir("web/static"))  //以main函数表示为当前目录
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// 指定路由信息(匹配请求)
	http.HandleFunc("/", app.LoginView)
	http.HandleFunc("/login", app.Login)
	http.HandleFunc("/loginout", app.LoginOut)

	http.HandleFunc("/addBill", app.IssueShow)	// 显示发布票据页面
	http.HandleFunc("/issue", app.Issue)	// 提交发布票据请求

	http.HandleFunc("/bills", app.FindBills)	// 查询当前持票人的票据列表
	http.HandleFunc("/billInfo", app.BillInfoByNo)	// 根据票据号码查询票据详情

	http.HandleFunc("/endorse", app.Endorse)	// 发起背书请求

	http.HandleFunc("/waitEndorseBills", app.WaitEndorseBills)	// 查询当前用户的待背书票据列表
	http.HandleFunc("/waitEndorseInfo", app.WaitEndorseInfo)	// 根据票据号码查询待签收票据详情

	http.HandleFunc("/accept", app.Accept)	// 背书签收处理
	http.HandleFunc("/reject", app.Reject)	// 背书拒绝处理



	fmt.Println("启动Web服务, 监听端口号为: 9000")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("Web服务启动失败: %v", err)
	}

}

/**
	<link rel="" type="" href="/static/css/style.css">
	<script src="/static/js/main.js"></script>
 */
