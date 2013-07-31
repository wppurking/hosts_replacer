package host

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
  "os/exec"
)


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

func AfterUpdate() []byte {
  if IsWindows() {
    cmd := exec.Command("ipconfig /flushdns")
    cmd.Run()
    output, _ := cmd.Output()
    return output
  } else {
    return []byte("hosts update successful.")
  }
}
