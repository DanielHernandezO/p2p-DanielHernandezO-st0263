package shell

import (
	"fmt"
	"strings"

	"github.com/DanielHernandezO/p2p/client/internal/business/constant"
)

func mapActions(dependencies *Dependencies) {
	var joined bool
	for {
		fmt.Println("Choose an option:")
		if !joined {
			fmt.Println("1. Join")
		}
		if joined {
			fmt.Println("2. Post files")
			fmt.Println("3. Search")
		}
		fmt.Println("0. exit")

		var opcion int
		fmt.Scanln(&opcion)

		switch opcion {
		case 0:
			fmt.Println("going out...")
			return
		case 1:
			answer := dependencies.executeJoinHandler.ExecuteJoin()
			fmt.Println(answer)
			if answer == constant.Connected {
				joined = true
			}
		case 2:
			if joined {
				answer := dependencies.fileHandler.PostFiles()
				fmt.Println(answer)
			} else {
				fmt.Println("You should join first")
			}
		case 3:
			if joined {
				var fileName string
				fmt.Scanln(&fileName)
				answer := dependencies.executeJoinHandler.Search(fileName)
				if strings.ToLower(answer) == constant.Found {
					fmt.Println("Downloading")
				} else {
					fmt.Println(answer)
				}
			} else {
				fmt.Println("You should join first")
			}
		default:
			fmt.Println("Invalid Option")
		}
	}
}
