package main

import (
	"fmt"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/ryo-del/devops-toolkit/internal/portscanner"
)

func main() {
	a := app.New()
	myWindow := a.NewWindow("Port Scanner")
	myWindow.Resize(fyne.NewSize(400, 300))

	infolabel := widget.NewLabel("Scanning ports on 127.0.0.1")
	output := widget.NewLabel("Results will appear here")

	btn := widget.NewButton("Сканировать", func() {
		output.SetText("") // очищаем
		var sb strings.Builder

		go func() {
			portscanner.ScanPorts(func(port int) {
				sb.WriteString(fmt.Sprintf("[OPEN] Port %d\n", port))
				// обновление UI должно быть в главном потоке
				time.Sleep(10 * time.Millisecond) // чтобы не моргал
				a.SendNotification(&fyne.Notification{Title: "Порт найден", Content: fmt.Sprintf("Port %d открыт", port)})
				output.SetText(sb.String())
			})
		}()
	})

	myWindow.SetContent(container.NewVBox(
		infolabel,
		btn,
		output,
	))
	myWindow.ShowAndRun()
}
