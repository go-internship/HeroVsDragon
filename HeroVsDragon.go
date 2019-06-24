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
			CheckCurrentHp(hpHero, hpDragon)
		} else if isGameEnd {
			gameEnd()
			ShowWinner(hpHero, hpDragon)
			break
		}
	}
}

func inputHeroName() {
	someThing := bufio.NewScanner(os.Stdin)
	someThing.Scan()
	if someThing.Text() == `` { //для обработки пустой строки
		data := map[string]string{}
		resp, err := http.Get("https://uinames.com/api/?amount=1&gender=male&region=russia")
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

func CheckCurrentHp(hpHeroo, hpDragoon int) bool {
	if hpHeroo|hpDragoon < 1 {
		isGameEnd = true
	}
	return isGameEnd
}

func showStep(step int) {
	fmt.Println(setGetGameDataText().step, step+1)
}

func gameEnd() {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println(setGetGameDataText().gameOver)
	showGameResult()
}

func ShowWinner(hpHeroo, hpDragoon int) int {
	if hpHeroo > hpDragoon {
		fmt.Println("")
		fmt.Println(setGetGameDataText().winner, setGetGameDataText().hero, heroName)
		return hpDragoon
	} else if hpDragoon > hpHeroo {
		fmt.Println(setGetGameDataText().winner, setGetGameDataText().dragon)
		return hpDragoon
	} else if hpDragoon == hpHeroo {
		fmt.Println(setGetGameDataText().standoff)
		return hpDragoon
	}
	return 0
}

func showWeaponHero() {
	fmt.Println("")
	fmt.Println("")
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
	fmt.Println("")
	fmt.Println("")
}

func randomize(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func attackToDragon() {
	switch weaponHero {
	case 1:
		randomized := randomize(0, 20)
		CasesAttackToDragon(harmOfSword)
		CasesAttackToHero(randomized)
	case 2:
		randomized := randomize(10, 30)
		CasesAttackToDragon(harmOfArrow)
		CasesAttackToHero(randomized)
	case 3:
		randomized := randomize(20, 40)
		CasesAttackToDragon(harmOfFrstn)
		CasesAttackToHero(randomized)
	default:
		fmt.Println(setGetMainMenuText().incorrectInp)
	}
}

func CasesAttackToDragon(harm int) int {
	hpDragon = hpDragon - harm
	fmt.Println(setGetGameDataText().harmHeroToDragon, harm)
	return hpDragon
}

func CasesAttackToHero(randomized int) int {
	hpHero = hpHero - randomized
	if randomized == 0 {
		fmt.Println(setGetGameDataText().dragonMiss)
	} else {
		fmt.Println(setGetGameDataText().harmDragonToHero, randomized)
	}
	fmt.Println("")
	fmt.Println("")
	return hpHero
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
