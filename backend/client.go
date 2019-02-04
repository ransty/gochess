package main

import (
  "os"
  "fmt"
  "net"
  "bufio"
)

func main() {
  p := make([]byte, 2048)
  conn, err := net.Dial("udp", "127.0.0.1:666")
  var arg = "nil"
  if err != nil {
    fmt.Printf("Something went wrong %v\n", err)
    return
  }
  
  if len(os.Args) > 1 {
    arg = os.Args[1]
  } else {
    fmt.Printf("Missing a message argument, using default message\n")
  }
  
  fmt.Printf(conn, arg)
  _, err = bufio.NewReader(conn).Read(p)
  if err == nil {
    fmt.Printf("%s\n", p)
  } else {
    fmt.Printf("Something went wrong %v\n", err)
  }
  conn.Close()
}
