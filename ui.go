package main

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (app *Config) makeUI() {

	label := widget.NewLabel("sdfdsf")
	// get player board
	player_board := app.getModeratorBoard()

	//get moderator board
	moderator_board := app.getModeratorBoard()

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Prices", theme.HomeIcon(), player_board),
		container.NewTabItemWithIcon("Holdings", theme.InfoIcon(), moderator_board),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	// add container to window
	finalContent := container.NewVBox(label, tabs)

	app.MainWindow.SetContent(finalContent)
}
