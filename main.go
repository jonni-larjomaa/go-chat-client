package main

import (
  "fmt"
  "net"
  "flag"
  "bufio"
  "os"
  "strings"
  "encoding/json"
  "strconv"
)

type Message struct {
  Cmd     string
  Receiv  []int
  Msg     []byte
}

var msg string
var addr string

func handleClient(c net.Conn) {

    writer := bufio.NewWriter(c)
    reader := bufio.NewReader(c)

    scanner := bufio.NewReader(os.Stdin)

    go clientReader(reader)
    go clientWriter(scanner, writer)
}

func clientWriter(scanner *bufio.Reader, writer *bufio.Writer){
  for {

    msg, _ := scanner.ReadString('\n')

    jsonmsg := newMessage(msg);

    json, _ := json.Marshal(jsonmsg)

    if len(jsonmsg.Cmd) == 0 {
      continue;
    }

    writer.Write(append(json, byte('\n')))
    writer.Flush()

  }
}

func clientReader(reader *bufio.Reader){
  for {
    resp, err := reader.ReadString('\n')

    if err != nil {
      return
    }

    fmt.Print(resp)
  }
}

func newMessage(msg string) *Message {

  smsg := strings.SplitN(strings.TrimSpace(msg), " ", 3)

  switch {

    case "getid" == smsg[0], "list" == smsg[0]:
      return &Message{smsg[0], make([]int, 0), make([]byte,0)}
    case "send" == smsg[0]:
      return &Message{smsg[0], getReceiversList(strings.Split(smsg[1],",")), getMsg(smsg[2])}
    default:
      fmt.Println("Command not understood use: getid,list or send")
      return &Message{"", make([]int, 0), make([]byte,0)}

  }

}

func getMsg(msg string) []byte {

  b := make([]byte,0);
  maxlen := 1024;

  if len(msg) < maxlen {
    maxlen = len(msg)
  }

  for i :=0;i < maxlen; i++ {
    b = append(b, byte(msg[i]))
  }

  return b
}

func getReceiversList(strs []string) []int {
    list := make([]int, 0)

    for _, str := range strs {
      num, _ := strconv.ParseInt(str, 10, 0)
      list = append(list, int(num))
    }

    return list
}

func main() {

    flag.StringVar(&addr, "addr", "127.1:9999", "Conenction address (addr:port, addr or :port)")
    flag.Parse()

    conn, err := net.Dial("tcp", addr)

    if err != nil {
      fmt.Println("Could not connect to: " + addr)
      os.Exit(1)
    }

    handleClient(conn)

    for{}
}
