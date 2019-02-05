package main

import (
  "bufio"
  "flag"
  "net"
  "fmt"
  "os"
  "time"
  "strconv"
)

var addr = flag.String("addr", "", "Address to listen to; default is \"\" (all interfaces).")
var port = flag.Int("port", 8080, "The port to listen on; default is 8080.")

func handleConnection(conn net.Conn) {
  remoteAddr := conn.RemoteAddr().String()
  fmt.Println("[Client] connected from " + remoteAddr)
  
  scanner := bufio.NewScanner(conn)
  
  for {
    ok := scanner.Scan()
    
    if !ok {
      break
    }
    
    handleMessage(scanner.Text(), conn)
  }
  
  fmt.Println("[Client] @ " + remoteAddr + " disconnected.")
}

func handleMessage(message string, conn net.Conn) {
  fmt.Println("> " + message)
  
  if len(message) > 0 && message[0] == '::' {
    switch {
    case message == "::time":
      resp := "It is " + time.Now().String() + "\n"
      fmt.Print("< " + resp)
      conn.Write([]byte(resp))
    case message == "::quit":
      fmt.Println("Quitting.")
      conn.Write([]byte("[Server] Shutting down....\n"))
      fmt.Println("< " + "%quit%")
      os.Exit(0)
    default:
      conn.Write([]byte("Unrecognized command.\n"))
    }
  } 
}
