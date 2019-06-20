/* Author - sula7 (@sulafpv)
Please read a manual on GitHub on how to run this application
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
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

type menuData struct { //Тексты главного меню
	point1, point2, point3, langRU, langEN, loading, bye, incorrectInp string
}

type gameDataText struct { //Игровые тексты
	hp, dragonMiss, harmHeroToDragon, harmDragonToHero, hero, dragon, entHeroName,
	gameOver, selWeapon, weapon1, weapon2, weapon3, winner, step, standoff string
}

func setGetMainMenuText() menuData { //Тексты главного меню
	mainMenu := menuData{
		point1:       "1. Начать новую игру",
		point2:       "2. Выход",
		loading:      "Загружается...",
		bye:          "До скорой встречи!",
		incorrectInp: "Неверный выбор, повторите снова",
	}
	return mainMenu
}

func setGetGameDataText() gameDataText { //Игровые тексты
	data := gameDataText{
		hp:               "hp",
		dragonMiss:       "Дракон промахнулся и не нанёс вам урона :)",
		harmHeroToDragon: "Вы нанесли урон Дракону",
		harmDragonToHero: "Дракон нанёс Вам урон",
		hero:             "Герой",
		dragon:           "Дракон Драконыч",
		entHeroName:      "Введите имя Героя: (нажмите Enter для случайного)",
		gameOver:         "Игра завершилась",
		selWeapon:        "Выберите оружие:",
		weapon1:          "1. Меч",
		weapon2:          "2. Стрела",
		weapon3:          "3. Огненный камень",
		winner:           "Победил",
		step:             "Ход #",
		standoff:         "Победила дружба :)",
	}
	return data
}

func showMainMenu() {
	fmt.Println(setGetMainMenuText().point1)
	fmt.Println(setGetMainMenuText().point2)
}

func selectMainMenuItem() {
	someThing := bufio.NewScanner(os.Stdin)
	someThing.Scan()
	inputMainMenuItem = strings.TrimSpace(someThing.Text()) //Убирает пробелы в начале и в конце
	switch inputMainMenuItem {
	case "1": //Начать новую игру
		fmt.Println(setGetMainMenuText().loading)
		isGameStart = true
	case "2": //Выход
		fmt.Println(setGetMainMenuText().bye)
		os.Exit(0)
	default:
		fmt.Println(setGetMainMenuText().incorrectInp)
		selectMainMenuItem()
	}
}

func gameStart() {
	fmt.Println(setGetGameDataText().entHeroName)
	inputHeroName()
	for i := 0; ; i++ {
		if !isGameEnd {
			showGameResult()
			showStep(i)
			showWeaponHero()
			selectWeapon()
			attackToDragon()
			checkCurrentHp()
		} else if isGameEnd {
			gameEnd()
			showWinner()
			break
		}
	}
}

func inputHeroName() {
	someThing := bufio.NewScanner(os.Stdin)
	someThing.Scan()
	if someThing.Text() == `` { //для обработки пустой строки
		data := map[string]string{}
		resp, err := http.Get("https://uinames.com/api/?amount=1&gender=male&region=kyrgyz+republic")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &data)
		heroName = data["name"]
	} else {
		heroName = strings.TrimSpace(someThing.Text()) //Убирает пробелы в начале и в конце
	}
	fmt.Println("")
}

func showGameResult() {
	fmt.Println(setGetGameDataText().hero, heroName,
		"\t\t\t", setGetGameDataText().dragon)
	showCurrentHP()
}

func showCurrentHP() {
	fmt.Println(hpHero, setGetGameDataText().hp, "\t\t\t\t", hpDragon)
}

func checkCurrentHp() {
	if hpHero|hpDragon < 1 {
		isGameEnd = true
	}
}

func showStep(step int) {
	fmt.Println(setGetGameDataText().step, step+1)
}

func gameEnd() {
	fmt.Println("\n\n")
	fmt.Println(setGetGameDataText().gameOver)
	showGameResult()
}

func showWinner() {
	if hpHero > hpDragon {
		fmt.Println("")
		fmt.Println(setGetGameDataText().winner, setGetGameDataText().hero, heroName)
	} else if hpDragon > hpHero {
		fmt.Println(setGetGameDataText().winner, setGetGameDataText().dragon)
	} else if hpDragon == hpHero {
		fmt.Println(setGetGameDataText().standoff)
	}
}

func showWeaponHero() {
	fmt.Println("\n")
	fmt.Println(setGetGameDataText().selWeapon)
	weaponHero := [3]string{
		setGetGameDataText().weapon1,
		setGetGameDataText().weapon2,
		setGetGameDataText().weapon3,
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

func attackToDragon() {
	rand.Seed(time.Now().UTC().UnixNano())
	switch weaponHero {
	case 1:
		randomized := randomize(0, 20)
		casesAttackToDragon(randomized, harmOfSword)
	case 2:
		randomized := randomize(10, 30)
		casesAttackToDragon(randomized, harmOfArrow)
	case 3:
		randomized := randomize(20, 40)
		casesAttackToDragon(randomized, harmOfFrstn)
	default:
		fmt.Println(setGetMainMenuText().incorrectInp)
	}
}

func casesAttackToDragon(randomized int, harm int) {
	hpHero = hpHero - randomized
	hpDragon = hpDragon - harm
	fmt.Println(setGetGameDataText().harmHeroToDragon, harm)
	if randomized == 0 {
		fmt.Println(setGetGameDataText().dragonMiss)
	} else {
		fmt.Println(setGetGameDataText().harmDragonToHero, randomized)
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
