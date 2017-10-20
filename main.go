package main

import (
	"github.com/doxxan/appindicator"
	"github.com/doxxan/appindicator/gtk-extensions/gotk3"
	"github.com/conformal/gotk3/gtk"
)

func main() {
	gtk.Init(nil)

	// menu init
	menu, _ := gtk.MenuNew()

	menuItem1, _ := gtk.MenuItemNewWithLabel("Function Rate - ##")
	menuItem1.Show()
	menu.Append(menuItem1)

	menuItem2, _ := gtk.MenuItemNewWithLabel("Uptime - ##:##:##")
	menuItem2.Show()
	menu.Append(menuItem2)

	menuQuit, _ := gtk.MenuItemNewWithLabel("Quit")
	menuQuit.Connect("activate", func() {
		gtk.MainQuit()
	})
	menuQuit.Show()
	menu.Append(menuQuit)

	// indicator
	indicator := gotk3.NewAppIndicator("test-openfaas-indicator-ubuntu", "system-run-symbolic", appindicator.CategoryCommunications)
	indicator.SetStatus(appindicator.StatusActive)
	indicator.SetMenu(menu)

	gtk.Main()
}
