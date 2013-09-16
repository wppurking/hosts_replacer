package host

// 找到 hosts 文件路径
func LocalHostsPath() string {
	return "/etc/hosts"
}

// 找到执行浏览器的命令
func BrowserCmd() string {
	return "open"
}

// 最后执行更新
func AfterUpdate() []byte {
	return []byte("hosts update successful.")
}
