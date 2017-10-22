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
func Record(w http.ResponseWriter, r *http.Request) {
	pageIndex := r.URL.Query().Get("pageIndex")
	db := dao.GetConn()
	defer db.Close()
	fmt.Println("select * from collection where idcardNo="+r.URL.Query().Get("idcardNo")+" limit "+pageIndex+",20")
	rows, err := db.Query("select Id, DispeopleId, IdcardNo, DiscardNo, ServiceId, AgencyId, AccessRecord,Photos, Latitude,Longitude,Location,AccessTime,LivePhoto,RecordPhoto,Name,Gender,Phone,Level,Birth,PickerId,Avatar,DisableModeId,pickerName from collection where del=0 and idcardNo="+r.URL.Query().Get("idcardNo")+" order by accessTime desc  limit "+pageIndex+",20")
	checkErr(err)
	var Arr []module.Collection
	for rows.Next() {
		var Id int
		var DispeopleId string
		var IdcardNo string
		var DiscardNo string
		var ServiceId string
		var AgencyId string
		var AccessRecord string
		var Photos string
		var Latitude string
		var Longitude string
		var Location string
		var AccessTime string
		var LivePhoto string
		var RecordPhoto string
		var Name string
		var Gender int
		var Phone string
		var Level string
		var Birth string
		var PickerId string
		var Avatar string
		var DisableModeId string
		var PickerName string

		err = rows.Scan(&Id, &DispeopleId, &IdcardNo, &DiscardNo, &ServiceId, &AgencyId, &AccessRecord,&Photos, &Latitude,&Longitude,&Location,&AccessTime,&LivePhoto,&RecordPhoto,&Name,&Gender,&Phone,&Level,&Birth,&PickerId,&Avatar,&DisableModeId,&PickerName)
		checkErr(err)


		var __col module.Collection
		__col.Id = Id
		__col.DispeopleId = DispeopleId
		__col.IdcardNo = IdcardNo
		__col.DiscardNo = DiscardNo
		__col.ServiceId = ServiceId
		__col.AgencyId = AgencyId
		__col.AccessRecord = AccessRecord
		__col.Latitude = Latitude
		__col.Longitude = Longitude
		__col.Location = Location
		__col.AccessTime = AccessTime
		__col.LivePhoto = LivePhoto
		__col.RecordPhoto = RecordPhoto
		__col.Name = Name
		__col.Gender = Gender
		__col.Phone = Phone
		__col.Level = Level
		__col.Birth = Birth
		__col.Photos = Photos
		__col.PickerId = PickerId
		__col.Avatar = Avatar
		__col.DisableModeId = DisableModeId
		__col.PickerName = PickerName

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