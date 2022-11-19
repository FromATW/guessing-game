package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNum := rand.Intn(maxNum)
	fmt.Println("The secretNum is", secretNum)

	fmt.Println("Please input your number")
	for {
		var guess int
		_, err := fmt.Scanf("%d\n", &guess)
		if err != nil {
			fmt.Println("Invalid Num")
			return
		}
		fmt.Println("Your guess is", guess)
		if guess > secretNum {
			fmt.Println("Your guess is bigger than the secret number. Please try again")
		} else if guess < secretNum {
			fmt.Println("Your guess is smaller than the secret number. Please try again")
		} else {
			fmt.Println("Correct, you Legend")
			break
		}
	}
}
