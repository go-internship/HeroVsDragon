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

type Menu struct { //Тексты главного меню
	point1            string
	point2            string
	point3            string
	langRU            string
	langEN            string
	bye               string
	incorrectInp      string
	inputMainMenuItem string
}

var menuText = Menu{ //Тексты главного меню
	point1:       "1. Начать новую игру",
	point2:       "2. Выход",
	bye:          "До скорой встречи!",
	incorrectInp: "Неверный выбор, повторите снова",
}

type Game struct { //Игровые тексты
	hp               string
	dragonMiss       string
	harmHeroToDragon string
	harmDragonToHero string
	hero             string
	dragon           string
	entHeroName      string
	gameOver         string
	selWeapon        string
	weapon1          string
	weapon2          string
	weapon3          string
	winner           string
	step             string
	standoff         string
	isGameStart      bool
	isGameEnd        bool
	whoIsWinner      int
}

var gameText = Game{ //Игровые тексты
	hp:               "hp",
	dragonMiss:       "Дракон промахнулся и не нанёс вам урона :)",
	harmHeroToDragon: "Вы нанесли урон Дракону",
	harmDragonToHero: "Дракон нанёс Вам урон",
	hero:             "Герой",
	dragon:           "Дракон",
	entHeroName:      "Введите имя Героя: (нажмите Enter для случайного)",
	gameOver:         "Игра завершилась",
	selWeapon:        "Выберите оружие:",
	weapon1:          "1. Меч",
	weapon2:          "2. Стрела",
	weapon3:          "3. Огненный камень\n",
	winner:           "Победил",
	step:             "Ход #",
	standoff:         "Победила дружба :)",
	isGameStart:      false,
	isGameEnd:        false,
}

type Hero struct {
	hp     int
	damage int
	weapon string
	name   string
}

var heroData = Hero{
	hp: 100,
}

type Dragon struct {
	hp   int
	name string
}

var dragonData = Dragon{
	hp:   100,
	name: "Драконыч",
}

func ScanInput(input string) string { //CAN'T TEST
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	input = strings.TrimSpace(scan.Text())
	return input
}

func SelectMainMenuItem(inputData string) bool { //TESTED
	if inputData == "" { //For test
		inputData = (ScanInput(menuText.inputMainMenuItem))
	}
	switch inputData {
	case "1": //Начать новую игру
		gameText.isGameStart = true
		return gameText.isGameStart
	case "2": //Выход
		fmt.Println(menuText.bye)
		gameText.isGameStart = false
		os.Exit(0)
		return gameText.isGameStart
	default:
		fmt.Println(menuText.incorrectInp)
		gameText.isGameStart = false
		return gameText.isGameStart
	}
}

func InputHeroName(inputData string) string { //TESTED
	if inputData == "" {
		inputData = ScanInput(heroData.name)
	}
	if inputData == `` { //для обработки пустой строки
		heroData.name = FetchHeroName()
	} else {
		heroData.name = strings.TrimSpace(inputData) //Убирает пробелы в начале и в конце
	}
	fmt.Print("\n")
	return heroData.name
}

func FetchHeroName() string { //TESTED
	data := map[string]string{}
	resp, err := http.Get("https://uinames.com/api/?amount=1&gender=male&region=russia")
	if err != nil {
		fmt.Println(err)
		return "Error fetching name"
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &data)
	return data["name"]
}

func ShowGameResult() { //DRY //CAN'T TEST
	fmt.Println(gameText.hero, heroData.name,
		"\t\t\t", gameText.dragon, dragonData.name)
	fmt.Println(heroData.hp, gameText.hp, "\t\t\t\t", dragonData.hp)
}

func CheckCurrentHp(hpHero, hpDragon int) bool { //TESTED
	if hpHero|hpDragon < 1 {
		gameText.isGameEnd = true
	}
	return gameText.isGameEnd
}

func CheckWinner(hpHero, hpDragon int) int { //TESTED
	if hpHero > hpDragon {
		gameText.whoIsWinner = 1
		return gameText.whoIsWinner
	} else if hpDragon > hpHero {
		gameText.whoIsWinner = 2
		return gameText.whoIsWinner
	} else if hpDragon == hpHero {
		gameText.whoIsWinner = 3
		return gameText.whoIsWinner
	}
	return 0
}

func Randomize(min, max int) int { //CAN'T TEST
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func AttackHeroAndDragon(inputData string) bool { //TESTED
	if inputData == "" {
		inputData = (ScanInput(heroData.weapon))
	}
	switch inputData {
	case "1": //Меч
		randomized := Randomize(0, 20)
		heroData.damage = 10
		CasesAttackToDragon(heroData.damage)
		CasesAttackToHero(randomized)
		return true
	case "2": //Стрела
		randomized := Randomize(10, 30)
		heroData.damage = 15
		CasesAttackToDragon(heroData.damage)
		CasesAttackToHero(randomized)
		return true
	case "3": //Огненный камень
		randomized := Randomize(20, 40)
		heroData.damage = 30
		CasesAttackToDragon(heroData.damage)
		CasesAttackToHero(randomized)
		return false
	default:
		fmt.Println(menuText.incorrectInp)
		return false
	}
}

func CasesAttackToDragon(damage int) int { //TESTED
	dragonData.hp -= damage
	fmt.Println(gameText.harmHeroToDragon, damage)
	return dragonData.hp
}

func CasesAttackToHero(randomized int) int { //TESTED
	heroData.hp -= randomized
	if randomized == 0 {
		fmt.Println(gameText.dragonMiss)
	} else {
		fmt.Println(gameText.harmDragonToHero, randomized)
	}
	fmt.Print("\n\n")
	return heroData.hp
}

func GameStart() {
	fmt.Println(gameText.entHeroName)
	InputHeroName(heroData.name)
	step := 1
	for {
		if !gameText.isGameEnd {
			ShowGameResult()

			fmt.Print(gameText.step, step, "\n\n") //Shows step
			step++

			fmt.Println(gameText.selWeapon)
			fmt.Println(gameText.weapon1)
			fmt.Println(gameText.weapon2)
			fmt.Print(gameText.weapon3)

			fmt.Scan(&heroData.weapon) //Input weapon of hero

			AttackHeroAndDragon(heroData.weapon)
			CheckCurrentHp(heroData.hp, dragonData.hp)
		} else if gameText.isGameEnd {
			fmt.Println(gameText.gameOver) //Shows Game Over
			ShowGameResult()
			fmt.Print("\n")
			switch CheckWinner(heroData.hp, dragonData.hp) {
			case 1:
				fmt.Println(gameText.winner, gameText.hero, heroData.name)
			case 2:
				fmt.Println(gameText.winner, gameText.dragon)
			case 3:
				fmt.Println(gameText.standoff)
			}
			break
		}
	}
}

func main() { //CAN'T TEST
	for {
		fmt.Println(menuText.point1) //Shows main menu
		fmt.Println(menuText.point2)

		SelectMainMenuItem(menuText.inputMainMenuItem)
		if gameText.isGameStart {
			break
		}
	}
	GameStart()
}
