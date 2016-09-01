# GO-chat client for GO-chat server

usage:

0. server needs to be running -> check go-chat-server

1. go run main.go
  - defaults to 127.1:9999 when connecting
  - can be changed using -addr flag
2. write following command on prompt
  - getid, gets clients id
  - list, lists all clients on hub
  - send 1,2,3 <msg>, send message to clients separated by comma
