package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)
	defer gtk.MainQuit()

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("My Window")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	label, err := gtk.LabelNew("Hello, world!")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}
	win.Add(label)

	win.ShowAll()
	gtk.Main()
}
