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
	point1       string
	point2       string
	point3       string
	langRU       string
	langEN       string
	loading      string
	bye          string
	incorrectInp string
}

var menuText = Menu{ //Тексты главного меню
	point1:       "1. Начать новую игру",
	point2:       "2. Выход",
	loading:      "Загружается...",
	bye:          "До скорой встречи!",
	incorrectInp: "Неверный выбор, повторите снова",
}

type Game struct { //Игровые тексты
	hp                string
	dragonMiss        string
	harmHeroToDragon  string
	harmDragonToHero  string
	hero              string
	dragon            string
	entHeroName       string
	gameOver          string
	selWeapon         string
	weapon1           string
	weapon2           string
	weapon3           string
	winner            string
	step              string
	standoff          string
	isGameStart       bool
	isGameEnd         bool
	inputMainMenuItem string
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
	weapon3:          "3. Огненный камень",
	winner:           "Победил",
	step:             "Ход #",
	standoff:         "Победила дружба :)",
	isGameStart:      false,
	isGameEnd:        false,
}

type Hero struct {
	hp     int
	damage int
	weapon int
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

func showMainMenu() {
	fmt.Println(menuText.point1)
	fmt.Println(menuText.point2)
}

func selectMainMenuItem() {
	someThing := bufio.NewScanner(os.Stdin)
	someThing.Scan()
	gameText.inputMainMenuItem = strings.TrimSpace(someThing.Text()) //Убирает пробелы в начале и в конце
	switch gameText.inputMainMenuItem {
	case "1": //Начать новую игру
		fmt.Println(menuText.loading)
		gameText.isGameStart = true
	case "2": //Выход
		fmt.Println(menuText.bye)
		os.Exit(0)
	default:
		fmt.Println(menuText.incorrectInp)
		selectMainMenuItem()
	}
}

func gameStart() {
	fmt.Println(gameText.entHeroName)
	inputHeroName()
	step := 1
	for {
		if !gameText.isGameEnd {
			showGameResult()

			fmt.Println(gameText.step, step) //Shows step
			step++

			showWeaponHero()
			selectWeapon()
			attackToDragon()
			CheckCurrentHp(heroData.hp, dragonData.hp)
		} else if gameText.isGameEnd {
			fmt.Println(gameText.gameOver) //Shows Game Over
			showGameResult()

			ShowWinner(heroData.hp, dragonData.hp)
			break
		}
	}
}

func inputHeroName() {
	someThing := bufio.NewScanner(os.Stdin)
	someThing.Scan()
	if someThing.Text() == `` { //для обработки пустой строки
		heroData.name = FetchHeroName()
	} else {
		heroData.name = strings.TrimSpace(someThing.Text()) //Убирает пробелы в начале и в конце
	}
	fmt.Println("")
}

func FetchHeroName() string {
	data := map[string]string{}
	resp, err := http.Get("https://uinames.com/api/?amount=1&gender=male&region=russia")
	if err != nil {
		fmt.Println(err)
		return "Error fetching name from API"
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &data)
	return data["name"]
}

func showGameResult() {
	fmt.Println(gameText.hero, heroData.name,
		"\t\t\t", gameText.dragon, dragonData.name)
	showCurrentHP()
}

func showCurrentHP() {
	fmt.Println(heroData.hp, gameText.hp, "\t\t\t\t", dragonData.hp)
}

func CheckCurrentHp(hpHero, hpDragon int) bool {
	if hpHero|hpDragon < 1 {
		gameText.isGameEnd = true
	}
	return gameText.isGameEnd
}

func ShowWinner(hpHero, hpDragon int) int {
	if hpHero > hpDragon {
		fmt.Println("")
		fmt.Println(gameText.winner, gameText.hero, heroData.name)
		return hpDragon
	} else if hpDragon > hpHero {
		fmt.Println(gameText.winner, gameText.dragon)
		return hpDragon
	} else if hpDragon == hpHero {
		fmt.Println(gameText.standoff)
		return hpDragon
	}
	return 0
}

func showWeaponHero() {
	fmt.Println("")
	fmt.Println("")
	fmt.Println(gameText.selWeapon)
	weaponHero := [3]string{
		gameText.weapon1,
		gameText.weapon2,
		gameText.weapon3,
	}
	for i := 0; i < len(weaponHero); i++ {
		fmt.Println(weaponHero[i])
	}
}

func selectWeapon() {
	fmt.Scan(&heroData.weapon)
	fmt.Println("")
	fmt.Println("")
}

func randomize(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func attackToDragon() {
	switch heroData.weapon {
	case 1:
		randomized := randomize(0, 20)
		heroData.damage = 10
		CasesAttackToDragon(heroData.damage)
		CasesAttackToHero(randomized)
	case 2:
		randomized := randomize(10, 30)
		heroData.damage = 15
		CasesAttackToDragon(heroData.damage)
		CasesAttackToHero(randomized)
	case 3:
		randomized := randomize(20, 40)
		heroData.damage = 30
		CasesAttackToDragon(heroData.damage)
		CasesAttackToHero(randomized)
	default:
		fmt.Println(menuText.incorrectInp)
	}
}

func CasesAttackToDragon(damage int) int {
	dragonData.hp -= damage
	fmt.Println(gameText.harmHeroToDragon, damage)
	return dragonData.hp
}

func CasesAttackToHero(randomized int) int {
	heroData.hp -= randomized
	if randomized == 0 {
		fmt.Println(gameText.dragonMiss)
	} else {
		fmt.Println(gameText.harmDragonToHero, randomized)
	}
	fmt.Println("")
	fmt.Println("")
	return heroData.hp
}

func main() {
	for {
		showMainMenu()
		selectMainMenuItem()
		if gameText.isGameStart {
			break
		}
	}
	gameStart()
}
