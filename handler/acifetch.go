package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"fmt"
	"strings"

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
    aciId := utils.MD5(fmt.Sprintf("%s", acname))
    a := new(models.Aci)

    if asc := strings.Contains(acname, ".aci.asc"); asc == true {
        if imgpath, err = a.GetSignPath(aciId); err != nil {

		    log.Error("[ACI API] Get ACI SignPath failed: %v", err.Error())
		    result, _ := json.Marshal(map[string]string{"message": "Get ACI SignPath failed"})
		    return http.StatusNotFound, result
		}
    } else {
        if imgpath, err = a.GetAciPath(aciId); err != nil {

			log.Error("[ACI API] Get ACI AciPath failed: %v", err.Error())
			result, _ := json.Marshal(map[string]string{"message": "Get ACI AciPath failed"})
			return http.StatusNotFound, result
		}
    }

	if img, err = ioutil.ReadFile(imgpath); err != nil {
		// TBD: consider to fetch image from other storage medium

		log.Error("[ACI API] Get ACI file failed: %v", err.Error())
		result, _ := json.Marshal(map[string]string{"message": "Get ACI file failed"})
		return http.StatusNotFound, result
	}

	return http.StatusOK, img

}
