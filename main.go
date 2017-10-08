package main

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"help/lib"

	"fmt"
)


func main() {
	startHttpServer()
}

func startHttpServer() {
	//登录
	http.HandleFunc("/login", lib.Login)
	//签到
	http.HandleFunc("/sign", lib.Sign)
	//签到记录
	http.HandleFunc("/signlist", lib.Signlist)
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
	//explore
	http.HandleFunc("/explore",lib.Explore)
	//hash
	http.HandleFunc("/hash",lib.Hash)
	//update
	http.HandleFunc("/update",lib.Update)
	//delete
	http.HandleFunc("/delete",lib.Delete)
	//sign
	http.HandleFunc("/signal",lib.Signal)
	//init listener
	http.Handle("/img/", http.StripPrefix("/img", http.FileServer(http.Dir("/usr/img"))))
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	fmt.Println("server start success...")
}




