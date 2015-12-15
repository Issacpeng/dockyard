<<<<<<< HEAD
package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/astaxie/beego/logs"
	"github.com/containerops/dockyard/models"
	"gopkg.in/macaron.v1"
)

func GetRktPukkeysHandler(ctx *macaron.Context, log *logs.BeeLogger) (int, []byte) {

	fmt.Println("GetRktPukkeysHandler")

	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}

func GetRktfileHandler(ctx *macaron.Context, log *logs.BeeLogger) (int, []byte) {

	fmt.Println("GetRktSignfileHandler")
	version := ctx.Params(":version")
	name := ctx.Params(":name")

	fmt.Printf("GetRktfileHandler version:%v\r\n", version)
	fmt.Printf("GetRktfileHandler name:%v\r\n", name)

	asc := strings.Contains(name, ".asc")
	if asc == true {
		status, result := GetRktSignfileHandler(name, ctx, log)
		return status, result
	} else {
		status, result := GetRktImagefileHandler(name, ctx, log)
		return status, result
	}
	result, _ := json.Marshal(map[string]string{})
	return http.StatusOK, result
}

func GetRktSignfileHandler(name string, ctx *macaron.Context, log *logs.BeeLogger) (int, []byte) {

	var jsonInfo string
	var payload string
	var err error

	fmt.Println("GetRktSignfileHandler")

	digest := "sha256:a3ed95caeb02ffe68cdd9fd84406680ae93d633cb16422d00e8a7c22955b46d4"
	tarsum := strings.Split(digest, ":")[1]

	i := new(models.Image)
	has, _ := i.HasTarsum(tarsum)
	if has == false {
		log.Error("[RKT PULL] Digest not found: %v", tarsum)

		result, _ := json.Marshal(map[string]string{"message": "Digest not found"})
		return http.StatusNotFound, result
	}

	if jsonInfo, err = i.GetJSON(tarsum); err != nil {
		log.Error("[RKT PULL] Search Image SIGN Error: %v", err.Error())

		result, _ := json.Marshal(map[string]string{"message": "Search Image SIGN Error"})
		return http.StatusNotFound, result
	}

	ctx.Resp.Header().Set("Rkt-Checksum-Payload", fmt.Sprintf("sha256:%v", payload))
	ctx.Resp.Header().Set("Rkt-Size", fmt.Sprint(i.Size))
	ctx.Resp.Header().Set("Content-Length", fmt.Sprint(len(jsonInfo)))

	return http.StatusOK, []byte(jsonInfo)

}

func GetRktImagefileHandler(name string, ctx *macaron.Context, log *logs.BeeLogger) (int, []byte) {

	fmt.Println("GetRktImagefileHandler")

	digest := "sha256:a3ed95caeb02ffe68cdd9fd84406680ae93d633cb16422d00e8a7c22955b46d4"
	tarsum := strings.Split(digest, ":")[1]

	i := new(models.Image)
	has, _ := i.HasTarsum(tarsum)
	if has == false {
		log.Error("[RKT PULL] Digest not found: %v", tarsum)

		result, _ := json.Marshal(map[string]string{"message": "Digest not found"})
		return http.StatusNotFound, result
	}

	layerfile := i.Path
	if _, err := os.Stat(layerfile); err != nil {
		log.Error("[RKT PULL] File path is invalid: %v", err.Error())

		result, _ := json.Marshal(map[string]string{"message": "File path is invalid"})
		return http.StatusBadRequest, result
	}

	file, err := ioutil.ReadFile(layerfile)
	if err != nil {
		log.Error("[RKT PULL] Read file failed: %v", err.Error())

		result, _ := json.Marshal(map[string]string{"message": "Read file failed"})
		return http.StatusBadRequest, result
	}

	ctx.Resp.Header().Set("Content-Type", "application/x-gzip")
	ctx.Resp.Header().Set("Rkt-Content-Digest", name)
	ctx.Resp.Header().Set("Content-Length", fmt.Sprint(len(file)))

	return http.StatusOK, file

}
=======
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
>>>>>>> acpush-opt
