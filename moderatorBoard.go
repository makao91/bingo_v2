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
	{
		{fieldName: "PM", activated: false},
		{fieldName: "meble", activated: false},
		{fieldName: "życie", activated: false},
	},
	{
		{fieldName: "dupy z Tindera", activated: false},
		{fieldName: "temperatura", activated: false},
		{fieldName: "internet w firmie", activated: false},
	},
	{
		{fieldName: "pieniądze", activated: false},
		{fieldName: "polityka", activated: false},
		{fieldName: "juniorzy", activated: false},
	},
	{
		{fieldName: "devops", activated: false},
		{fieldName: "gnój w projekcie", activated: false},
		{fieldName: "estymacje", activated: false},
	},
	{
		{fieldName: "dracha na muszli", activated: false},
		{fieldName: "wyjebane puszki z kosza", activated: false},
		{fieldName: "zmielona kawa", activated: false},
	},
	{
		{fieldName: "smród", activated: false},
		{fieldName: "monitoring", activated: false},
		{fieldName: "asd", activated: false},
	},
	{
		{fieldName: "anime", activated: false},
		{fieldName: "żeglarz", activated: false},
		{fieldName: "opluwanie pracowników", activated: false},
	},
}

type ModeratorBoardButton struct {
	widget.Button
	fieldName string
	activated bool
}

func (app *Config) getModeratorBoard() *fyne.Container {
	moderatorGrid := app.getModeratorGrid()

	moderatorBoardContainer := container.NewBorder(
		nil,
		nil,
		nil,
		nil,
		container.NewAdaptiveGrid(1, moderatorGrid),
	)

	return moderatorBoardContainer
}

func (app *Config) getModeratorGrid() *fyne.Container {

	grid := container.NewGridWithColumns(3)
	for r := 0; r < 10; r++ {
		for c := 0; c < 3; c++ {
			grid.Add(app.newModeratorField(r, c))
		}
	}

	return grid
}
func (app *Config) newModeratorField(row, column int) *ModeratorBoardButton {
	i := &ModeratorBoardButton{fieldName: fullModeratorBoard[row][column].fieldName, activated: fullModeratorBoard[row][column].activated}
	i.SetText(fullModeratorBoard[row][column].fieldName)
	i.OnTapped = app.launchModeratorButton(i)
	i.Importance = widget.LowImportance
	i.ExtendBaseWidget(i)

	return i
}

func (app *Config) launchModeratorButton(i *ModeratorBoardButton) func() {
	return func() {
		i.activated = true
		for _, board := range app.AllPlayersBoards {
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
