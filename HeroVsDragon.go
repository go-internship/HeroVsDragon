/* Author - sula7 (sulafpv)
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
var inputMainMenuItem string     //input
var inputLangMenuItem string     //input
var hpHero int = 100
var hpDragon int = 100
var weaponHero int = 0 //input

const harmOfSword = 10
const harmOfArrow = 15
const harmOfFrstn = 30

var heroName string //input

type menuData struct {
	point1, point2, point3, langRU, langEN, loading, bye, incorrectInp string
}

type gameDataText struct {
	hp, dragonMiss, harmHeroToDragon, harmDragonToHero, hero, dragon, entHeroName,
	gameOver, selWeapon, weapon1, weapon2, weapon3 string
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
		bye:          "До скорой встречи!",
		incorrectInp: "Неверный выбор, повторите снова",
	}
	return mainMenu
}

func setGetMainMenuTextEN() menuData {
	mainMenu := menuData{
		point1:       "1. Start a new game",
		point2:       "2. Choose lang",
		point3:       "3. Exit",
		loading:      "Loading...",
		bye:          "Good bye :)",
		incorrectInp: "Incorrect selection, try again",
	}
	return mainMenu
}

func setGetGameDataTextRU() gameDataText {
	data := gameDataText{
		hp:               "hp",
		dragonMiss:       "Дракон промахнулся и не нанёс вам урона :)",
		harmHeroToDragon: "Вы нанесли урон Дракону",
		harmDragonToHero: "Дракон нанёс Вам урон",
		hero:             "Герой",
		dragon:           "Дракон Драконыч",
		entHeroName:      "Введите имя Героя:",
		gameOver:         "Игра завершилась",
		selWeapon:        "Выберите оружие:",
		weapon1:          "1. Меч",
		weapon2:          "2. Стрела",
		weapon3:          "3. Огненный камень",
	}
	return data
}

func setGetGameDataTextEN() gameDataText {
	data := gameDataText{
		hp:               "hp",
		dragonMiss:       "Dragon missed, no harm to Hero :)",
		harmHeroToDragon: "You damaged Dragon to",
		harmDragonToHero: "Dragon damaged You to",
		hero:             "Hero",
		dragon:           "Dragon Dragoner",
		entHeroName:      "Enter Hero name:",
		gameOver:         "Game over",
		selWeapon:        "Select weapon:",
		weapon1:          "1. Sword",
		weapon2:          "2. Arrow",
		weapon3:          "3. Firestone",
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
	fmt.Println(checkLangwReturn(setGetMainMenuTextRU().point1, setGetMainMenuTextEN().point1))
	fmt.Println(checkLangwReturn(setGetMainMenuTextRU().point2, setGetMainMenuTextEN().point2))
	fmt.Println(checkLangwReturn(setGetMainMenuTextRU().point3, setGetMainMenuTextEN().point3))
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
		fmt.Println(checkLangwReturn(setGetMainMenuTextRU().incorrectInp, setGetMainMenuTextEN().incorrectInp))
		showLangMenu()
		selectLangMenuItem()
	}
}

func checkLangwReturn(textRU string, textEN string) string {
	if selectedMenuLang == true {
		return textRU
	}
	return textEN
}

func selectMainMenuItem() { //Основные действия в меню
	someThing := bufio.NewScanner(os.Stdin)
	someThing.Scan()
	inputMainMenuItem = someThing.Text()
	switch inputMainMenuItem {
	case "1": //Начать новую игру
		fmt.Println(checkLangwReturn(setGetMainMenuTextRU().loading, setGetMainMenuTextEN().loading))
		isGameStart = true
	case "2": //Выбрать язык
		showLangMenu()
		selectLangMenuItem()
	case "3": //Выход
		fmt.Println(checkLangwReturn(setGetMainMenuTextRU().bye, setGetMainMenuTextEN().bye))
		os.Exit(0)
	default:
		fmt.Println(checkLangwReturn(setGetMainMenuTextRU().incorrectInp, setGetMainMenuTextEN().incorrectInp))
		selectMainMenuItem()
	}
}

func gameStart() {
	fmt.Println(checkLangwReturn(setGetGameDataTextRU().entHeroName, setGetGameDataTextEN().entHeroName))
	inputHeroName()
	for {
		if !isGameEnd {
			showGameResult()
			showWeaponHero()
			selectWeapon()
			attackToDragonName()
			checkCurrentHp()
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
	fmt.Println("")
}

func showGameResult() {
	fmt.Println(checkLangwReturn(setGetGameDataTextRU().hero, setGetGameDataTextEN().hero), heroName, "\t\t\t", checkLangwReturn(setGetGameDataTextRU().dragon, setGetGameDataTextEN().dragon))
	showCurrentHP()
}

func showCurrentHP() {
	fmt.Println(hpHero, setGetGameDataTextRU().hp, "\t\t\t\t", hpDragon, setGetGameDataTextEN().hp)
}

func checkCurrentHp() {
	if hpHero|hpDragon < 1 {
		isGameEnd = true
	}
}

func gameEnd() {
	fmt.Println("\n\n")
	fmt.Println(checkLangwReturn(setGetGameDataTextRU().gameOver, setGetGameDataTextEN().gameOver))
	showGameResult()
}

func showWeaponHero() {
	fmt.Println("\n")
	fmt.Println(setGetGameDataTextRU().selWeapon)
	weaponHero := [3]string{
		checkLangwReturn(setGetGameDataTextRU().weapon1, setGetGameDataTextEN().weapon1),
		checkLangwReturn(setGetGameDataTextRU().weapon2, setGetGameDataTextEN().weapon2),
		checkLangwReturn(setGetGameDataTextRU().weapon3, setGetGameDataTextEN().weapon3),
	}
	for i := 0; i < len(weaponHero); i++ {
		fmt.Println(weaponHero[i])
	}
}

func selectWeapon() {
	fmt.Scan(&weaponHero)
	fmt.Println("\n")
}

func randomize(min int, max int) int {
	return min + rand.Intn(max-min)
}

func attackToDragonName() {
	rand.Seed(time.Now().UTC().UnixNano())
	switch weaponHero {
	case 1:
		randomized := randomize(0, 10)
		casesAttackToDragon(randomized, harmOfSword)
	case 2:
		randomized := randomize(10, 20)
		casesAttackToDragon(randomized, harmOfArrow)
	case 3:
		randomized := randomize(20, 30)
		casesAttackToDragon(randomized, harmOfFrstn)
	default:
		fmt.Println(checkLangwReturn(setGetMainMenuTextRU().incorrectInp, setGetMainMenuTextEN().incorrectInp))
	}
}

func casesAttackToDragon(randomized int, harm int) {
	hpHero = hpHero - randomized
	hpDragon = hpDragon - harm
	fmt.Println(checkLangwReturn(setGetGameDataTextRU().harmHeroToDragon, setGetGameDataTextEN().harmHeroToDragon), harm)
	if randomized == 0 {
		fmt.Println(checkLangwReturn(setGetGameDataTextRU().dragonMiss, setGetGameDataTextEN().dragonMiss))
	} else {
		fmt.Println(checkLangwReturn(setGetGameDataTextRU().harmDragonToHero, setGetGameDataTextEN().harmDragonToHero), randomized)
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
