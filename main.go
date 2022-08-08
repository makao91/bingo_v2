package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"log"
	"os"
)

type Config struct {
	App             fyne.App
	MainWindow      fyne.Window
	InfoLog         *log.Logger
	ErrorLog        *log.Logger
	ModeratorFields [][]interface{}
	ModeratorTable  *widget.Table
}

func main() {

	var myApp Config

	//create application
	fyneApp := app.NewWithID("makao320")
	myApp.App = fyneApp

	//create loggers
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	myApp.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// create and size a fyne window
	myApp.MainWindow = fyneApp.NewWindow("GoldWatcher")
	myApp.MainWindow.Resize(fyne.NewSize(770, 410))
	myApp.MainWindow.SetFixedSize(true)
	myApp.MainWindow.SetMaster()

	myApp.makeUI()

	// show and run the application
	myApp.MainWindow.ShowAndRun()
}
