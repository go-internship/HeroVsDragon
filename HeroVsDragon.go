package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

//Если не нужна, то убрать эту ф-ию нафиг
/*func pauseEachLevel() {
	//fmt.Println("Нажмите Enter для продолжения")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}*/

func sleepEachStep() {
	//желательно применить метод DRY
	fmt.Print(".")
	duration := time.Duration(1) * time.Second
	time.Sleep(duration)
	fmt.Print(".")
	time.Sleep(duration)
	fmt.Print(".\n\n")
}

func randomIndexWeaponOfDragon() int {
	rand.Seed(time.Now().UTC().UnixNano())
	e := rand.Intn(5)
	return e
}

func main() {
	weaponOfHero := [10]string{
		"мечом",
		"огнём",
		"ионной пушкой",
		"базукой",
		"пистолетом",
		"руками",
		"сковородкой",
		"поварёжкой",
		"ноутбуком",
		"камнем",
	}

	weaponOfDragon := [5]string{
		"дышащим огнём",
		"лапами",
		"камнем",
		"прыжками",
		"хвостом",
	}

	harmOfDragon := make(map[int]int)
	harmOfDragon[0] = 30 //дышащим огнём
	harmOfDragon[1] = 20 //ударом лап
	harmOfDragon[2] = 25 //камнями
	harmOfDragon[3] = 10 //прыжками
	harmOfDragon[4] = 15 //ударом хвоста

	attackHero := true
	fmt.Println("\nПривет, герой!\n")

	hpHero := 100 //инициализация жизни Героя
	//Цикл отображения отображает и итерирует Уровни
	for i := 0; ; i++ {
		//Цикл отображает и итерирует Ходы
		for j := 0; ; j++ {
			fmt.Println("Уровень", i+1, "\nХод", j+1)
			fmt.Println("У вас", hpHero, "hp")

			//Надо реализовать эту конструкцию в функции
			if attackHero == true {
				attackHero = false
				//Логика оружия героя
				weapon := 0
				fmt.Println("Выберите оружие для атаки на Дракона:\n")
				for o := 0; o < 10; o++ {
					fmt.Println(o+1, weaponOfHero[o])
				}
				fmt.Fscan(os.Stdin, &weapon) //А вот тут надо сделать обработчик ошибок, т.к. будет вылетать out of range, если юзер напишет < 10
				//Вывод оружия героя
				fmt.Println("\nВы атакуете Дракона", weaponOfHero[weapon-1], "\n")

				sleepEachStep()
			} else if attackHero == false {
				attackHero = true
				randLocal := randomIndexWeaponOfDragon() //чтобы ф-ия вызывалась один раз иначе при каждом вызове будет другое значение
				fmt.Println("\nДракон атакует Вас", weaponOfDragon[randLocal])
				hpHero = hpHero - harmOfDragon[randLocal]
				sleepEachStep()
				fmt.Println("Дракон нанёс вам урон", harmOfDragon[randLocal], "hp")
				sleepEachStep()
			}
		}
	}
}

//TODO: [РЕШЕНО]1. Решить вопрос с функц. рандомизации выбора оружия. 2. Отобразить урон, hp дракона и героя. После каждого хода.
