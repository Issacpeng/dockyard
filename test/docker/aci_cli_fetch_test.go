package main

import (
	"os/exec"
	"testing"

	"github.com/containerops/dockyard/test"
)


func TestAciFetch(t *testing.T) {
	var cmd *exec.Cmd
	var err error
	var out string

    tests := []struct{
    	domains string
		name    string
		version string
		os      string
		arch    string
		ext     string
    }{
      	{test.Domains, "etcd", "v2.2.2", "linux", "amd64", ".aci"},
    }

    for _, tt := range tests {
    	aciname := tt.domains + "/" + tt.name + ":" + tt.version

    	cmd = exec.Command(test.RktBinary, "fetch", aciname)
		if out, err = test.ParseCmdCtx(cmd); err != nil {
			t.Fatalf("fetch aci image %v failed: [Info]%v, [Error]%v", aciname, out, err)
		}

        imagename := tt.name + "-" + tt.version + "-" + tt.os + "-" + tt.arch + tt.ext
		acihttps := "https://" + tt.domains + "/" + "ac-image" + "/" + imagename

    	cmd = exec.Command(test.RktBinary, "fetch", acihttps)
		if out, err = test.ParseCmdCtx(cmd); err != nil {
			t.Fatalf("fetch aci image %v failed: [Info]%v, [Error]%v", acihttps, out, err)
		}
    }
}