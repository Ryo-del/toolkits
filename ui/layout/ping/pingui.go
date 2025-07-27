package ping

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/ryo-del/devops-toolkit/internal/ping"
)

func NewpingTab() fyne.CanvasObject {
	interval := ping.GetInterval() // получаем интервал обновления

	hostLabel := widget.NewLabel("Loading host...")
	statusLabel := widget.NewLabel("Loading status...")

	// Обновление статуса пинга
	go func() {
		for {
			ping.RunPing()
			status := ping.GetPingStatus() // статус: ✅ Alive / ❌ Unreachable
			statusLabel.SetText(status)    // обновляем UI
			time.Sleep(time.Duration(interval) * time.Second)
		}
	}()

	// Обновление хоста
	go func() {
		for {
			ping.RunPing()
			host := ping.GetSource() // получаем (google.com)
			hostLabel.SetText(host)
			time.Sleep(time.Duration(interval) * time.Second)
		}
	}()

	return container.NewVBox(
		hostLabel,
		statusLabel,
	)
}
