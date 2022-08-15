package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"time"
)

func (app *Config) checkForBingo(board PlayersBoard) {
	go func() {
		for range time.Tick(time.Second * 1) {
			var winners []string
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
			if len(winners) > 0 {
				showWinner(winners, app)
				break
			}
		}
		return
	}()
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
			app.AllPlayersBoards = nil
			app.AllPlayerContainer = nil
			app.FinalPlayerContent = nil
			app.makeUI()
			app.MainWindow.Resize(fyne.NewSize(770, 410))
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
