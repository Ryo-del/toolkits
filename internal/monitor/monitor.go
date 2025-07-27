package monitor

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

type Config struct {
	Monitor MonitorConfig `json:"monitor"`
}
type MonitorConfig struct {
	CPU      bool `json:"cpu"`      //config.CPU
	Interval int  `json:"interval"` // в секундах, например
}

var config MonitorConfig

func init() {
	//json
	file, err := os.Open("internal/settings/main.json")
	if err != nil {
		panic(err) // если не нашли конфиг — останавливаем программу
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		panic(err) // если не удалось прочитать JSON
	}

	config = cfg.Monitor
}

func GetCPUUsage() (string, error) {
	percent, err := cpu.Percent(time.Duration(config.Interval)*time.Second, config.CPU)
	if err != nil || len(percent) == 0 {
		return "", err
	}
	return fmt.Sprintf("CPU Usage: %.2f%%", percent[0]), nil
}

func GetMemoryUsage() (string, error) {
	m, err := mem.VirtualMemory()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Memory Usage: %.2f%%", m.UsedPercent), nil
}

func GetDiskUsage() (string, error) {
	diskUse, err := disk.Usage("/")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Disk Usage: %.2f%%", diskUse.UsedPercent), nil
}

func GetNetworkIO() (string, error) {
	netIO, err := net.IOCounters(false)
	if err != nil || len(netIO) == 0 {
		return "", err
	}

	// Переводим байты в мегабайты
	mbSent := float64(netIO[0].BytesSent) / (1024 * 1024)
	mbRecv := float64(netIO[0].BytesRecv) / (1024 * 1024)

	return fmt.Sprintf("Network IO: %.2f MB sent, %.2f MB received", mbSent, mbRecv), nil
}

func GetHostInfo() (string, error) {
	hostInfo, err := host.Info()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Host Info: %s, OS: %s, Platform: %s", hostInfo.Hostname, hostInfo.OS, hostInfo.Platform), nil
}
