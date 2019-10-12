package handlers

import (
	"fmt"
	"os"
	"strconv"

	"github.com/boltdb/bolt"
)

//Console manages client UI
func Console(db *bolt.DB) {
	var option int
	for {
		fmt.Println("")
		fmt.Println("Choose Option: ")
		fmt.Println("1. GetList")
		fmt.Println("2. Register Files")
		fmt.Println("3. Download")
		fmt.Println("999. Quit")

		fmt.Scan(&option)
		switch option {
		case 1:
			list := GetList()
			if list != nil {
				for _, l := range list {
					fmt.Println(l.Name)
				}
			} else {
				fmt.Println("Failed ... No files?")
			}
		case 2:
			fmt.Println("Enter absolute file path to share...")
			var path, name string
			fmt.Scanln(&path)
			fmt.Println("Enter the name with which to share file: ")
			fmt.Scanln(&name)
			info, err := os.Stat(path)
			if os.IsNotExist(err) {
				fmt.Println("Invalid PATH")
			} else {
				fmt.Println("Sharing File: " + path)
				fmt.Println("Share Name: " + name)
				fmt.Println("File Size: " + strconv.FormatInt(info.Size(), 10))
				if FileRegister(name, strconv.FormatInt(info.Size(), 10)) {
					FileMap(name, path, db)
				}
			}
		case 3:
			fmt.Println("Enter file Name to download: ")
			var name string
			fmt.Scanln(&name)
			list, _ := GetLocation(name)
			if list == nil {
				fmt.Printf("%s doesnt exist ...", name)
				continue
			}
			fmt.Println("Enter absolute file path to download...")
			var path string
			fmt.Scanln(&path)
			info, err := os.Stat(path)
			if os.IsNotExist(err) || !info.IsDir() {
				fmt.Println("Invalid PATH")
				continue
			} else {
				FileMap(name, path+"/"+name, db)
				Download(list, name, path)
			}
		case 999:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid option ... try again")
		}
	}
}
