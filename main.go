package main

import (
  "fmt"
  "net"
  "flag"
  "bufio"
  "os"
)

var msg string
var addr string

func handleClient(c net.Conn) {

    writer := bufio.NewWriter(c)
    reader := bufio.NewReader(c)

    scanner := bufio.NewReader(os.Stdin)

    for {

      msg, _ := scanner.ReadString('\n')

      writer.WriteString(msg)
      writer.Flush()

      resp, err := reader.ReadString('\n')

      if err != nil {
        return
      }

      fmt.Print(resp)
    }
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
}
