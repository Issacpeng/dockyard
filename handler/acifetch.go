package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"fmt"
	"strings"
//	"time"

	"github.com/astaxie/beego/logs"
	"gopkg.in/macaron.v1"

//	"github.com/containerops/wrench/setting"
	"github.com/containerops/dockyard/models"
	"github.com/containerops/wrench/utils"
)

/* TBD:
current implementation as blow just be added for testing ACI fetch,
they would be updated after ACI ac-push finished
*/

func GetPubkeysHandler(ctx *macaron.Context, log *logs.BeeLogger) (int, []byte) {
	var pubkey []byte
	var err error

	servername := ctx.Params(":servername")

    p := new(models.Pubkey)
    if pubkey, err = p.GetPubkey(servername); err != nil {

		log.Error("[ACI API] Get pubkey file failed: %v", err.Error())
		result, _ := json.Marshal(map[string]string{"message": "Get pubkey file failed"})
		return http.StatusNotFound, result
	}

	return http.StatusOK, pubkey
}

func GetACIHandler(ctx *macaron.Context, log *logs.BeeLogger) (int, []byte) {
	var img []byte
	var imgpath string
	var err error

	acname := ctx.Params(":acname")
    imagename := strings.Trim(acname, ".asc") //Trim ".asc" of aciID to find currect aci image signature and file record
    aciId := utils.MD5(fmt.Sprintf("%s", imagename))
    fmt.Printf("#########  imagename:%v  #########\r\n", imagename) 
    fmt.Printf("#########  aciId:%v  #########\r\n", aciId) 

    a := new(models.Aci)
/*
    a.Imagename= "etcd-v2.2.2-linux-amd64.aci"
    a.AciId = utils.MD5(fmt.Sprintf("%s", a.Imagename))
    a.SignPath= setting.ImagePath + "/" + "etcd-v2.2.2-linux-amd64.aci.asc"
    a.AciPath = setting.ImagePath + "/" + "etcd-v2.2.2-linux-amd64.aci"
    a.CreatedTime = time.Now().UnixNano() / int64(time.Millisecond)
    a.Save()
*/
    if _, _, err := a.Has(aciId); err != nil {

     	log.Error("[ACI API] Get ACI file failed: %v", err.Error())
		result, _ := json.Marshal(map[string]string{"message": "Searching ACI file failed"})
		return http.StatusNotFound, result
    }  	
    if asc := strings.Contains(acname, ".aci.asc"); asc == true {
        imgpath = a.SignPath
    } else {
        imgpath = a.AciPath
    }

    fmt.Printf("#########  imgpath:%v  #########\r\n", imgpath)

	if img, err = ioutil.ReadFile(imgpath); err != nil {
		// TBD: consider to fetch image from other storage medium

		log.Error("[ACI API] Get ACI file failed: %v", err.Error())
		result, _ := json.Marshal(map[string]string{"message": "Get ACI file failed"})
		return http.StatusNotFound, result
	}

	return http.StatusOK, img

}
