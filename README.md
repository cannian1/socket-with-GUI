# About socket with GUI

基于 walk框架的 go语言 和python 在局域网下socket通信

***

**说明**

​        半成品，学习 golang 中的小玩具

**使用**

参考https://github.com/lxn/walk

将 server.exe.manifest 放在 与 server.go 同级的目录下，执行 go build 。

需要保证电脑此时连接WiFi，即在 cmd 中运行 ipconfig，可以看到 “无线局域网适配器 WLAN:”字段。



python 客户端代码内需要将 IP 改为上述字段下的 IP 地址。



先启动服务端，再启动客户端，

客户端中发送的数据，在点击“编辑参数1”后会回显到绑定的只读文本框中。
