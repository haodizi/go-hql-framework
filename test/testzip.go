package main

import (
	"awesomeProject/utils"
	"fmt"
)

func main() {
	sourceFile := []string{"logs/a1.txt", "logs/a2.txt"}
	error := utils.MakeZip(sourceFile, "logs/a12.zip")
	if error != nil {
		fmt.Println("Make zip file failed,", error)
	}
}
