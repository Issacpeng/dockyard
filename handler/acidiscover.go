package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path"

	"github.com/astaxie/beego/logs"
	"gopkg.in/macaron.v1"

	"github.com/containerops/wrench/setting"
)

// TBD: discovery template should be updated to keep in line with ACI
func DiscoveryACIHandler(ctx *macaron.Context, log *logs.BeeLogger) {
	img := ctx.Params(":imagename")
    fmt.Println("############## renderListOfACIs ##############\r\n")
	os.RemoveAll(path.Join(directory, "tmp"))
	err := os.MkdirAll(path.Join(directory, "tmp"), 0755)
	if err != nil {
        fmt.Println("############## MkdirAll fail ##############\r\n")
		fmt.Fprintf(os.Stderr, "%v", err)
	}

    fmt.Println("############## ParseFiles ##############\r\n")
	t, err := template.ParseFiles("conf/acitemplate.html")
	if err != nil {
		log.Error("[ACI API] Discovery parse template file failed: %v", err.Error())
		ctx.Resp.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(ctx.Resp, fmt.Sprintf("%v", err))
		return
	}

	acis, err := listACIs()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}

	err = t.Execute(ctx.Resp, struct {
		Name       string
		ServerName string
		ListenMode string
		ACIs       []aci
	}{
		Name:       img,
		ServerName: setting.Domains,
		ListenMode: setting.ListenMode,
		ACIs:       acis,
	})
	if err != nil {
		log.Error("[ACI API] Discovery ACIlist failed: %v", err.Error())
		fmt.Fprintf(ctx.Resp, fmt.Sprintf("%v", err))
	}
}
