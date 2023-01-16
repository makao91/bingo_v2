package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"math/rand"
	"time"
)

var name string

func (app *Config) getPlayerBoard(playerName string) *fyne.Container {
	name = playerName
	playerGrid := app.getPlayerGrid()
	player := widget.NewLabel(playerName)
	player.Alignment = fyne.TextAlignCenter

	playerBoardContainer := container.NewBorder(
		player,
		widget.NewSeparator(),
		widget.NewSeparator(),
		widget.NewSeparator(),
		container.NewAdaptiveGrid(1, playerGrid),
	)

	return playerBoardContainer
}

func (app *Config) getPlayerGrid() *fyne.Container {
	newBoard := fullModeratorBoard
	app.shuffleBingo(newBoard)
	playersBoard := PlayersBoard{playerName: name}

	grid := container.NewGridWithColumns(3)
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			field := newPlayerField(r, c, newBoard)
			playersBoard.board = append(playersBoard.board, field)
			grid.Add(field)
		}
	}
	app.checkForBingo(playersBoard)
	app.AllPlayersBoards = append(app.AllPlayersBoards, playersBoard)

	return grid
}
func newPlayerField(row, column int, board [][]ModeratorBoardButton) *ModeratorBoardButton {
	i := &ModeratorBoardButton{fieldName: board[row][column].fieldName, activated: board[row][column].activated}
	i.SetText(board[row][column].fieldName)
	i.ExtendBaseWidget(i)

	return i
}

func (app *Config) shuffleBingo(board [][]ModeratorBoardButton) {
	rand.Seed(time.Now().UnixNano())
	for index := 0; index < 3; index++ {
		rand.Shuffle(len(board), func(i, j int) { board[i], board[j] = board[j], board[i] })
		for j := 0; j < 3; j++ {
			rand.Shuffle(len(board[index]), func(i, j int) {
				board[i][index], board[j][index] = board[j][index], board[i][index]
			})
		}
	}
}
