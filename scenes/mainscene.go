package scenes

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type MainScene struct {
	window fyne.Window
	app    fyne.App
}

func NewMainScene(title string, size fyne.Size) *MainScene {
	mainApp := app.New()
	mainW := mainApp.NewWindow(title)
	mainW.CenterOnScreen()
	mainW.SetFixedSize(true)
	mainW.Resize(size)
	return &MainScene{
		app:    mainApp,
		window: mainW,
	}
}

func (m *MainScene) ShowMenu() {
	scene := NewMainScenes(m.window, m.app)
	scene.Inicio()
	m.window.ShowAndRun()

}
