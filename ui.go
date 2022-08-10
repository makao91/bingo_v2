package main

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"time"
)

type PlayersBoard struct {
	board      []*ModeratorBoardButton
	playerName string
}

var AllPlayersBoards []PlayersBoard

func (app *Config) makeUI() {

	// get player board
	playerBoardOne := app.getPlayerBoard("Krzysiu")
	playerBoardTwo := app.getPlayerBoard("Maciek")
	playerBoardThree := app.getPlayerBoard("Długi")

	go func() {
		for range time.Tick(time.Second * 1) {
			winnerChecker := app.checkWinner()
			if winnerChecker == true {
				break
			}
		}
		return
	}()

	//get moderator board
	moderator_board := app.getModeratorBoard()

	// add container to window
	finalContent := container.NewBorder(
		nil,
		moderator_board,
		nil,
		nil,
		container.NewHBox(playerBoardOne, layout.NewSpacer(), playerBoardTwo, layout.NewSpacer(), playerBoardThree),
	)

	app.MainWindow.SetContent(finalContent)
}

func (app *Config) checkWinner() bool {
	var winners []string
	for _, board := range AllPlayersBoards {
		if board.board[0].activated == true && board.board[1].activated == true && board.board[2].activated == true {
			winners = setWinners(winners, board)
		}
		if board.board[3].activated == true && board.board[4].activated == true && board.board[5].activated == true {
			winners = setWinners(winners, board)
		}
		if board.board[6].activated == true && board.board[7].activated == true && board.board[8].activated == true {
			winners = setWinners(winners, board)
		}
		if board.board[0].activated == true && board.board[3].activated == true && board.board[6].activated == true {
			winners = setWinners(winners, board)
		}
		if board.board[1].activated == true && board.board[4].activated == true && board.board[7].activated == true {
			winners = setWinners(winners, board)
		}
		if board.board[2].activated == true && board.board[5].activated == true && board.board[8].activated == true {
			winners = setWinners(winners, board)
		}
		if board.board[0].activated == true && board.board[4].activated == true && board.board[8].activated == true {
			winners = setWinners(winners, board)
		}
		if board.board[2].activated == true && board.board[4].activated == true && board.board[6].activated == true {
			winners = setWinners(winners, board)
		}
	}
	if len(winners) > 0 {
		showWinner(winners, app)
		return true
	}

	return false
}

func setWinners(winners []string, board PlayersBoard) []string {
	if contains(winners, board.playerName) == false {
		winners = append(winners, board.playerName)
	}
	return winners
}

func showWinner(winners []string, app *Config) {

	var winnersString string
	for _, playerName := range winners {
		winnersString += playerName + " "
	}

	dialog.ShowConfirm("WINNER", "Championem został: "+winnersString+"\n Nowa gra wariacie?", func(b bool) {
		if b == false {
			app.MainWindow.Close()
		}
		if b == true {
			AllPlayersBoards = nil
			app.makeUI()
			return
		}
	}, app.MainWindow)
}
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
