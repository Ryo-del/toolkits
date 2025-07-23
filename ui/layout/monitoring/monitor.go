package monitor

import (
	"time"

	"github.com/ryo-del/devops-toolkit/internal/monitor"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewMonitorTab() fyne.CanvasObject {
	cpulabel := widget.NewLabel("Loading CPU...")

	// Горутина для обновления текста каждую 1 секунды
	go func() {
		for {
			usage, err := monitor.GetCPUUsage()
			if err != nil {
				usage = "Error getting CPU"
			}
			cpulabel.SetText(usage) // обновляем лейбл
			time.Sleep(1 * time.Second)
		}
	}()
	memlabel := widget.NewLabel("Loading MEM...")

	// Горутина для обновления текста каждую 1 секунды
	go func() {
		for {
			usage, err := monitor.GetMemoryUsage()
			if err != nil {
				usage = "Error getting MEM"
			}
			memlabel.SetText(usage) // обновляем лейбл
			time.Sleep(1 * time.Second)
		}
	}()

	return container.NewVBox(
		widget.NewLabel("Системный мониторинг"),
		cpulabel,
		memlabel,
	)
}
