/**
* 1. 下载 hosts 文件的方法
* 2. 替换本机 hosts 文件的方法
* 3. 运行本地浏览器的方法
 */
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	// TODO 这里需要纠结一下 Golang 多线程的问题.
	startServer()
	Browser()
	//filename := GetHosts()
	//UpdateLocalHosts(filename)
}

func startServer() {
	http.HandleFunc("/", homePage)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		os.Exit(2)
	}
}

func homePage(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<html>Hello</html>")
}

// 判断是什么环境
func IsWindows() bool {
	return os.DevNull == "NULL"
}

// 找到执行浏览器的命令
func BrowserCmd() string {
	if IsWindows() {
		return "start"
	} else {
		return "open"
	}
}

// 打开浏览器
func Browser() {
	cmd := exec.Command(BrowserCmd(), "http://0.0.0.0:8081")
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("open url")
}

// 下载 hosts 文件
func GetHosts() string {
	url := "https://raw.github.com/wppurking/wyatt_hosts/master/hosts"
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return "获取最新 hosts 文件失败"
	}
	body, _ := ioutil.ReadAll(resp.Body)
	err = ioutil.WriteFile("./hosts", body, 0644)
	if err != nil {
		return "写入数据失败"
	}
	return "./hosts"
}

func LocalHostsPath() string {
	if IsWindows() {
		return "c:/windows/system32/drivers/etc/hosts"
	} else {
		return "/etc/hosts"
	}
}

// 更新本地的 hosts 文件
func UpdateLocalHosts(filename string) {
	bytes, _ := ioutil.ReadFile(filename)
	bak, _ := ioutil.ReadFile(LocalHostsPath())
	err := ioutil.WriteFile(LocalHostsPath(), bytes, 0644)
	if err != nil {
		_ = ioutil.WriteFile(LocalHostsPath(), bak, 0644)
	}
	fmt.Println("hosts 替换成功")
	AfterUpdate()
}

func AfterUpdate() {
	if IsWindows() {
		cmd := exec.Command("ipconfig /flushdns")
		cmd.Run()
		fmt.Println(cmd.Output())
	}
}
