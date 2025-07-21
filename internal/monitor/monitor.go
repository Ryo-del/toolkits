package monitor

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

func GetCPUUsage() (string, error) {
	percent, err := cpu.Percent(time.Second, false)
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
	return fmt.Sprintf("Network IO: %d bytes sent, %d bytes received", netIO[0].BytesSent, netIO[0].BytesRecv), nil
}

func GetHostInfo() (string, error) {
	hostInfo, err := host.Info()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Host Info: %s, OS: %s, Platform: %s", hostInfo.Hostname, hostInfo.OS, hostInfo.Platform), nil
}
