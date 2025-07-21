package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()                      // Create a new application instance
	myWindow := myApp.NewWindow("toolkits") // Create a new window with the title "toolkits"
	myWindow.Resize(fyne.NewSize(800, 600)) // Set the window size to 800x600 pixel

	myWindow.Show()
	myApp.Run()
	tidyUp()
}

func tidyUp() {
	fmt.Println("Exited")
}
