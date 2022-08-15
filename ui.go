package main

import (
	"fyne.io/fyne/v2/container"
)

type PlayersBoard struct {
	board      []*ModeratorBoardButton
	playerName string
}

func (app *Config) makeUI() {

	//make player container
	app.AllPlayerContainer = container.NewGridWithColumns(1)

	//create new player button
	createPlayerButton := app.createPlayerButton()

	//get moderator board
	moderator_board := app.getModeratorBoard()

	// add container to window
	finalContent := container.NewBorder(
		createPlayerButton,
		moderator_board,
		nil,
		nil,
		app.AllPlayerContainer,
	)
	app.FinalPlayerContent = finalContent
	app.MainWindow.SetContent(finalContent)
}
