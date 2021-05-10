#!/usr/bin/python3
# -*-coding:utf-8 -*-
from socket import *
HOST = '192.168.1.120' #服务端ip
PORT = 12345 #服务端端口号
BUFSIZ = 1024 #缓冲区大小
ADDR = (HOST, PORT)
tcpCliSock = socket(AF_INET, SOCK_STREAM) #创建socket对象
tcpCliSock.connect(ADDR) #连接服务器
while True:
    data = input('>>').strip()
    if not data:
        break
    tcpCliSock.send(data.encode('utf-8')) #发送消息
    data = tcpCliSock.recv(BUFSIZ) #读取消息
    if not data:
        break
    print(data.decode('utf-8'))
tcpCliSock.close() #关闭客户端