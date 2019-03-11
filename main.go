package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

// Keyboard struct
type Keyboard struct {
	buttons []*widget.Button

	window fyne.Window
}

var qwertyLayout = [...]string{
	"tab", "Q", "W", "E", "R", "T", "Y", "U", "I", "O", "P", "'",
	"esc", "A", "S", "D", "F", "G", "H", "J", "K", "L", ";", "ret",
	"(", "Z", "X", "C", "V", "B", "N", "M", ",", ".", "/", ")",
	"[", "{", "<", "@", " ", " ", " ", " ", "@", ">", "}", "]",
}

var colemakLayout = [...]string{
	"", "Q", "W", "F", "P", "G", "J", "L", "U", "Y", ";", "",
	"", "A", "R", "S", "T", "D", "H", "N", "E", "I", "O", "",
	"", "Z", "X", "C", "V", "B", "K", "M", ",", ".", "/", "",
	"", "", "", "", "", "", "", "", "", "", "", "",
}

var numberLayout = [...]string{
	"", "Q", "W", "F", "P", "G", "J", "L", "U", "Y", ";", "",
	"", "A", "R", "S", "T", "D", "H", "N", "E", "I", "O", "",
	"", "Z", "X", "C", "V", "B", "K", "M", ",", ".", "/", "",
	"", "", "", "", "", "", "", "", "", "", "", "",
}

func (k *Keyboard) addButton(text string) *widget.Button {
	action := func() {}
	button := widget.NewButton(text, action)
	k.buttons = append(k.buttons, button)
	return button
}

func (k *Keyboard) loadUI(app fyne.App) {

	k.window = app.NewWindow("Plank")
	k.window.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewGridLayout(1),
			fyne.NewContainerWithLayout(
				layout.NewGridLayout(12),
				k.addButton("tab"),
				k.addButton("Q"),
				k.addButton("W"),
				k.addButton("E"),
				k.addButton("R"),
				k.addButton("T"),
				k.addButton("Y"),
				k.addButton("U"),
				k.addButton("I"),
				k.addButton("O"),
				k.addButton("P"),
				k.addButton("'"),
			),
			fyne.NewContainerWithLayout(
				layout.NewGridLayout(12),
				k.addButton("esc"),
				k.addButton("A"),
				k.addButton("S"),
				k.addButton("D"),
				k.addButton("F"),
				k.addButton("G"),
				k.addButton("H"),
				k.addButton("J"),
				k.addButton("K"),
				k.addButton("L"),
				k.addButton(";"),
				k.addButton("ret"),
			),
			fyne.NewContainerWithLayout(
				layout.NewGridLayout(12),
				k.addButton("("),
				k.addButton("Z"),
				k.addButton("X"),
				k.addButton("C"),
				k.addButton("V"),
				k.addButton("B"),
				k.addButton("N"),
				k.addButton("M"),
				k.addButton(","),
				k.addButton("."),
				k.addButton("/"),
				k.addButton(")"),
			),
			fyne.NewContainerWithLayout(
				layout.NewGridLayout(12),
				k.addButton("["),
				k.addButton("{"),
				k.addButton("<"),
				k.addButton("@"),
				k.addButton(" "),
				layout.NewSpacer(),
				layout.NewSpacer(),
				k.addButton(" "),
				k.addButton("@"),
				k.addButton(">"),
				k.addButton("}"),
				k.addButton("]"),
			),
		),
	)
}

func main() {
	a := app.New()

	k := &Keyboard{}
	k.loadUI(a)

	k.window.ShowAndRun()
	fmt.Println("Exit")
}
