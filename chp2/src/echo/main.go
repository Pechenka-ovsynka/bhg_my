package main

import (
    "log"
    "net"
    "bufio"
)

func echo(conn net.Conn){
    // отложенное закрытие соединения после выполнения функции
    defer conn.Close()
    reader := bufio.NewReader(conn)
    s, err := reader.ReadString('\n')
    if err != nil {
        log.Fatalln("Unable to read data", err)
    }
    log.Printf("Read %d bytes from %s, data: %s", len(s), conn.RemoteAddr(), s)

    log.Println("Echo data")
    writer := bufio.NewWriter(conn)
    _, err = writer.WriteString(s)
    if err != nil {
        log.Fatalln("Unable to write data", err)
    }
    writer.Flush()
}

func main() {
    port := "20080"
    addres := "127.0.0.1"
    addres = addres + ":" + string(port)
    listener, err := net.Listen("tcp", addres)

    if err != nil {
        log.Fatalln("Unable to bind port", port, "error:", err)
    }
    log.Printf("Listening on 0.0.0.0:%s", port)
    for {
        conn, err := listener.Accept()
        log.Println("new connection")
        if err != nil {
            log.Fatalln("Unable to acept connection", err)
        }
        go echo(conn)

    }

}
