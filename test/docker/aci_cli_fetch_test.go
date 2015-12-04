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

    nametests := []struct{
    	domains string
		name    string
		version string
    }{
      	{test.Domains, "etcd", "v2.2.2"},
    }

    for _, tt := range nametests {
    	aciname := tt.domains + "/" + tt.name + ":" + tt.version

    	cmd = exec.Command(test.RktBinary, "fetch", aciname)
		if out, err = test.ParseCmdCtx(cmd); err != nil {
			t.Fatalf("fetch aci image %v failed: [Info]%v, [Error]%v", aciname, out, err)
		}
    }
}