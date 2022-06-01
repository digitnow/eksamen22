package main

import (
   "fmt"
   "net"
   "bytes"
   "github.com/digitnow/eksamen22/mycrypt"
)

func main() { 
    serverAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8888")
	if err != nil {
		return
	}
    
    client, err := net.ListenPacket("udp", "127.0.0.1:")
    if err != nil {
        fmt.Println(err)
    }
    defer func() { _ = client.Close() }()

    ALF_OPPG8 := []rune("abcdefghijklmnopqrstuvwxyzæøå, ")
    kryptert_streng := string(mycrypt.Krypter([]rune("w, x og y møtes i ålesund"), ALF_OPPG8, 4))
    fmt.Printf("Kryptert melding: %s\n", kryptert_streng)

    msg := []byte(kryptert_streng)
    _, err = client.WriteTo(msg, serverAddr)
    if err != nil {
        fmt.Println(err)
    }

    buf := make([]byte, 1024)
    n, addr, err := client.ReadFrom(buf)
    if err != nil {
        fmt.Println(err)
    }

    if addr.String() != serverAddr.String() {
        fmt.Printf("received reply from %q instead of %q", addr, serverAddr)
    }

    if !bytes.Equal(msg, buf[:n]) {
        fmt.Printf("expected reply %q; actual reply %q", msg, buf[:n])
    }

}
