package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type BingoField struct {
	name      string
	activated bool
}

func (app *Config) getModeratorBoard() *fyne.Container {
	app.ModeratorFields = app.getFieldsToPlayString()
	app.ModeratorTable = app.getModeratorTable()

	holdingsContainer := container.NewBorder(
		nil,
		nil,
		nil,
		nil,
		container.NewAdaptiveGrid(1, app.ModeratorTable),
	)

	//holdingsContainer.Resize(fyne.NewSize(300, 300))

	return holdingsContainer
}

func (app *Config) getModeratorTable() *widget.Table {
	t := widget.NewTable(
		func() (int, int) {
			return len(app.ModeratorFields), len(app.ModeratorFields[0])
		},
		func() fyne.CanvasObject {
			ctr := container.NewVBox(widget.NewLabel(""))
			return ctr
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {

			// we're just putting in textual information
			o.(*fyne.Container).Objects = []fyne.CanvasObject{
				widget.NewLabel(app.ModeratorFields[i.Row][i.Col].(string)),
			}
		})

	colWidths := []float32{100, 100, 100}
	for i := 0; i < len(colWidths); i++ {
		t.SetColumnWidth(i, colWidths[i])
	}

	return t
}

func (app *Config) getFieldsToPlay() [][]BingoField {

	var board = [][]BingoField{
		{BingoField{"kod Witka", false}, BingoField{"Teamsy", false}, BingoField{"Unitree", false}},
		{BingoField{"pogoda", false}, BingoField{"AgileBoard", false}, BingoField{"TimeTracker", false}},
		{BingoField{"BitBucket", false}, BingoField{"syf w kiblu", false}, BingoField{"syf w kuchni", false}},
	}

	return board
}

func (app *Config) getFieldsToPlayString() [][]interface{} {

	var slice [][]interface{}

	slice = append(slice, []interface{}{"kod Witka", "Teamsy", "Unitree"})
	slice = append(slice, []interface{}{"pogoda", "AgileBoard", "TimeTracker"})
	slice = append(slice, []interface{}{"BitBucket", "syf w kiblu", "syf w kuchni"})

	return slice
}
