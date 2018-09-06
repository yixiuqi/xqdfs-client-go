package xqdfs_client_go

import (
	"testing"
	"fmt"
)

func TestClientGet(t *testing.T) {
	url:="MSwxLDIsNjU5ODAwNzc3MTc3MTkzMjE0NCwxMTIyMzM0NA=="
	data,err:=ClientGet("http://192.168.10.25:10087/opt/get",url)
	if err!=nil {
		t.Error(err)
	}else{
		fmt.Println("get image len:",len(data))
	}
}

