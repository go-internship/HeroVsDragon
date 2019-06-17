/* Author - sula7 (Sultan Moldobaev)
Please read a manual on GitHub on how to run this application
*/

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var isGameStart bool = false
var isGameEnd bool = false

var selectedMenuLang bool = true //true = RU, false = EN
var inputMainMenuItem string
var inputLangMenuItem string
var hpHero int = 100
var hpDragon int = 100
var weaponHero int = 0

const harmOfSword = 10
const harmOfArrow = 15
const harmOfFrstn = 30

var heroName string

type menuData struct {
	point1, point2, point3, langRU, langEN, loading, hero, dragonName, bye, incorrectInp, entHeroName,
	gameOver, selWeapon, weapon1, weapon2, weapon3 string
}

type gameDataText struct {
	hp string
}

type gameDataLogic struct {
	harmOfSword, harmOfArrow, harmOfFrstn int
}

func setGetMainMenuTextRU() menuData {
	mainMenu := menuData{
		point1:       "1. Начать новую игру",
		point2:       "2. Выбрать язык",
		point3:       "3. Выход",
		langRU:       "1. Русский",
		langEN:       "2. English",
		loading:      "Загружается...",
		hero:         "Герой",
		dragonName:   "Дракон Драконыч",
		bye:          "До скорой встречи",
		incorrectInp: "Неверный выбор, повторите снова",
		entHeroName:  "Введите имя Героя:",
		gameOver:     "Игра завершилась",
		selWeapon:    "Выберите оружие:",
		weapon1:      "1. Меч",
		weapon2:      "2. Стрела",
		weapon3:      "3. Огненный камень",
	}
	return mainMenu
}

func setGetMainMenuTextEN() menuData {
	mainMenu := menuData{
		point1:       "1. Start a new game",
		point2:       "2. Choose lang",
		point3:       "3. Exit",
		loading:      "Loading...",
		hero:         "Hero",
		dragonName:   "DragonName DragonNameer",
		bye:          "Good bye :)",
		incorrectInp: "Incorrect selection, try again",
		entHeroName:  "Enter Hero name:",
		gameOver:     "Game over",
		selWeapon:    "Select weapon:",
		weapon1:      "1. Sword",
		weapon2:      "2. Arrow",
		weapon3:      "3. Firestone",
	}
	return mainMenu
}

func setGetGameDataTextRU() gameDataText {
	data := gameDataText{
		hp: "hp",
	}
	return data
}

func setGetGameDataTextEN() gameDataText {
	data := gameDataText{
		hp: "hp",
	}
	return data
}

func setGetGameDataLogic() gameDataLogic {
	data := gameDataLogic{
		harmOfSword: 10,
		harmOfArrow: 15,
		harmOfFrstn: 30,
	}
	return data
}

func showLangMenu() {
	fmt.Println(setGetMainMenuTextRU().langRU)
	fmt.Println(setGetMainMenuTextRU().langEN)
}

func showMainMenu() {
	checkLang(setGetMainMenuTextRU().point1, setGetMainMenuTextEN().point1)
	checkLang(setGetMainMenuTextRU().point2, setGetMainMenuTextEN().point2)
	checkLang(setGetMainMenuTextRU().point3, setGetMainMenuTextEN().point3)
}

func selectLangMenuItem() {
	someThing := bufio.NewScanner(os.Stdin)
	someThing.Scan()
	inputLangMenuItem = someThing.Text()
	switch inputLangMenuItem {
	case "1":
		selectedMenuLang = true
	case "2":
		selectedMenuLang = false
	default:
		checkLang(setGetMainMenuTextRU().incorrectInp, setGetMainMenuTextEN().incorrectInp)
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
	someThing := bufio.NewScanner(os.Stdin)
	someThing.Scan()
	inputMainMenuItem = someThing.Text()
	switch inputMainMenuItem {
	case "1": //Начать новую игру
		checkLang(setGetMainMenuTextRU().loading, setGetMainMenuTextEN().loading)
		isGameStart = true
	case "2": //Выбрать язык
		showLangMenu()
		selectLangMenuItem()
	case "3": //Выход
		checkLang(setGetMainMenuTextRU().bye, setGetMainMenuTextEN().bye)
		os.Exit(0)
	default:
		checkLang(setGetMainMenuTextRU().incorrectInp, setGetMainMenuTextEN().incorrectInp)
		selectMainMenuItem()
	}
}

func gameStart() {
	checkLang(setGetMainMenuTextRU().entHeroName, setGetMainMenuTextEN().entHeroName)
	inputHeroName()
	for {
		if !isGameEnd {
			showGameResultRU()
			showWeaponHero()
			selectWeapon()
			attackToDragonName()
		} else if isGameEnd {
			gameEnd()
			break
		}
	}
}

func inputHeroName() {
	someThing := bufio.NewScanner(os.Stdin)
	someThing.Scan()
	heroName = someThing.Text()
}

func showGameResultRU() { //Сделать перевод
	fmt.Println(setGetMainMenuTextRU().hero, heroName, "\t\t\t", setGetMainMenuTextRU().dragonName)
	showCurrentHP()
}

func showCurrentHP() {
	fmt.Println(hpHero, setGetGameDataTextRU().hp, "\t\t\t\t", hpDragon, setGetGameDataTextRU().hp)
	checkCurrentHp()
}

func checkCurrentHp() {
	if hpHero|hpDragon < 1 {
		isGameEnd = true
	}
}

func gameEnd() {
	fmt.Println("\n\n")
	fmt.Println(setGetMainMenuTextRU().gameOver)
	showGameResultRU()
}

func showWeaponHero() {
	fmt.Println("\n")
	fmt.Println(setGetMainMenuTextRU().selWeapon)
	weaponHero := [3]string{setGetMainMenuTextRU().weapon1, setGetMainMenuTextRU().weapon2, setGetMainMenuTextRU().weapon3}
	for i := 0; i < len(weaponHero); i++ {
		fmt.Println(weaponHero[i])
	}
}

func selectWeapon() {
	fmt.Scan(&weaponHero)
	fmt.Println("\n")
}

func randmize(min int, max int) int {
	return min + rand.Intn(max-min)
}

func attackToDragonName() {
	rand.Seed(time.Now().UTC().UnixNano())
	switch weaponHero {
	case 1:
		caseOneAttackToDragonName()
	case 2:
		caseTwoAttackToDragonName()
	case 3:
		caseThreeAttackToDragonName()
	default:
		checkLang(setGetMainMenuTextRU().incorrectInp, setGetMainMenuTextEN().incorrectInp)
	}
}

func caseOneAttackToDragonName() {
	randomized := randmize(0, 10)
	hpHero = hpHero - randomized
	hpDragon = hpDragon - harmOfSword
	fmt.Println("Вы нанесли Дракону 10 урона")
	if randomized == 0 {
		fmt.Println("Дракон промахнулся и не нанёс вам урона :)")
	} else {
		fmt.Println("Дракон нанёс вам", randomized, "урона")
	}
	fmt.Println("\n")
}

func caseTwoAttackToDragonName() {
	randomized := randmize(10, 20)
	hpHero = hpHero - randomized
	hpDragon = hpDragon - harmOfArrow
	fmt.Println("Вы нанесли Дракону 15 урона")
	if randomized == 0 {
		fmt.Println("Дракон промахнулся и не нанёс вам урона :)")
	} else {
		fmt.Println("Дракон нанёс вам", randomized, "урона")
	}
	fmt.Println("\n")
}

func caseThreeAttackToDragonName() {
	randomized := randmize(20, 30)
	hpHero = hpHero - randomized
	hpDragon = hpDragon - harmOfFrstn
	fmt.Println("Вы нанесли Дракону 30 урона")
	if randomized == 0 {
		fmt.Println("Дракон промахнулся и не нанёс вам урона :)")
	} else {
		fmt.Println("Дракон нанёс вам", randomized, "урона")
	}
	fmt.Println("\n")
}

func main() {
	for {
		showMainMenu()
		selectMainMenuItem()
		if isGameStart {
			break
		}
	}
	gameStart()
}
