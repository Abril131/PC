package scenes

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type EndScenes struct {
	window fyne.Window
	app    fyne.App
}

func NewEndScenes(window fyne.Window, app fyne.App) *EndScenes {
	return &EndScenes{
		window: window,
		app:    app,
	}
}

func (ends *EndScenes) End() {
	backgroundImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/end.png"))
	backgroundImage.Resize(fyne.NewSize(1240, 720))
	backgroundImage.Move(fyne.NewPos(0, 0))

	playButton := widget.NewButton("Volver a jugar", ends.StarEnd)
	playButton.Resize(fyne.NewSize(150, 30))
	playButton.Move(fyne.NewPos(400, 450))

	exitButton := widget.NewButton("Salir", ends.ExitEnd)
	exitButton.Resize(fyne.NewSize(150, 30))
	exitButton.Move(fyne.NewPos(650, 450))

	ends.window.SetContent(container.NewWithoutLayout(backgroundImage, playButton, exitButton))
}

func (ends *EndScenes) StarEnd() {
	scene := NewGamesScene(ends.window, ends.app)
	scene.Show()
}
func (ends *EndScenes) ExitEnd() {
	ends.window.Close()
}
