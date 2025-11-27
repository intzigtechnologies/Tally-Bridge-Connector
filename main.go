package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"net"
)

func main() {
	a := app.New()
	w := a.NewWindow("Tally Bridge Connector")

	indicator_localhost := widget.NewLabel("")

	btn_check_localhost := widget.NewButton("Check Tally Server", func() {
		connection, err := net.Dial("tcp", "localhost:9000")
		if err != nil {
			indicator_localhost.SetText("Please turn on Tally Server: " + err.Error())
			return
		}
		defer connection.Close()

		indicator_localhost.SetText("Tally Server is running")
	})

	w.SetContent(
		container.NewVBox(
			indicator_localhost,
			btn_check_localhost,
		),
	)

	w.ShowAndRun()
}
