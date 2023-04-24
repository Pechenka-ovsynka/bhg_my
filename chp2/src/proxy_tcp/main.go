package main

import(
  "log"
  "io"
  "net"
)

func handle(src net.Conn) {
    dst, err := net.Dial("tcp","google.com:80")
    if err != nil {
        log.Fatalln("Unable to connect to target",err)
    }
    defer dst.Close()

    go func() {
        _, err := io.Copy(dst, src)
        if err != nil {
            log.Fatalln(err)
        }
    }()
    _, err = io.Copy(src,dst)
    if err !=  nil {
        log.Fatalln(err)
    }
}

func main() {
    listener, err := net.Listen("tcp",":20081")
    if err != nil {
        log.Fatalln("Unuble to bind port", err)
    }
    for {
        conn, err :=listener.Accept()
        if err != nil {
            log.Fatalln("Unuble to listener accept", err)
        }
        go handle(conn)
    }
}
