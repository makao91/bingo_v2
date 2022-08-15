package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func (app *Config) createPlayerButton() fyne.CanvasObject {

	button := widget.NewButton("Join the Game!", app.createNewPlayerDialog())

	app.CreatePlayerButton = button

	return button
}

func (app *Config) createNewPlayerDialog() func() {
	return func() {
		insertedPlayerName := widget.NewEntry()

		// create a dialog
		addForm := dialog.NewForm(
			"New Player",
			"Create",
			"Cancel",
			[]*widget.FormItem{
				{Text: "Name", Widget: insertedPlayerName},
			},
			func(valid bool) {
				if valid {
					playerBoard := app.getPlayerBoard(insertedPlayerName.Text)
					app.AllPlayerContainer.Add(playerBoard)
					app.setAdaptiveGridLayout()
					app.AllPlayerContainer.Refresh()
				}
			},
			app.MainWindow)

		// size and show the dialog
		addForm.Resize(fyne.Size{Width: 400})
		addForm.Show()
	}
}

func (app *Config) setAdaptiveGridLayout() {
	if len(app.AllPlayerContainer.Objects) < 4 {
		app.AllPlayerContainer.Layout = container.NewAdaptiveGrid(len(app.AllPlayerContainer.Objects)).Layout
	}
}
