package main

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"help/lib"

)


func main() {
	startHttpServer()
}

func startHttpServer() {
	//登录
	http.HandleFunc("/login", lib.Login)
	//签到
	http.HandleFunc("/sign", lib.Sign)
	//
	//http.HandleFunc("/collection", lib.Collection)
	//康复对象
	http.HandleFunc("/people", lib.People)
    //康复采集
	http.HandleFunc("/upload", lib.Upload)
	//搜索
	http.HandleFunc("/search", lib.Search)
	//我服务的
	http.HandleFunc("/mypeople", lib.MyPeople)
	//意见反馈
	http.HandleFunc("/feedback",lib.FeedBack)
	//init listener
	err := http.ListenAndServe(":82", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}




