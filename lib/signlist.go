package lib

import (
	"net/http"
	"encoding/json"
	"log"
	"io"

	"help/dao"
	"help/module"

	"fmt"
)

//帮扶列表
func Signlist(w http.ResponseWriter, r *http.Request) {
	workerId := r.URL.Query().Get("workerId")
	pageIndex := r.URL.Query().Get("pageIndex")
	db := dao.GetConn()
	defer db.Close()
	//_, err := strconv.Atoi(pageIndex)
	fmt.Println("select * from sign  where workerId='"+workerId+"'  limit "+pageIndex+",20")
	rows, err := db.Query("select * from sign  where workerId='"+workerId+"' order by workTime desc  limit "+pageIndex+",20")
	checkErr(err)
	var Arr []module.Sign
	for rows.Next() {
		var Id         int
		var WorkerId   string
		var WorkerName string
		var Latitude   float64
		var Longitude  float64
		var Location   string
		var WorkTime   string

		err = rows.Scan(&Id, &WorkerId, &WorkerName, &Latitude, &Longitude, &Location, &WorkTime)
		checkErr(err)

		var __col module.Sign
		__col.Id = Id
		__col.WorkerId = WorkerId
		__col.WorkerName = WorkerName
		__col.Latitude = Latitude
		__col.Longitude = Longitude
		__col.Location = Location
		__col.WorkTime = WorkTime

		Arr = append(Arr, __col)

	}

	//
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
}