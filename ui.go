package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"log"
)

type PlayersBoard struct {
	board      []*ModeratorBoardButton
	playerName string
}

type Config struct {
	App                fyne.App
	MainWindow         fyne.Window
	InfoLog            *log.Logger
	ErrorLog           *log.Logger
	ModeratorFields    [][]interface{}
	AllPlayersBoards   []PlayersBoard
	AllPlayerContainer *fyne.Container
	FinalPlayerContent *fyne.Container
	CreatePlayerButton fyne.CanvasObject
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
