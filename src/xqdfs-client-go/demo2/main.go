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

	for i:=int64(0);i<256;i++{
		go func(){
			for {
				lock.Lock()
				count++
				lock.Unlock()
				
				url,err:=xqdfs_client_go.ClientUpload("http://127.0.0.1:10087/opt/upload",img)
				if err!=nil{
					fmt.Println(err)
					continue
				}
				
				data,err:=xqdfs_client_go.ClientGet("http://127.0.0.1:10087/opt/get",url)
				if err!=nil{
					fmt.Println(err)
					continue
				}

				if (count%1000)==0{
					fmt.Println(count," ",len(data))
				}
			}
		}()
	}

	for{
		time.Sleep(time.Second)
	}
}
