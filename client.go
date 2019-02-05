package main

import (
  "flag"
  "bufio"
  "net"
  "fmt"
  "os"
  "strconv"
  "regexp"
  "time"
)

var host = flag.String("host", "localhost", "Host to connect to; defaults to \"localhost\".")
var port = flag.Int("port", 8080, "The port to connect to; defaults to 8080.")

func readConnection(conn net.Conn) {
  for {
    scanner := bufio.NewScanner(conn)
    
    for {
      ok := scanner.Scan()
      text := scanner.Text()
      
      command := handleCommands(text)
      if !command {
        fmt.Printf("\b\b** %s\n> ", text)
      }
      
      if !ok {
        fmt.Println("EOF Server connection.")
        break
      }
    }
  }
}

func handleCommands(text string) bool {
  r, err := regexp.Compile("^%.*%$")
  if err == nil {
    if r.MatchString(text) {
      switch {
      case text == "%quit%":
        fmt.Println("\b\b[Server] disconnecting..")
        os.Exit(0)
      }
      return true
    }
  }
  return false
}

func main() {
  flag.Parse()
  
  dest := *host + ":" + strconv.Itoa(*port)
  fmt.Printf("[Client] Connecting to %s...\n", dest)
  
  conn, err := net.Dial("tcp", dest)
  
  if err != nil {
    if _, t := err.(*net.OpError); t{
      fmt.Println("Error connecting...")
    } else {
      fmt.Println("Unknown error: " + err.Error())
    }
    os.Exit(1)
  }
  
  go readConnection(conn)

  for {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("> ")
    text, _ := reader.ReadString('\n')
    
    conn.setWriteDeadline(time.Now().Add(1 * time.Second))
    _, err := conn.Write([]byte(text))
    if err != nil {
      fmt.Println("Error writing to stream.")
      break
    }
  }
}
