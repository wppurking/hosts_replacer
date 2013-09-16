package web

import (
	"host"
	"io"
	"net/http"
	"os"
	"strings"
)

var exitChannel chan bool

func StartServer(ch chan bool) {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/update", updateHosts)
	http.HandleFunc("/exit", stopServer)

	exitChannel = ch

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		os.Exit(2)
	}
}

func header() string {
	return "<head><meta http-equiv='content-type' content='text/html;charset=utf-8'></head>"
}

func homePage(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<html>" + header() + "<body>" + closeLink() + 
		"<br><a href='http://localhost:8081/update'>更新 Hosts 文件</a></body></html>")
}

func closeLink() string {
	return "<a href='http://localhost:8081/exit'>退出</a>"
}

func updateHosts(w http.ResponseWriter, req *http.Request) {
	filename := host.GetHosts()
	host.UpdateLocalHosts(filename)

	hosts, _ := host.Hosts()
	cmdOutput := host.AfterUpdate()

	io.WriteString(w, "<html>" + header() + "<body>" + closeLink() + 
		"<h3>运行结果:" + string(cmdOutput) + "</h3><br><hr><br>" + 
		strings.Join(strings.Split(string(hosts), "\n"), "<br>") + "</body></html>")
}

func stopServer(w http.ResponseWriter, req *http.Request) {
	os.Remove("./hosts")
	exitChannel <- true
}
