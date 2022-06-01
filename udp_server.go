package main

import (
    "net"
    "fmt"
    "context"
)

func main() {
    client2Addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8900")
    if err != nil {
        return
    }

    _, cancel := context.WithCancel(context.Background())
    s, err := net.ListenPacket("udp", "127.0.0.1:8888")
    if err != nil {
        fmt.Errorf("binding to udp %s: %w", "127.0.0.1:8888", err)
    }
    defer cancel()
    //go func() {
        //go func() {
          //  <-ctx.Done()
           // _ = s.Close()
       // }()
        buf := make([]byte, 1024)
        for {
            n, clientAddr, err := s.ReadFrom(buf) // client to server
            fmt.Println("after s.ReadFrom") 
            if err != nil {
                fmt.Println("error s.ReadFrom")
                return
            }
            _, err = s.WriteTo(buf[:n], clientAddr) // server to client
            if err != nil {
                return
            } 
            _, err = s.WriteTo(buf[:n], client2Addr) // server to client2
            if err != nil {
                return
            } 
        }
        //cancel()
    //}()
}
