/* Author - sula7 (@sulafpv)
Please read a manual on GitHub on how to run this application
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
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
	Dragon
	hp     int
	name   string
	damage int
	weapon string
}

type Dragon struct {
	hp     int
	name   string
	damage int
	weapon string
}

func ScanInput() string { //CAN'T TEST
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	input := strings.TrimSpace(scan.Text())
	return input
}

func SelectMainMenuItem() bool {
	input := ScanInput()
	switch input {
	case "1":
		return true
	case "2":
		fmt.Println(menuText.bye)
		os.Exit(0)
		return false
	default:
		fmt.Println(menuText.incorrectInp)
		SelectMainMenuItem()
	}
	return false
}

func (h *Hero) SetHeroName() {
	input := ScanInput()
	if input == `` {
		h.name = FetchHeroName()
	} else if input != `` {
		h.name = input
	}
}

func (h *Hero) SetHeroHP() {
	input := ScanInput()
	switch input {
	case "1":
		h.Dragon.damage = Randomize(0, 20)
		h.hp -= h.Dragon.damage
		h.damage = 10
	case "2":
		h.Dragon.damage = Randomize(10, 30)
		h.hp -= h.Dragon.damage
		h.damage = 15
	case "3":
		h.Dragon.damage = Randomize(20, 40)
		h.hp -= h.Dragon.damage
		h.damage = 30
	default:
		fmt.Printf(menuText.incorrectInp)
	}
	fmt.Print("\n")
}

func (h *Hero) SetDragonHP() {
	h.Dragon.hp -= h.damage
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

func ShowGameResult(hName string, hHP int, dName string, dHP int) { //DRY //DON'T NEED A TEST
	s := fmt.Sprintf("Герой %s \t\t\t Дракон %s \n", hName, dName)
	io.WriteString(os.Stdout, s)
	fmt.Println(hHP, gameText.hp, "\t\t\t\t", dHP, gameText.hp)
}

func CheckGameEnd(hpHero, hpDragon int) bool {
	if hpHero|hpDragon < 1 {
		gameText.isGameEnd = true
	}
	return gameText.isGameEnd
}

func CheckWinner(hHP, dHP int) int {
	if dHP < 0 {
		gameText.whoIsWinner = 1
		return gameText.whoIsWinner
	} else if hHP < 0 {
		gameText.whoIsWinner = 2
		return gameText.whoIsWinner
	} else if dHP < 0 || hHP < 0 {
		gameText.whoIsWinner = 3
		return gameText.whoIsWinner
	}
	return 0
}

func Randomize(min, max int) int { //CAN'T TEST
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func GameStart() bool {
	hData := Hero{
		Dragon: Dragon{
			hp:     100,
			name:   "Драконыч",
			damage: 0,
		},
		hp:     100,
		damage: 0,
		name:   "default",
	}

	fmt.Println(gameText.entHeroName)
	hData.SetHeroName()
	step := 1
	for {
		if !gameText.isGameEnd {
			ShowGameResult(hData.name, hData.hp, hData.Dragon.name, hData.Dragon.hp)

			fmt.Print(gameText.step, step, "\n\n")
			step++

			fmt.Println(gameText.selWeapon)
			fmt.Println(gameText.weapon1)
			fmt.Println(gameText.weapon2)
			fmt.Print(gameText.weapon3)

			hData.SetHeroHP()
			hData.SetDragonHP()
			fmt.Println(gameText.harmDragonToHero, hData.Dragon.damage)
			fmt.Println(gameText.harmHeroToDragon, hData.damage)

			CheckGameEnd(hData.hp, hData.Dragon.hp)
		} else if gameText.isGameEnd {
			break
		}
	}
	fmt.Println(gameText.gameOver) //Shows Game Over
	fmt.Print("\n")
	ShowGameResult(hData.name, hData.hp, hData.Dragon.name, hData.Dragon.hp)
	fmt.Print("\n")
	switch CheckWinner(hData.hp, hData.Dragon.hp) {
	case 1:
		fmt.Println(gameText.winner, gameText.hero, hData.name)
	case 2:
		fmt.Println(gameText.winner, gameText.dragon, hData.Dragon.name)
	case 3:
		fmt.Println(gameText.standoff)
	}
	return true
}

func main() { //CAN'T TEST
	fmt.Println(menuText.point1)
	fmt.Println(menuText.point2)

	gameText.isGameStart = SelectMainMenuItem()
	if gameText.isGameStart {
		GameStart()
	}
}
