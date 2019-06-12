package main

import (
	"bufio"
	"fmt"
	"os"
)

func pauseGame() {
	//fmt.Println("Нажмите Enter для продолжения")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func isDragonAttacks(attack bool) bool {
	s := attack
	if s == true {
		s = false
	} else if s == false {
		s = true
	}
	return s
}

func isYouAttack(attack bool) bool {
	s := attack
	return s
}

func main() {
	const mech = "мечом"
	const fire = "огнём"
	attackHero := true

	fmt.Println("\nПривет, герой!\n")

	//Цикл отображения отображает и итерирует Уровни
	for i := 0; i < 10; i++ {
		//Цикл отображает и итерирует Ходы
		for j := 0; j < 10; j++ {
			fmt.Println("Уровень", i+1, "\nХод", j+1)

			//Надо реализовать конструкцию в функции
			if attackHero == true {
				attackHero = false
				fmt.Println("Вы атакуете Дракона", mech)
			} else if attackHero == false {
				attackHero = true
				fmt.Println("Дракон атакует Вас", fire)
			}
			pauseGame()
		}

		//fmt.Println(attackHero)
		/*fmt.Println(isDragonAttacks(false))
		fmt.Println(isYouAttack(true))*/
	}
}
