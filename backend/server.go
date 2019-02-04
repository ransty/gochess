package main

import (
  "fmt"
  "net"
)

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
  _, err := conn.WriteToUDP([]byte("[Server] : message received "), addr)
  if err != nil {
    fmt.Printf("Couldn't send response %v", err)
  }
}

func main() {
  p := make([]byte, 2048)
  addr := net.UDPAddr {
    Port: 666,
    IP: net.ParseIP("127.0.0.1"),
  }
  ser, err := net.ListenUDP("udp", &addr)
  if err != nil {
    fmt.Printf("Something went wrong %v\n", err)
    return
  }
  for {
    _, remoteaddr, err := ser.ReadFromUDP(p)
    fmt.Printf("Read a message from %v %s \n", remoteaddr, p)
    if err != nil {
      fmt.Printf("Something went wrong %v", err)
      continue
    }
    go sendResponse(ser, remoteaddr)
}
    
