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
func PeopleDetail(w http.ResponseWriter, r *http.Request) {
	idCardNo := r.URL.Query().Get("idCardNo")

	var sql string
	sql = "select Id, IdCard, CardNo, CardStateId, DisableType, DisableModeId, DisableLevel, IsChild,Name, Gender, Birth, Nation, Phone, Address, Guardian,GuardPhone, Relation, Coordinator, PickTime,Avatar,Szxq,Szxz,Szc,Kfxm,Kfjg from people where IdCard='"+idCardNo+"' "

	db := dao.GetConn()
	defer db.Close()
	fmt.Println(sql)
	var Id int
	var IdCard string
	var CardNo string
	var CardStateId int
	var DisableType string
	var DisableModeId string
	var DisableLevel string
	var IsChild int
	var Name string
	var Gender int
	var Birth string
	var Nation string
	var Phone string
	var Address string
	var Guardian string
	var GuardPhone string
	var Relation string
	var Coordinator string
	var PickTime string
	var Avatar string
	var Szxq string
	var Szxz string
	var Szc string
	var Kfxm string
	var Kfjg string
	err := db.QueryRow(sql).Scan(&Id, &IdCard, &CardNo, &CardStateId, &DisableType, &DisableModeId, &DisableLevel, &IsChild, &Name, &Gender, &Birth, &Nation, &Phone, &Address, &Guardian, &GuardPhone, &Relation, &Coordinator, &PickTime, &Avatar, &Szxq, &Szxz, &Szc, &Kfxm, &Kfjg)
	if err!=nil{
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Server", "NanjingYouzi")
		defer func() {
			io.WriteString(w, "error")
			//fmt.Print(string(data))
			//w.Write(data)
		}()
	}else {
			var __col module.People
			__col.Id = Id
			__col.IdCard = IdCard
			__col.CardNo = CardNo
			__col.CardStateId = CardStateId
			__col.DisableType = DisableType
			__col.DisableModeId = DisableModeId
			__col.DisableLevel = DisableLevel
			__col.IsChild = IsChild
			__col.Name = Name

			__col.Gender = Gender
			__col.Birth = Birth
			__col.Nation = Nation
			__col.Phone = Phone
			__col.Address = Address
			__col.Guardian = Guardian
			__col.GuardPhone = GuardPhone
			__col.Relation = Relation
			__col.Coordinator = Coordinator
			__col.PickTime = PickTime
			__col.Avatar = Avatar
			__col.Szxq = Szxq
			__col.Szxz = Szxz
			__col.Szc = Szc
			__col.Kfxm = Kfxm
			__col.Kfjg = Kfjg
		//
		data, err := json.Marshal(__col)
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
}