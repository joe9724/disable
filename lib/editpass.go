package lib

import (
	"net/http"
	"io"
	"help/dao"
	"fmt"
)

//反馈
func EditPass(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	uid := r.URL.Query().Get("uid")
	oldpass := r.URL.Query().Get("oldpass")
	newpass := r.URL.Query().Get("newpass")

	var Uid string
	db := dao.GetConn()
	defer db.Close()
	fmt.Println("SELECT * from t_user where u_id =? and u_password=?",uid,oldpass)
	err := db.QueryRow("SELECT u_id from t_user where u_id =? and u_password=?",uid,oldpass).Scan(&Uid)
	if err!=nil {
		fmt.Println("erris",err.Error())
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Server", "NanjingYouzi")
		defer func() {
			io.WriteString(w,"error")
			//fmt.Print(string(data))
			//w.Write(data)
		}()
		return
	}else{
		//修改密码
		_,err1 := db.Exec(
			"UPDATE t_user set u_password=? where u_id=?",
			newpass,uid)
		fmt.Println("UPDATE t_user set u_password=? where u_id=?",
			newpass,uid)
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
}