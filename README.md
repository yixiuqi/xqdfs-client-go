 
### 1、Upload
函数:ClientUpload

参数:
* host string-服务器地址  
* img []byte-图片数据

返回:
* url string-图片url，后续获取图片用  
* err error

```bash
img:=make([]byte,1024)
url,err:=ClientUpload("http://192.168.10.25:10087/opt/upload",img)
if err!=nil {
	fmt.Println(err)
}else{
	fmt.Println(url)
}
```

***

### 2、Get
函数:ClientGet

参数:
* host string-服务器地址  
* url string-图片url

返回:
* img []byte-图片数据 
* err error

```bash
url:="MSwxLDIsNjU5ODAwNzc3MTc3MTkzMjE0NCwxMTIyMzM0NA=="
data,err:=ClientGet("http://192.168.10.25:10087/opt/get",url)
if err!=nil {
	t.Error(err)
}else{
	fmt.Println("get image len:",len(data))
}
```

***

### 3、Delete
函数:ClientDelete

参数:
* host string-服务器地址  
* url string-图片url

返回:
* err error

```bash
url:="MSwxLDIsNjU5ODAwNzc3MTc3MTkzMjE0NCwxMTIyMzM0NA=="
err:=ClientDelete("http://192.168.10.25:10087/opt/delete",url)
if err!=nil {
	t.Error(err)
}else{
	fmt.Println("delete ok")
}
```

