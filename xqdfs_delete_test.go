package xqdfs_client_go

import (
	"testing"
	"fmt"
)

func TestClientDelete(t *testing.T) {
	url:="MSwxLDIsNjU5ODAwNzc3MTc3MTkzMjE0NCwxMTIyMzM0NA=="
	err:=ClientDelete("http://192.168.10.25:10087/opt/delete",url)
	if err!=nil {
		t.Error(err)
	}else{
		fmt.Println("delete ok")
	}
}

