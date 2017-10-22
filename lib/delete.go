package lib

import (
	"net/http"
	//"log"
	"io"
	"fmt"

	"help/dao"


)

//删除采集
func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("client:", r.RemoteAddr, "method:", r.Method)
	//db := dao.GetConn()
	//defer db.Close()
	r.ParseForm()
	//r.ParseMultipartForm(32 << 20) //最大内存为32M

	//读取参数
	id := r.URL.Query().Get("id")
	db1 := dao.GetConn()
	defer db1.Close()
	var  Diff int
	err := db1.QueryRow("SELECT TIMESTAMPDIFF(HOUR,accessTime,now()) as diff from collection where id =?",id).Scan(&Diff)
	if err!=nil {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Server", "NanjingYouzi")
		defer func() {
			io.WriteString(w,"error")
			//fmt.Print(string(data))
			//w.Write(data)
		}()
		return
	}

	if Diff>=6 ||Diff<=-6 {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Server", "NanjingYouzi")
		defer func() {
			io.WriteString(w,"timeout")
			//fmt.Print(string(data))
			//w.Write(data)
		}()
		return
	}
	db := dao.GetConn()
	defer db.Close()

	//fmt.Println("id is",id)
	//fmt.Println("delete sql is","UPDATE collection set del=1 where id="+id)
	_,err1 := db.Exec(
		"UPDATE collection set del=1 where id=?",
		id)
	if err1 != nil {
		//log.Fatal("err post data2: ", err)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Server", "NanjingYouzi")
		fmt.Println("update err is",err.Error())
		defer func() {
			io.WriteString(w,"error")
			//fmt.Print(string(data))
			//w.Write(data)
		}()
		return
	}else{
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Server", "NanjingYouzi")
		defer func() {
			io.WriteString(w,"ok")
			//fmt.Print(string(data))
			//w.Write(data)
		}()
	}

}