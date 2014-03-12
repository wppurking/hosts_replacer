#hosts_replacer
使用 Golang 写的 hosts 文件替换应用

支持 windows XP 和 Mac OS.

编译使用下面参数, 减少将近 50% 体积
go build -ldflags '-w -s'

下面的解释都是 Google 来的
-w : 去掉DWARF调试信息，得到的程序就不能用gdb调试了
-s : 去掉符号表（然后panic时候的stack trace就没有任何文件名/行号信息了. 这个等价于普通C/C++程序被strip的效果）
