package main

import (
	"fmt"
	"net"
	"time"
	"sync"
	"os"
)

func worker(ports chan int, address string, wg *sync.WaitGroup){
		for p := range ports {
				scan_port(p, address)
				wg.Done()
		}
}

func scan_port(port int, address string){
		target := fmt.Sprintf("%s:%d",address, port)
		conn, err := net.DialTimeout("tcp", target,  1 * time.Second)
		if err != nil {
				return
		}
		fmt.Printf("Port %d: open\n", port)
		conn.Close()
}

func main() {
	var wg sync.WaitGroup
	ports := make(chan int, 1000)
	if len(os.Args) <= 1{
		fmt.Println("use test_scan target")
		return
	}
  address :=  os.Args[1]
	fmt.Println(address)
	for  i := 0; i < cap(ports); i++ {
		go worker(ports, address, &wg)
	}

	for i := 0; i <= 65535; i++  {
			wg.Add(1)
			ports <- i
	}
		wg.Wait()
}
