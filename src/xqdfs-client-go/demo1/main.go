package main

import (
	"fmt"
	"sync"
	"time"
	"xqdfs-client-go"
)

func main() {
	img:=make([]byte,1024*64)
	count:=uint64(0)
	lock:=&sync.Mutex{}
	for i:=int64(0);i<64;i++{
		go func(){
			for {
				lock.Lock()
				count++
				lock.Unlock()

				url,err:=xqdfs_client_go.ClientUpload("http://192.168.10.25:10087/opt/upload",img)
				if err!=nil{
					fmt.Println(err)
					continue
				}
				err=xqdfs_client_go.ClientDelete("http://192.168.10.25:10087/opt/delete",url)
				if err!=nil{
					fmt.Println(err)
					continue
				}

				if (count%1000)==0{
					fmt.Println(count," ",url)
				}
			}
		}()
	}

	for{
		time.Sleep(time.Second)
	}
}
