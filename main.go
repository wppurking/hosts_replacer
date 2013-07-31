/**
* 1. 下载 hosts 文件的方法
* 2. 替换本机 hosts 文件的方法
* 3. 运行本地浏览器的方法
 */
package main

import (
	"host"
	"web"
)

var ch = make(chan bool)

func main() {
	go web.StartServer(ch)
	host.Browser()
	<-ch
}
