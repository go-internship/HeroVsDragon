package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func pauseEachLevel() {
	//fmt.Println("Нажмите Enter для продолжения")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func main() {
	const mech = "мечом"
	const fire = "огнём"
	const sablya = "ионной пушкой"

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

	weaponOfDragon := [10]string{
		"дышащим огнём",
		"лапами",
		"камнем",
		"прыжками",
		"хвостом",
		"гопниками с района",
		"певицей без голоса",
		"гиперкубической призмой",
		"сферическими идиотами",
		"хомячками",
	}
	attackHero := true

	fmt.Println("\nПривет, герой!\n")

	//Цикл отображения отображает и итерирует Уровни
	for i := 0; i < 10; i++ {
		//Цикл отображает и итерирует Ходы
		for j := 0; j < 10; j++ {
			fmt.Println("Уровень", i+1, "\nХод", j+1)

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
			} else if attackHero == false {
				attackHero = true
				rand.Seed(time.Now().UTC().UnixNano())
				fmt.Println("\nДракон атакует Вас", weaponOfDragon[rand.Intn(10)])
			}
			pauseEachLevel()
		}

		//fmt.Println(attackHero)
		/*fmt.Println(isDragonAttacks(false))
		fmt.Println(isYouAttack(true))*/
	}
}

//TODO: [РЕШЕНО]1. Решить вопрос с функц. рандомизации выбора оружия. 2. Отобразить урон, hp дракона и героя. После каждого хода.
