package portscanner

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

type Config struct {
	Scaner ScanerConfig `json:"scaner"`
}

type ScanerConfig struct {
	Protocol string `json:"protocol"`
	Ports    []int  `json:"ports"`
	IP       string `json:"ip"`
	Work     int    `json:"Work"`
	Time     int    `json:"time"`
}

// worker проверяет порты и вызывает коллбек при открытии
func worker(cfg ScanerConfig, ports <-chan int, wg *sync.WaitGroup, onOpen func(int)) {
	defer wg.Done()

	timeout := time.Duration(cfg.Time) * time.Millisecond // переводим в time.Duration

	for port := range ports {
		// формируем адрес в нужном формате
		address := fmt.Sprintf("%s:%d", cfg.IP, port)

		// DialTimeout: протокол, адрес, таймаут
		conn, err := net.DialTimeout(cfg.Protocol, address, timeout*time.Second)
		if err == nil {
			onOpen(port) // вызываем коллбек, если порт открыт
			conn.Close()
		}
	}
}

// ScanPorts запускает сканирование с заданной функцией onOpen
func ScanPorts(onOpen func(int)) {
	// Читаем конфиг из файла
	file, err := os.Open("internal/settings/main.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		panic(err)
	}

	ports := make(chan int, len(cfg.Scaner.Ports))
	var wg sync.WaitGroup

	// Запускаем воркеры, передаём cfg.Scaner в каждый
	for i := 0; i < cfg.Scaner.Work; i++ {
		wg.Add(1)
		go worker(cfg.Scaner, ports, &wg, onOpen)
	}

	// Отправляем порты в канал
	for _, port := range cfg.Scaner.Ports {
		ports <- port
	}
	close(ports)

	wg.Wait()
}
