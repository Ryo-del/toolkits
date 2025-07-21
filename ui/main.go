package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()                      // Create a new application instance
	myWindow := myApp.NewWindow("toolkits") // Create a new window with the title "Hello"
	// Set the size of the window
	myWindow.Resize(fyne.NewSize(600, 400))

	myWindow.Show()
	myApp.Run()
	tidyUp()
}

func tidyUp() { // Perform any necessary cleanup before exiting
	fmt.Println("Exited")
}
