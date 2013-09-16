package host

import (
 "os/exec"
)

// 找到 hosts 文件路径
func LocalHostsPath() string {
	return "c:/windows/system32/drivers/etc/hosts"
}

// 找到执行浏览器的命令
func BrowserCmd() string {
	return "explorer"
}

// 最后执行更新
func AfterUpdate() []byte {
	cmd := exec.Command("ipconfig", "/flushdns")
	output, err := cmd.Output()
	if err != nil {
		return []byte(err.Error())
	} else {
		return output
	}
}
