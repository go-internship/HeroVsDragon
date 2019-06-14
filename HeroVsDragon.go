/* Author - sula7 (Sultan Moldobaev)
Please read a manual on GitHub on how to run this application
*/

package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var menuItemMainMenuRU [3]string
var menuItemMainMenuEN [3]string

var menuItemLangMenu [3]string
var isGameStart bool = false

var selectedMenuLang bool = true //true = RU, false = EN
var inputMainMenuItem int
var inputLangMenuItem int
var hpHero int = 100
var hpDragon int = 100
var weaponHero int = 0

var heroName string

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
		selectedMenuLang = true
	case 2:
		selectedMenuLang = false
	default:
		checkLang("Неверный выбор, введите снова!", "Incorrect selection, try again")
		showLangMenu()
		selectLangMenuItem()
	}
}

func checkLang(textRU string, textEN string) {
	if selectedMenuLang == true {
		fmt.Println(textRU)
	} else if selectedMenuLang == false {
		fmt.Println(textEN)
	}
}

func selectMainMenuItem() { //Основные действия в меню
	fmt.Scan(&inputMainMenuItem)
	switch inputMainMenuItem {
	case 1: //Начать новую игру
		checkLang("Загружается...", "Starting...")
		isGameStart = true
	case 2: //Выбрать язык
		showLangMenu()
		selectLangMenuItem()
	case 3: //Выход
		checkLang("До скорой встречи!", "Good Bye!")
		os.Exit(0)
	default:
		checkLang("Неверный выбор, введите снова!", "Incorrect selection, try again")
		selectMainMenuItem()
	}
}

func gameStart() {
	checkLang("Введите имя Героя:", "Enter Hero name:")
	inputHeroName()
	showWeaponHero()
	selectWeapon()
	attackToDragon()
}

func inputHeroName() {
	fmt.Scan(&heroName)
}

func showGameResultRU() { //Сделать перевод
	fmt.Println("Герой", heroName, "\t\t\t", "Дракон Драконович")
	showCurrentHP()
}

func showCurrentHP() {
	fmt.Println(hpHero, "hp", "\t\t\t\t", hpDragon, "hp")
}

func showWeaponHero() {
	fmt.Println("\n")
	fmt.Println("Выберите оружие:")
	weaponHero := [3]string{"1. Меч", "2. Стрела", "3. Огенный камень"}
	for i := 0; i < len(weaponHero); i++ {
		fmt.Println(weaponHero[i])
	}
}

func selectWeapon() {
	fmt.Scan(&weaponHero)
}

func attackToDragon() {
	rand.Seed(time.Now().UnixNano())
	switch weaponHero {
	case 1:
		hpDragon = hpDragon - 10
		fmt.Println("Вы нанесли Дракону 10 урона")
		randomized := rand.Intn(20)
		hpHero = hpHero - randomized
		fmt.Println("Дракон нанёс вам", randomized, "урона")
	}
}

func main() {
	for {
		if selectedMenuLang {
			showMainMenuRU()
			selectMainMenuItem()
			if isGameStart {
				break
			}
		} else if !selectedMenuLang {
			showMainMenuEN()
			selectMainMenuItem()
			if isGameStart {
				break
			}
		}
	}
	gameStart()
}
