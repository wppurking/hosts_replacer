package host

import (
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

// 打开浏览器
func Browser() {
	cmd := exec.Command(BrowserCmd(), "http://localhost:8081")
	err := cmd.Start()
	if err != nil {
		log.Println(err)
	}

	log.Println("Open url:", "http://localhost:8081", "in browser to start.")
}

// 下载 hosts 文件
func GetHosts() string {
	url := "https://raw.github.com/wppurking/wyatt_hosts/master/hosts"
	resp, err := http.Get(url)
	if err != nil {
		return "获取最新 hosts 文件失败"
	}
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	err = ioutil.WriteFile("./hosts", body, 0644)
	if err != nil {
		return "写入数据失败"
	}
	return "./hosts"
}

func Hosts() ([]byte, error) {
	return ioutil.ReadFile(LocalHostsPath())
}

// 更新本地的 hosts 文件
func UpdateLocalHosts(filename string) {
	bytes, _ := ioutil.ReadFile(filename)
	bak, _ := Hosts()
	err := ioutil.WriteFile(LocalHostsPath(), bytes, 0644)
	if err != nil {
		_ = ioutil.WriteFile(LocalHostsPath(), bak, 0644)
	}
}
