package scenes

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type WinScenes struct {
	window fyne.Window
	app    fyne.App
}

func NewWinScenes(window fyne.Window, app fyne.App) *WinScenes {
	return &WinScenes{
		window: window,
		app:    app,
	}
}

func (wi *WinScenes) Win() {
	backgroundImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/ganado.png"))
	backgroundImage.Resize(fyne.NewSize(1240, 720))
	backgroundImage.Move(fyne.NewPos(0, 0))

	playButton := widget.NewButton("Volver a jugar", wi.StarWin)
	playButton.Resize(fyne.NewSize(150, 30))
	playButton.Move(fyne.NewPos(400, 450))

	exitButton := widget.NewButton("Salir", wi.ExitWin)
	exitButton.Resize(fyne.NewSize(150, 30))
	exitButton.Move(fyne.NewPos(650, 450))

	wi.window.SetContent(container.NewWithoutLayout(backgroundImage, playButton, exitButton))
}

func (wi *WinScenes) StarWin() {
	scene := NewGamesScene(wi.window, wi.app)
	scene.Show()
}
func (wi *WinScenes) ExitWin() {
	wi.window.Close()
}
