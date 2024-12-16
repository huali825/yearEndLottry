package main

import (
	"fmt"
	"fyne.io/fyne/v2/widget"
	"testing"
)

func TestName(t *testing.T) {
	fmt.Printf("hello")
	out, _ := makeUI()

	if out.Text != "Hello world!" {
		t.Error("Incorrect initial greeting")
	}

}

func makeUI() (*widget.Label, *widget.Entry) {
	out := widget.NewLabel("Hello world!")
	in := widget.NewEntry()

	in.OnChanged = func(content string) {
		out.SetText("Hello " + content + "!")
	}
	return out, in
}
