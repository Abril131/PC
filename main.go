// main.go
package main

import (
	"Juego/scenes"
	"fmt"
	"fyne.io/fyne/v2"
)

func main() {
	//Cargar y mostrar la escena principal
	mainW := scenes.NewMainScene("HOBILAN", fyne.NewSize(1240, 720))
	mainW.ShowMenu()
	fmt.Println("hhh")
}
