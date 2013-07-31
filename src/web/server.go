package web

import (
  "io"
  "net/http"
  "os"
  "host"
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

func homePage(w http.ResponseWriter, req *http.Request) {
  io.WriteString(w, "<html>" + closeLink() + "<br><a href='http://0.0.0.0:8081/update'>更新 Hosts 文件</a></html>")
}

func closeLink() string {
  return "<a href='http://0.0.0.0:8081/exit'>退出</a>"
}

func updateHosts(w http.ResponseWriter, req *http.Request) {
  filename := host.GetHosts()
  host.UpdateLocalHosts(filename)

  hosts, _ := host.Hosts()
  cmdOutput := host.AfterUpdate()

  io.WriteString(w, "<html>" + closeLink() + "<h3>运行结果:" + string(cmdOutput) + "</h3><br><hr><br>" + strings.Join(strings.Split(string(hosts), "\n"), "<br>") + "</html>")
}

func stopServer(w http.ResponseWriter, req *http.Request) {
  os.Remove("./hosts")
  exitChannel <- true
}