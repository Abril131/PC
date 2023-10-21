package scenes

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type MainScenes struct {
	window fyne.Window
	app    fyne.App
}

func NewMainScenes(window fyne.Window, app fyne.App) *MainScenes {
	return &MainScenes{
		window: window,
		app:    app,
	}
}

func (sce *MainScenes) Inicio() {
	backgroundImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/menu.jpeg"))
	backgroundImage.Resize(fyne.NewSize(1240, 720))
	backgroundImage.Move(fyne.NewPos(0, 0))

	playButton := widget.NewButton("Jugar", sce.StarMain)
	playButton.Resize(fyne.NewSize(150, 30))
	playButton.Move(fyne.NewPos(400, 450))

	exitButton := widget.NewButton("Salir", sce.ExitMain)
	exitButton.Resize(fyne.NewSize(150, 30))
	exitButton.Move(fyne.NewPos(650, 450))

	sce.window.SetContent(container.NewWithoutLayout(backgroundImage, playButton, exitButton))
}

func (sce *MainScenes) StarMain() {
	scene := NewGamesScene(sce.window, sce.app)
	scene.Show()
}
func (sce *MainScenes) ExitMain() {
	sce.window.Close()
}
