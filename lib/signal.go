package lib

import (
	"net/http"
	//"log"
	"io"
	"fmt"



)

//删除采集
func Signal(w http.ResponseWriter, r *http.Request) {
	fmt.Println("recieve signal:", r.RemoteAddr, "method:", r.Method)
	//db := dao.GetConn()
	//defer db.Close()
	r.ParseForm()
	//r.ParseMultipartForm(32 << 20) //最大内存为32M


		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Server", "NanjingYouzi")
		defer func() {
			io.WriteString(w,"ok")
			//fmt.Print(string(data))
			//w.Write(data)
		}()
	}
