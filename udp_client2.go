package main

import (
   "fmt"
   "net"
//   "bytes"
   "github.com/digitnow/eksamen22/mycrypt"
)

func main() { 
    /*
    serverAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8888")
	if err != nil {
		return
	}
    */
    
    client, err := net.ListenPacket("udp", "127.0.0.1:8900")
    if err != nil {
        fmt.Println(err)
    }
    defer func() { _ = client.Close() }()

    buf := make([]byte, 1024)
    n, _, err := client.ReadFrom(buf)
    if err != nil {
        fmt.Println(err)
    }

    ALF_OPPG8 := []rune("abcdefghijklmnopqrstuvwxyzæøå, ")
    kryptert_streng := string(buf[:n])
    dekryptert_tilbake := string(mycrypt.Krypter([]rune(kryptert_streng), ALF_OPPG8, len(ALF_OPPG8)-4))
    fmt.Printf("Dekryptert melding: %s\n", dekryptert_tilbake)

}
