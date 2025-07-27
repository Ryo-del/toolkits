package monitor

import (
	"encoding/json"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/ryo-del/devops-toolkit/internal/monitor"
)

type Config struct {
	Monitor MonitorConfig `json:"monitor"`
}
type MonitorConfig struct {
	Interval int `json:"interval"` // интервал обновления в секундах
}

var config MonitorConfig

func NewMonitorTab() fyne.CanvasObject {
	// читаем конфиг
	file, err := os.Open("internal/settings/main.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		panic(err)
	}
	config = cfg.Monitor

	// создаём лейблы
	cpuLabel := widget.NewLabel("Loading CPU...")
	memLabel := widget.NewLabel("Loading Memory...")
	diskLabel := widget.NewLabel("Loading Disk...")
	netLabel := widget.NewLabel("Loading Network...")
	hostLabel := widget.NewLabel("Loading Host Info...")

	// отдельная функция для обновления лейбла
	updateLabel := func(label *widget.Label, getData func() (string, error)) {
		go func() {
			for {
				text, err := getData()
				if err != nil {
					label.SetText("Error")
				} else {
					label.SetText(text)
				}
				time.Sleep(time.Duration(config.Interval) * time.Second)
			}
		}()
	}

	// запускаем обновление всех лейблов
	updateLabel(cpuLabel, monitor.GetCPUUsage)
	updateLabel(memLabel, monitor.GetMemoryUsage)
	updateLabel(diskLabel, monitor.GetDiskUsage)
	updateLabel(netLabel, monitor.GetNetworkIO)
	updateLabel(hostLabel, monitor.GetHostInfo)

	// собираем UI
	return container.NewVBox(
		widget.NewLabel("Системный мониторинг"),
		cpuLabel,
		memLabel,
		diskLabel,
		netLabel,
		hostLabel,
	)
}
