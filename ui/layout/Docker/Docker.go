package Docker

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	docker "github.com/ryo-del/devops-toolkit/internal/Docker"
)

func NewDockerTab() fyne.CanvasObject {
	title := widget.NewLabel("Docker Management")
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	options := []string{"Python", "Go", "Java", "Node.js"}
	LangLabel := widget.NewLabel("Выберите язык программирования:")

	dynamicBox := container.NewVBox()

	// лейбл для вывода результата (пустой изначально)
	resultLabel := widget.NewLabel("")

	selectLang := widget.NewSelect(options, func(selected string) {
		docker.ProgramLang = selected
		dynamicBox.Objects = nil

		switch selected {
		case "Python":
			PythonImages := []string{
				"python:3.14", "python:3.14-slim", "python:3.14-alpine",
				"python:3.13", "python:3.13-slim", "python:3.13-alpine",
			}
			imageText := widget.NewLabel("Выберите образ для контейнера:")
			PythonSelectImages := widget.NewSelect(PythonImages, func(sel string) {
				docker.Python_Image = sel
			})

			workDirLabel := widget.NewLabel("Укажите рабочую директорию (например, /app):")
			workDirEntry := widget.NewEntry()
			workDirEntry.SetPlaceHolder("/app")
			workDirEntry.OnChanged = func(text string) { docker.WorkDir = text }

			fileLabel := widget.NewLabel("Укажите файл зависимостей:")
			fileEntry := widget.NewEntry()
			fileEntry.SetPlaceHolder("requirements.txt")
			fileEntry.OnChanged = func(text string) { docker.File = text }

			launchLabel := widget.NewLabel("Введите команду запуска:")
			launchEntry := widget.NewEntry()
			launchEntry.SetPlaceHolder(`["python", "main.py"]`)
			launchEntry.OnChanged = func(text string) { docker.Launch = text }

			dynamicBox.Add(imageText)
			dynamicBox.Add(PythonSelectImages)
			dynamicBox.Add(workDirLabel)
			dynamicBox.Add(workDirEntry)
			dynamicBox.Add(fileLabel)
			dynamicBox.Add(fileEntry)
			dynamicBox.Add(launchLabel)
			dynamicBox.Add(launchEntry)

		case "Go":
			GoImages := []string{
				"golang:1.22", "golang:1.21", "golang:1.20-alpine",
			}
			imageText := widget.NewLabel("Выберите образ для контейнера:")
			GoSelectImages := widget.NewSelect(GoImages, func(sel string) {
				docker.Go_Image = sel
			})

			workDirLabel := widget.NewLabel("Укажите рабочую директорию (например, /app):")
			workDirEntry := widget.NewEntry()
			workDirEntry.SetPlaceHolder("/app")
			workDirEntry.OnChanged = func(text string) { docker.WorkDir = text }

			mainLabel := widget.NewLabel("Укажите главный файл:")
			mainEntry := widget.NewEntry()
			mainEntry.SetPlaceHolder("main.go")
			mainEntry.OnChanged = func(text string) { docker.File = text }

			launchLabel := widget.NewLabel("Введите команду запуска:")
			launchEntry := widget.NewEntry()
			launchEntry.SetPlaceHolder(`["go", "run", "main.go"]`)
			launchEntry.OnChanged = func(text string) { docker.Launch = text }

			dynamicBox.Add(imageText)
			dynamicBox.Add(GoSelectImages)
			dynamicBox.Add(workDirLabel)
			dynamicBox.Add(workDirEntry)
			dynamicBox.Add(mainLabel)
			dynamicBox.Add(mainEntry)
			dynamicBox.Add(launchLabel)
			dynamicBox.Add(launchEntry)

		case "Java":
			JavaImages := []string{
				"openjdk:21", "openjdk:17", "openjdk:11",
			}
			imageText := widget.NewLabel("Выберите образ для контейнера:")
			JavaSelectImages := widget.NewSelect(JavaImages, func(sel string) {
				docker.Java_Image = sel
			})

			workDirLabel := widget.NewLabel("Укажите рабочую директорию (например, /app):")
			workDirEntry := widget.NewEntry()
			workDirEntry.SetPlaceHolder("/app")
			workDirEntry.OnChanged = func(text string) { docker.WorkDir = text }

			jarLabel := widget.NewLabel("Укажите JAR файл:")
			jarEntry := widget.NewEntry()
			jarEntry.SetPlaceHolder("app.jar")
			jarEntry.OnChanged = func(text string) { docker.File = text }

			launchLabel := widget.NewLabel("Введите команду запуска:")
			launchEntry := widget.NewEntry()
			launchEntry.SetPlaceHolder(`["java", "-jar", "app.jar"]`)
			launchEntry.OnChanged = func(text string) { docker.Launch = text }

			dynamicBox.Add(imageText)
			dynamicBox.Add(JavaSelectImages)
			dynamicBox.Add(workDirLabel)
			dynamicBox.Add(workDirEntry)
			dynamicBox.Add(jarLabel)
			dynamicBox.Add(jarEntry)
			dynamicBox.Add(launchLabel)
			dynamicBox.Add(launchEntry)

		case "Node.js":
			NodeImages := []string{
				"node:22", "node:20", "node:18-alpine",
			}
			imageText := widget.NewLabel("Выберите образ для контейнера:")
			NodeSelectImages := widget.NewSelect(NodeImages, func(sel string) {
				docker.Node_Image = sel
			})

			workDirLabel := widget.NewLabel("Укажите рабочую директорию (например, /app):")
			workDirEntry := widget.NewEntry()
			workDirEntry.SetPlaceHolder("/app")
			workDirEntry.OnChanged = func(text string) { docker.WorkDir = text }

			fileLabel := widget.NewLabel("Укажите главный JS файл:")
			fileEntry := widget.NewEntry()
			fileEntry.SetPlaceHolder("server.js")
			fileEntry.OnChanged = func(text string) { docker.File = text }

			launchLabel := widget.NewLabel("Введите команду запуска:")
			launchEntry := widget.NewEntry()
			launchEntry.SetPlaceHolder(`["node", "server.js"]`)
			launchEntry.OnChanged = func(text string) { docker.Launch = text }

			dynamicBox.Add(imageText)
			dynamicBox.Add(NodeSelectImages)
			dynamicBox.Add(workDirLabel)
			dynamicBox.Add(workDirEntry)
			dynamicBox.Add(fileLabel)
			dynamicBox.Add(fileEntry)
			dynamicBox.Add(launchLabel)
			dynamicBox.Add(launchEntry)
		}

		// кнопка генерации Dockerfile
		generateButton := widget.NewButton("Сгенерировать Dockerfile", func() {
			result := docker.GenerateDockerfile()
			resultLabel.SetText(result) // обновляем текст вместо ошибки
			resultLabel.Refresh()

			// показываем во всплывающем окне
			win := fyne.CurrentApp().NewWindow("Dockerfile")
			win.SetContent(widget.NewLabel(result))
			win.Resize(fyne.NewSize(400, 300))
			win.Show()
		})
		copyButton := widget.NewButton("📋 Скопировать Dockerfile", func() {
			clip := fyne.CurrentApp().Driver().AllWindows()[0].Clipboard()
			clip.SetContent(docker.GenerateDockerfile())
		})

		dynamicBox.Add(generateButton)
		dynamicBox.Add(resultLabel)
		dynamicBox.Add(copyButton)
		dynamicBox.Refresh()
	})

	selectBox := container.New(layout.NewGridWrapLayout(fyne.NewSize(200, 40)), selectLang)

	return container.NewVBox(
		title,
		LangLabel,
		selectBox,
		dynamicBox,
	)
}
