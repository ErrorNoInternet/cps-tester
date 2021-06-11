package main

import (
	"image/color"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var cpsCounter float64

func main() {
	app := app.New()
	window := app.NewWindow("CPS")

	cpsLabel := canvas.NewText("0.0 CPS", color.White)
	cpsLabel.TextSize = 18
	cpsLabel.Move(fyne.NewPos(4, 0))
	clickButton := widget.NewButton("Click", buttonClick)
	clickButton.Move(fyne.NewPos(4, 34))
	clickButton.Resize(fyne.Size{Width: 94, Height: 28})

	window.Resize(fyne.Size{Width: 110, Height: 72})
	window.SetFixedSize(true)
	window.SetContent(container.NewWithoutLayout(cpsLabel, clickButton))
	go manageCPS(cpsLabel)
	window.ShowAndRun()
}

func addClicks() {
	for i := 0; i < 10; i++ {
		cpsCounter += 0.1
		time.Sleep(100 * time.Millisecond)
	}
}

func buttonClick() {
	go addClicks()
}

func manageCPS(label *canvas.Text) {
	for {
		time.Sleep(time.Second)
		label.Text = strconv.FormatFloat(cpsCounter, 'f', 1, 64) + " CPS"
		label.Refresh()
		cpsCounter = 0
	}
}
