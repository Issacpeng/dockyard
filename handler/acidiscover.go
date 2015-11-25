package handler

import (
	//"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/macaron.v1"
)

func DiscoveryACIHandler(ctx *macaron.Context) (int, []byte) {
	repo := ctx.Params(":repository")
	fmt.Printf("\r\nGetRktDiscoveryHandler ctx.Req.URL: %v, repo:%v\r\n", ctx.Req.URL, repo)

	// TBD: generate aci template and endpoint like ac-push
	/*
		t, err := template.ParseFiles(path.Join(templatedir, "acitemplate.html"))
		if err != nil {
			fmt.Fprintf(ctx.Resp, fmt.Sprintf("%v", err))
			result, _ := json.Marshal(err)
			return http.StatusInternalServerError, result
		}
		fmt.Printf("renderListOfACIs t: %v:\r\n", t)
		acis, err := listACIs()
		if err != nil {
			result, _ := json.Marshal(err)
			return http.StatusInternalServerError, result
		}
		fmt.Printf("renderListOfACIs acis: %v:\r\n", acis)
		err = t.Execute(ctx.Resp, struct {
			ServerName string
			ACIs       []aci
			HTTPS      bool
		}{
			ServerName: serverName,
			ACIs:       acis,
			HTTPS:      *https,
		})
		if err != nil {
			result, _ := json.Marshal(err)
			return http.StatusInternalServerError, result
		}

		result, _ := json.Marshal(map[string]string{})
		return http.StatusOK, result
	*/

	//return the body of resp data from coreos.com, including template prepared for http requests
	//    result := "<html lang=\"en\">\r\n"

	//    result += "<head>\r\n"
	//    result += "<meta charset=\"utf-8\">\r\n"
	result := "<meta name=\"ac-discovery\" content="
	result += "\"containerops.me/etcd https://containerops.me/ac-pull/{version}/etcd-{version}-{os}-{arch}.{ext}\">\r\n"
	result += "<meta name=\"ac-discovery-pubkeys\" content="
	result += "\"containerops.me/etcd https://coreos.com/dist/pubkeys/aci-pubkeys.gpg\">\r\n"
	result += "</head>\r\n"

	//    result += "<body class=\"coreos-home co-m-main-nav-transparent co-p-header-large\">\r\n"
	//    result += "  </body>\r\n"

	//    result += "</html>\r\n"

	fmt.Printf("result %v\r\n", result)
	return http.StatusOK, []byte(result)

}
