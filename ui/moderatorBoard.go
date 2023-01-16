package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/joho/godotenv"
	"log"
	"os"
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
		{fieldName: "języki programowania", activated: false},
	},
	{
		{fieldName: "anime", activated: false},
		{fieldName: "żeglarz", activated: false},
		{fieldName: "opluwanie pracowników", activated: false},
	},
	{
		{fieldName: "wordpress", activated: false},
		{fieldName: "dedlajny", activated: false},
		{fieldName: "januszpol", activated: false},
	},
	{
		{fieldName: "jebanie na czyiś kod", activated: false},
		{fieldName: "klienci", activated: false},
		{fieldName: "testy", activated: false},
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
	for r := 0; r < len(fullModeratorBoard); r++ {
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
		app.createModeratorPasswordDialog(i)
	}
}

func (app *Config) createModeratorPasswordDialog(i *ModeratorBoardButton) {
	password := widget.NewPasswordEntry()
	password.Validator = validation.NewRegexp("champ", "Only true champion know the password")
	// create a dialog
	addForm := dialog.NewForm(
		"Champion Board",
		"Confirm",
		"Cancel",
		[]*widget.FormItem{
			{Text: "Password", Widget: password},
		},
		func(valid bool) {

			if valid {
				app.activateModeratorButton(i)
			}
		},
		app.MainWindow)

	// size and show the dialog
	addForm.Resize(fyne.Size{Width: 400})
	addForm.Show()
	app.MainWindow.Canvas().Focus(password)
}

func (app *Config) activateModeratorButton(i *ModeratorBoardButton) {
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

func (app *Config) getEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
