package settings

import (
	"encoding/json"
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Config struct {
	Scaner  ScanerConfig  `json:"scaner"`
	Parser  ParserConfig  `json:"parser"`
	Monitor MonitorConfig `json:"monitor"`
}

type ScanerConfig struct {
	Protocol string `json:"protocol"`
	Ports    []int  `json:"ports"`
	IP       string `json:"ip"`
	Work     int    `json:"Work"`
	Time     int    `json:"time"`
}

type ParserConfig struct {
	LogPath string `json:"log_path"`
	Format  string `json:"format"`
	SaveTo  string `json:"save_to"`
}

type MonitorConfig struct {
	CPU      bool `json:"cpu"`
	Memory   bool `json:"memory"`
	Interval int  `json:"interval"` // в секундах, например
}

func NewSettingsTab() fyne.CanvasObject {
	// Читаем JSON
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

	defaultProtocol := cfg.Scaner.Protocol

	infolabel := widget.NewLabel("Settings\n")

	radio := widget.NewRadioGroup([]string{"tcp", "udp"}, func(value string) {
		// Обновляем поле в структуре
		cfg.Scaner.Protocol = value

		// Пытаемся сохранить обратно в файл
		file, err := os.OpenFile("internal/settings/main.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Println("Ошибка открытия для записи:", err)
			return
		}
		defer file.Close()

		// Пишем красиво отформатированный JSON
		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")
		err = encoder.Encode(cfg)
		if err != nil {
			fmt.Println("Ошибка при записи JSON:", err)
		} else {
			fmt.Println("Протокол сохранён:", value)
		}
	})

	radio.Horizontal = true
	radio.SetSelected(defaultProtocol)

	return container.NewVBox(
		infolabel,
		radio,
	)
}
