package main

import (
  "os"

  "github.com/doxxan/appindicator"
  "github.com/doxxan/appindicator/gtk-extensions/gotk3"
  "github.com/johnmccabe/openfaas-bitbar/config"
  "github.com/conformal/gotk3/gtk"
  "gopkg.in/yaml.v2"
)

var (
  mainMenu *gtk.Menu
  images [10]*gtk.Image
)

func main() {
  // Read config (I've no care for writing one)
  newCfg, _ := config.Read()
  config.EnsureConfigDir()
  f, err := os.OpenFile(config.File(), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
  if err != nil {
    os.Exit(1)
  }

  defer f.Close()

  y, _ := yaml.Marshal(newCfg)

  if _, err = f.Write(y); err != nil {
    os.Exit(1)
  }

  if len(newCfg.Stacks) < 1 {
    // Do Nothing (maybe report error in future)
  }

  gtk.Init(nil)

  // menu init
  mainMenu, _ = gtk.MenuNew()

  // test putting picture in menu...
  i, _ := gtk.ImageNewFromFile("/home/lewis/Pictures/aaliyah.jpg")
  images[0] = i

  for _, stack := range newCfg.Stacks {
    mainMenuSub, _ := gtk.MenuItemNewWithLabel(stack.Name)
    mainMenuSub.Show()

    stackMenu, _ := gtk.MenuNew()

    menuItem1, _ := gtk.MenuItemNewWithLabel("Function Rate - ##")
    stackMenu.Append(menuItem1)

    menuItem2, _ := gtk.MenuItemNewWithLabel("Uptime - ##:##:##")
    stackMenu.Append(menuItem2)

    mainMenuSub.SetSubmenu(stackMenu)
    mainMenuSub.ShowAll()
    mainMenu.Append(mainMenuSub)
  }

  menuItemImg, _ := gtk.MenuItemNew()
  menuItemImg.Add(images[0])
  mainMenu.Append(menuItemImg)

  menuQuit, _ := gtk.MenuItemNewWithLabel("Quit")
  menuQuit.Connect("activate", func() {
    gtk.MainQuit()
  })
  menuQuit.Show()
  mainMenu.Append(menuQuit)

  // indicator
  indicator := gotk3.NewAppIndicator("test-openfaas-indicator-ubuntu", "system-run-symbolic", appindicator.CategoryCommunications)
  indicator.SetStatus(appindicator.StatusActive)
  indicator.SetMenu(mainMenu)
  mainMenu.ShowAll()
  gtk.Main()
}
