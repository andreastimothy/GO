package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"git.garena.com/sea-labs-id/batch-02/andreas-timothy/assignment-activity-reporter/socialgraph"
)

func scanLine(message string) string {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print(message)
	for reader.Scan() {
		return reader.Text()
	}
	return ""
}

func main() {
	system := socialgraph.NewSystem()
	a := true

	for a {
		var input int
		fmt.Println("Activity Reporter")
		fmt.Println("1. Setup")
		fmt.Println("2. Action")
		fmt.Println("3. Display")
		fmt.Println("4. Trending")
		fmt.Println("5. Exit")
		fmt.Print("Enter menu: ")
		fmt.Scan(&input)

		switch input {
		case 1:
			fmt.Println("")
			fmt.Print("Setup social graph: ")
			user1 := ""
			user2 := ""
			_, err := fmt.Scanf("%s follows %s\n", &user1, &user2)
			if err != nil {
				fmt.Println("Invalid Keyword")
				fmt.Scanln()
				break
			}

			fmt.Println("-------------------------------------------------------")
			system.Follow(user1, user2)
			fmt.Println()
		case 2:
			fmt.Println("")
			action := scanLine("Enter user Actions: ")
			fmt.Println("-------------------------------------------------------")

			actions := strings.Split(action, " ")
			if len(actions) == 3 {
				if actions[1] == "uploaded" && actions[2] == "photo" {
					user := actions[0]
					_, err := system.IsUserExist(user)
					if err != nil {
						fmt.Printf("unknown user {%s}", user)
						fmt.Scanln()
						} else {
							system.Upload(user)
						}
				} else {
					fmt.Println("Invalid Keyword")
					fmt.Scanln()
				}
			} else if len(actions) == 4 {
				if actions[1] == "likes" && actions[3] == "photo" {
					user1 := actions[0]
					user2 := actions[2]
					_, err1 := system.IsUserExist(user1)
					_, err2 := system.IsUserExist(user2)
					if err1 != nil {
						fmt.Printf("unknown user {%s}", user1)
						fmt.Scanln()
						} else if err2 != nil {
							fmt.Printf("unknown user {%s}", user2)
							fmt.Scanln()
						} else {
							system.Like(user1, user2)
						}
				} else {
					fmt.Println("Invalid Keyword")
					fmt.Scanln()
				}
			}
			fmt.Println()
		case 3:
			var name string
			fmt.Println("")
			fmt.Print("Display activity of: ")
			fmt.Scan(&name)
			fmt.Println()
			fmt.Println(name + " activities:")
			fmt.Println()

			v, _ := system.IsUserExist(name)
			for _, value := range system.Activities[v] {
				fmt.Println(value)
			}

			fmt.Println()
			fmt.Println("-------------------------------------------------------")
		case 4:
			fmt.Println("")
			fmt.Println("Trending photos: ")
			
			ranking := system.GetRanking()
			for index, value := range ranking {
				fmt.Printf("%v. %s photo got %v likes\n", index + 1, value.Name, value.Likes)
			}
			fmt.Println()
			fmt.Println("-------------------------------------------------------")
		case 5:
			fmt.Println("")
			fmt.Println("Good bye!")
			fmt.Println("-------------------------------------------------------")
			a = false
		default:
			fmt.Println("")
			fmt.Println("Invalid menu!")
			fmt.Println("-------------------------------------------------------")
		}
	}
}