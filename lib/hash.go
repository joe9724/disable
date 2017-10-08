package lib

import (
	"net/http"
	"encoding/json"
	"log"
	"io"
	"help/dao"
	"help/module"

)

//
func Hash(w http.ResponseWriter, r *http.Request) {
	kfxm := r.URL.Query().Get("kfxm")

	db := dao.GetConn()
	defer db.Close()
	//fmt.Println("select * from t_user where u_name=? and u_password=? limit 1")
	rows, err := db.Query("select distinct(kfjg) from serviceItem where itemName=? ",kfxm)
	checkErr(err)
	var Arr []module.ServiceItem
	for rows.Next() {
		//var ItemId int
		//var ItemName string
		//var Pinyin string
		var Kfjg string


		err = rows.Scan(&Kfjg)
		checkErr(err)

		var __col module.ServiceItem
		//__col.ItemId = ItemId
		//__col.ItemName = ItemName
		//__col.Pinyin = Pinyin
		//__col.U_Password = U_Password
		__col.Kfjg = Kfjg


		Arr = append(Arr, __col)

	}

    if len(Arr)>0{
	data, err := json.Marshal(Arr)
	if err != nil {
		log.Fatal("err get data: ", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Server", "NanjingYouzi")
	defer func() {
		io.WriteString(w, string(data))
		//fmt.Print(string(data))
		//w.Write(data)
	}()
	}}