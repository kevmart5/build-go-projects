package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CharacterType struct {
	characterType string
}

type Race struct {
	race string
}

type Character struct {
	name          string
	characterType CharacterType
	race          Race
}

func createNewCharacter(c *Character) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("********** Character menu **********")
	fmt.Println("")

	n := 0
	for n < 3 {
		if n == 0 {
			fmt.Println("Name of the character")
			fmt.Print("-> ")
			text, _ := reader.ReadString('\n')
			// convert CRLF to LF
			text = strings.Replace(text, "\n", "", -1)
			n++
		}

		if n == 1 {
			fmt.Println("Type of the character")
			fmt.Print("-> ")
			text, _ := reader.ReadString('\n')
			// convert CRLF to LF
			text = strings.Replace(text, "\n", "", -1)
			n++
		}

		if n == 2 {
			fmt.Println("Race of the character")
			fmt.Print("-> ")
			text, _ := reader.ReadString('\n')
			// convert CRLF to LF
			text = strings.Replace(text, "\n", "", -1)
			n++
		}
	}
	displayMainMenu()
}

func displayMainMenu() {
	reader := bufio.NewReader(os.Stdin)
	newCharacter := makeCharacter()
	fmt.Println("********** Menu **********")
	fmt.Println("1. Option 1")
	fmt.Println("2. Option 2")
	fmt.Println("0. Exit program")
	fmt.Println("---------------------")
	fmt.Println("")
	fmt.Println("%+v", newCharacter)
	fmt.Println("")
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("1", text) == 0 {
			fmt.Println("first option")
			createNewCharacter(*newCharacter)
		}

		if strings.Compare("2", text) == 0 {
			fmt.Println("Second option")
		}

		if strings.Compare("0", text) == 0 {
			fmt.Println("---------------------")
			fmt.Println("Bye Bye!")
			fmt.Println("---------------------")
			os.Exit(1)
		}

	}
}

func makeCharacter() Character {
	characterType := CharacterType{characterType: "Knight"}
	race := Race{race: "human"}
	defaultCharacter := Character{
		name:          "No name",
		characterType: characterType,
		race:          race,
	}

	return defaultCharacter
}

func main() {
	displayMainMenu()
}
