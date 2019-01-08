package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("ERROR: You need to specify the path of the project")
		return
	}

	project := os.Args[1]

	fmt.Print("Processing the code...\n")

	fmt.Println(project)
	// TODO: Validate the path and etc
}
