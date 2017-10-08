package lib

import (
	"net/http"
	//"log"
	"io"
	"fmt"
	"os"
	"help/dao"
	"time"
	"path"
)

//更新采集
func Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println("client:", r.RemoteAddr, "method:", r.Method)
	//db := dao.GetConn()
	//defer db.Close()
	r.ParseForm()
	r.ParseMultipartForm(32 << 20) //最大内存为32M

	//读取参数
	id := r.PostFormValue("id")
	fmt.Println("id is",id)
	idCardNo := r.PostFormValue("idCardNo")
	fmt.Println("idCardNo is",idCardNo)
	serviceId := r.PostFormValue("serviceId")
	fmt.Println("serviceId is",serviceId)
	agencyId := r.PostFormValue("agencyId")
	fmt.Println("agencyId is",agencyId)
	accessRecord := r.PostFormValue("accessRecord")
	fmt.Println("accessRecord is",accessRecord)
	livePhoto := r.PostFormValue("livePhoto")
	fmt.Println("livePhoto is",livePhoto)
	recordPhoto := r.PostFormValue("recordPhoto")
	fmt.Println("recordPhoto is",recordPhoto)


	//log.Println("userId=", DispeopleId, "cityId=", AccessRecord)

	mp := r.MultipartForm
	if mp == nil {
		//log.Println("not MultipartForm.")
		//w.Write(([]byte)("不是MultipartForm格式"))
		/*w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Server", "NanjingYouzi")
		defer func() {
			io.WriteString(w,"notMultipartForm")
			//fmt.Print(string(data))
			//w.Write(data)
		}()
		return*/
	}

	fileHeaders, findFile := mp.File["file"]
	if !findFile || len(fileHeaders) == 0 {
		//log.Println("file count == 0.")
		//w.Write(([]byte)("没有上传文件"))
		/*w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Server", "NanjingYouzi")
		defer func() {
			io.WriteString(w,"noUploadfile")
			//fmt.Print(string(data))
			//w.Write(data)
		}()
		return*/
	}

	//s :=time.Now().UnixNano()
	var i int = 0
	for _, v := range fileHeaders {
		fileName := v.Filename
		file, err := v.Open()
		checkErr(err)
		defer file.Close()

		var filenameWithSuffix string
		filenameWithSuffix = path.Base(fileName) //获取文件名带后缀
		fmt.Println("filenameWithSuffix =", filenameWithSuffix)
		var fileSuffix string
		fileSuffix = path.Ext(filenameWithSuffix) //获取文件后缀
		//rnd:=rand.Intn(100)
		var p string
		if livePhoto!=""{
			temp := GetRandomString(15)+fileSuffix
			//p = "D:/usr/gopath/src/help/"+ temp
			p="/usr/img/"+temp
			livePhoto = temp
		}else if recordPhoto!=""{
			temp := GetRandomString(16)+fileSuffix
			//p = "D:/usr/gopath/src/help/"+ temp
			p="/usr/img/"+temp
			recordPhoto = temp
		}
		outputFilePath := p
		writer, err := os.OpenFile(outputFilePath, os.O_WRONLY|os.O_CREATE, 0666)
		checkErr(err)
		io.Copy(writer, file)
		i++
	}

	//insert mysql
	db := dao.GetConn()
	defer db.Close()
	//fmt.Println(_lat)
	//fmt.Println(r.PostFormValue("workerId"))
	if livePhoto!=""&&recordPhoto!=""{
	_, err := db.Exec(
		"UPDATE collection set serviceId=?,agencyId=?,accessRecord=?,livePhoto=?,recordPhoto=? where id= ?",
		serviceId,agencyId,accessRecord,livePhoto,recordPhoto,id)
	if err != nil {
		fmt.Println("update err1 is",err.Error())
		//log.Fatal("err post data2: ", err)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Server", "NanjingYouzi")
		defer func() {
			io.WriteString(w,"error")
			//fmt.Print(string(data))
			//w.Write(data)
		}()
		return
	}}else if livePhoto==""&&recordPhoto!=""{
		_, err := db.Exec(
			"UPDATE collection set serviceId=?,agencyId=?,accessRecord=?,recordPhoto=? where id= ?",
			serviceId,agencyId,accessRecord,recordPhoto,id)
		if err != nil {
			fmt.Println("update err2 is",err.Error())
			//log.Fatal("err post data2: ", err)
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Server", "NanjingYouzi")
			defer func() {
				io.WriteString(w,"error")
				//fmt.Print(string(data))
				//w.Write(data)
			}()
			return
		}
	}else if recordPhoto==""&&livePhoto!=""{
		_, err := db.Exec(
			"UPDATE collection set serviceId=?,agencyId=?,accessRecord=?,livePhoto=? where id= ?",
			serviceId,agencyId,accessRecord,livePhoto,id)
		if err != nil {
			fmt.Println("update err3 is",err.Error())
			//log.Fatal("err post data2: ", err)
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Server", "NanjingYouzi")
			defer func() {
				io.WriteString(w,"error")
				//fmt.Print(string(data))
				//w.Write(data)
			}()
			return
		}
	}else if recordPhoto==""&&livePhoto==""{
		_, err := db.Exec(
			"UPDATE collection set serviceId=?,agencyId=?,accessRecord=? where id= ?",
			serviceId,agencyId,accessRecord,id)
		if err != nil {
			fmt.Println("update err3 is",err.Error())
			//log.Fatal("err post data2: ", err)
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Server", "NanjingYouzi")
			defer func() {
				io.WriteString(w,"error")
				//fmt.Print(string(data))
				//w.Write(data)
			}()
			return
		}
	}



	//同时更新people表的采集时间
	fmt.Println("UPDATE  people set pickTime='"+time.Now().Format("2006-01-02 15:04:05")+"'where idCard="+idCardNo)
	_, err2 := db.Exec("UPDATE  people set pickTime='"+time.Now().Format("2006-01-02 15:04:05")+"'where idCard='"+idCardNo+"'")
	if err2!=nil{
		fmt.Println("update err4 is",err2.Error())
		//log.Fatal("err post data1: ", err)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Server", "NanjingYouzi")
		defer func() {
			io.WriteString(w,"error")
			//fmt.Print(string(data))
			//w.Write(data)
		}()
	}
	//data, err := json.Marshal(Arr)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Server", "NanjingYouzi")
	defer func() {
		io.WriteString(w,"ok")
		//fmt.Print(string(data))
		//w.Write(data)
	}()

	//msg := fmt.Sprintf("成功上传了%d个文件", len(fileHeaders))
	//w.Write(([]byte)(msg))

}