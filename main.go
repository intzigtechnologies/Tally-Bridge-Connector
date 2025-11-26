package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"io"
	"net/http"
)

func main() {
	a := app.New()
	w := a.NewWindow("Tally Bridge Connector")

	label_ip := widget.NewLabel("Your ip address is: ")
	label_ip_display := widget.NewLabel("")

	btn := widget.NewButton("Disply ip", func() {
		resp, err := http.Get("https://api.ipify.org?format=text")
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(body)
		label_ip_display.SetText(string(body))
	})

	w.SetContent(
		container.NewVBox(
			label_ip,
			label_ip_display,
			btn,
		),
	)

	w.ShowAndRun()
}
