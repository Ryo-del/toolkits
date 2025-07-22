package portscanner

import (
	"fmt"
	"net"
	"sync"
	"time"
)

const (
	targetIP   = "127.0.0.1"
	numWorkers = 4
	timeout    = 500 * time.Millisecond
	protocols  = "tcp"
)

var portsToScan = []int{22, 80, 443, 8080, 8181}

// worker проверяет порты и вызывает коллбек при открытии
func worker(ports <-chan int, wg *sync.WaitGroup, onOpen func(int)) {
	defer wg.Done()

	for port := range ports {
		address := fmt.Sprintf("[%s]:%d", targetIP, port)
		conn, err := net.DialTimeout(protocols, address, timeout)
		if err == nil {
			onOpen(port) // вызываем переданный коллбек
			conn.Close()
		}
	}
}

// ScanPorts запускает сканирование с заданной функцией onOpen
func ScanPorts(onOpen func(int)) {
	ports := make(chan int, len(portsToScan))
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(ports, &wg, onOpen)
	}

	for _, port := range portsToScan {
		ports <- port
	}

	close(ports)
	wg.Wait()
}
