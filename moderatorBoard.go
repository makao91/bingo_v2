package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var fullModeratorBoard = [][]ModeratorBoardButton{
	{
		{fieldName: "kod Witka", activated: false},
		{fieldName: "Teamsy", activated: false},
		{fieldName: "Unitree", activated: false},
	},
	{
		{fieldName: "pogoda", activated: false},
		{fieldName: "AgileBoard", activated: false},
		{fieldName: "TimeTracker", activated: false},
	},
	{
		{fieldName: "BitBucket", activated: false},
		{fieldName: "syf w kiblu", activated: false},
		{fieldName: "syf w kuchni", activated: false},
	},
}

type ModeratorBoardButton struct {
	widget.Button
	fieldName string
	activated bool
}

func (app *Config) getModeratorBoard() *fyne.Container {
	app.ModeratorGrid = app.getModeratorGrid()

	moderatorBoardContainer := container.NewBorder(
		nil,
		nil,
		nil,
		nil,
		container.NewAdaptiveGrid(1, app.ModeratorGrid),
	)

	return moderatorBoardContainer
}

func (app *Config) getModeratorGrid() *fyne.Container {

	grid := container.NewGridWithColumns(3)
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			grid.Add(newModeratorField(r, c))
		}
	}

	return grid
}
func newModeratorField(row, column int) *ModeratorBoardButton {
	i := &ModeratorBoardButton{fieldName: fullModeratorBoard[row][column].fieldName, activated: fullModeratorBoard[row][column].activated}
	i.SetText(fullModeratorBoard[row][column].fieldName)
	i.OnTapped = launchModeratorButton(i)
	i.Importance = widget.LowImportance
	i.ExtendBaseWidget(i)

	return i
}

func launchModeratorButton(i *ModeratorBoardButton) func() {
	return func() {
		i.activated = true
		i.Importance = widget.HighImportance
		for _, board := range AllPlayersBoards {
			for _, field := range board.board {
				if field.fieldName == i.fieldName {
					field.activated = true
					field.Importance = widget.HighImportance
					field.Refresh()
				}
			}

		}
	}
}
