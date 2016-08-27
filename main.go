package main

import (
  "fmt"
  "net"
  "flag"
  "bufio"
  "os"
)

var addr string

func createClient(c net.Conn) {

    var msg string

    writer := bufio.NewWriter(c)
    reader := bufio.NewReader(c)

    go func(){
        for {
          fmt.Println(">> ")
          fmt.Scanln(&msg)

          writer.WriteString(msg)
          writer.Flush()

          resp, _ := reader.ReadString('\n')
          fmt.Println(resp)
        }
  }()
}

func main() {

    flag.StringVar(&addr, "addr", "127.1:9999", "Conenction address (addr:port, addr or :port)")
    flag.Parse()

    conn, err := net.Dial("tcp", addr)

    if err != nil {
      fmt.Println("Could not connect to: " + addr)
      os.Exit(1)
    }

    createClient(conn)
}
