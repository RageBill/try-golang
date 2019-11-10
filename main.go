package main

import (
	"fmt"
	"time"
)

func greetUser(nameChannel chan string) {
	fmt.Println("Greetings, user!")
	fmt.Println("The time now is", time.Now())
	var username string
	fmt.Println("What is your name?")
	fmt.Scan(&username)     // wait for user input
	nameChannel <- username // pipe username to channel output
}

func getFood(foodChannel chan []string) {
	foodChannel <- []string{"Apple", "Pepsi", "French Fries"}
}

func main() {
	// create channels
	usernameChannel := make(chan string)
	myFoodChannel := make(chan []string)

	// create goroutines
	go greetUser(usernameChannel)
	go getFood(myFoodChannel)

	switch username := <-usernameChannel; username {
	case "RageBill":
		fmt.Println("I see, so you are the developer!")
	default:
		fmt.Println("Nice to have you here, stranger!")
	}

	for index, food := range <-myFoodChannel {
		fmt.Println("Food I like, no.", index, ":", food)
	}
}
