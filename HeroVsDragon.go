package main

import (
	"fmt"
	"os"
)

var menuItemMainMenuRU [3]string
var menuItemMainMenuEN [3]string

var menuItemLangMenu [3]string

var selectedMenuLang int = 2 //1 = RU, 2 = EN
var inputMainMenuItem int
var inputLangMenuItem int

func showMainMenuItemsRU() {
	menuItemMainMenuRU[0] = "1. Начать новую игру"
	menuItemMainMenuRU[1] = "2. Выбрать язык"
	menuItemMainMenuRU[2] = "3. Выход"
}

func showMainMenuItemsEN() {
	menuItemMainMenuEN[0] = "1. Start a new game"
	menuItemMainMenuEN[1] = "2. Choose lang"
	menuItemMainMenuEN[2] = "3. Exit"
}

func showLangMenuItems() {
	menuItemLangMenu[0] = "1. Русский"
	menuItemLangMenu[1] = "2. English"
}

func showLangMenu() {
	showLangMenuItems()
	for i := 0; i < len(menuItemLangMenu); i++ {
		fmt.Println(menuItemLangMenu[i])
	}
}

func showMainMenuRU() {
	showMainMenuItemsRU()
	for i := 0; i < len(menuItemMainMenuRU); i++ {
		fmt.Println(menuItemMainMenuRU[i])
	}
}

func showMainMenuEN() {
	showMainMenuItemsEN()
	for i := 0; i < len(menuItemMainMenuEN); i++ {
		fmt.Println(menuItemMainMenuEN[i])
	}
}

func selectLangMenuItem() {
	fmt.Scan(&inputLangMenuItem)
	switch inputLangMenuItem {
	case 1:
		selectedMenuLang = 1
	case 2:
		selectedMenuLang = 2
	}
}

func selectMainMenuItem() {
	fmt.Scan(&inputMainMenuItem)
	switch inputMainMenuItem {
	case 1:
		fmt.Println("Новая игра загружается")
	case 2:
		showLangMenu()
		selectLangMenuItem()
	case 3:
		fmt.Println("Всего доброго!")
		os.Exit(2)
	default:
		fmt.Println("Неверный выбор")
	}
}

func main() {
	if selectedMenuLang == 1 {
		showMainMenuRU()
	} else if selectedMenuLang == 2 {
		showMainMenuEN()
	}
	selectMainMenuItem()
}
