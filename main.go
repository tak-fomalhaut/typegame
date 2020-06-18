package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	satisfaction := 0

	for n := 1; n <= 4; n++ {
		order(n, menu(n), &satisfaction)
	}
	if satisfaction == 4 {
		fmt.Println("Super delicious!")
	} else if satisfaction == 3 {
		fmt.Println("Yummy!")
	} else if satisfaction == 2 {
		fmt.Println("so-so")
	} else if satisfaction == 1 {
		fmt.Println("Unsatisfactory..")
	} else {
		fmt.Println("Are you hungry?")
	}
}

func order(num int, ques string, satis *int) {
	var input string
	fmt.Printf("[%d品目] 次と同じ値を入力してください : %s\n", num, ques)
	fmt.Scan(&input)
	if input == ques {
		fmt.Println("Get！")
		*satis++
	} else {
		fmt.Println("Miss！")
	}
}

func menu(m int) string {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(4)
	switch m {
	case 1:
		soup := []string{"cornpotage", "minestrone", "onionsoup", "clamchowder"}
		return soup[i]
	case 2:
		salad := []string{"caesarsalad", "cobbsalad", "caprese", "beanssalad"}
		return salad[i]
	case 3:
		mainDish := []string{"roastedlobster", "beefsteak", "lambchops", "swordfishmeuniere"}
		return mainDish[i]
	default:
		dessert := []string{"minticecream", "affogato", "freshmango", "strawberrysorbet"}
		return dessert[i]
	}
}
