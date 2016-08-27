package main

import (
  "fmt"
  "json"
  "net"
  "flag"
  "bufio"
)

var addr := flag.String("addr", "127.1:9999", "Conenction address (addr:port, addr or :port)")

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

          resp := reader.ReadString('\n')
          fmt.Println(resp)
        }
  }()
}

func main() {

    flag.Parse()

    conn, err := net.Dial("tcp", addr)
    defer conn.Close()

    if err != nil {
      fmt.Println("Could not connect to: " + addr)
      os.Exit(1)
    }

    createClient(conn)
}
