package portscanner

import (
	"fmt"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/ryo-del/devops-toolkit/internal/portscanner"
)

func NewParserTab() fyne.CanvasObject {
	infolabel := widget.NewLabel("Scanning ports on 127.0.0.1")
	output := widget.NewLabel("Results will appear here")

	btn := widget.NewButton("Сканировать", func() {
		output.SetText("") // очищаем
		var sb strings.Builder
		var portsFound bool

		go func() {
			portscanner.ScanPorts(func(port int) {
				portsFound = true
				sb.WriteString(fmt.Sprintf("[OPEN] Port %d\n", port))
				// обновление UI должно быть в главном потоке
				time.Sleep(10 * time.Millisecond) // чтобы не моргал
				output.SetText(sb.String())
			})

			// Если ни одного порта не найдено
			if !portsFound {
				output.SetText("Все порты закрыты")
			}
		}()
	})

	return container.NewVBox(
		infolabel,
		btn,
		output,
	)

}
