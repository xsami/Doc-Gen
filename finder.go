// The purpose of this function is to find all the instance of the variable X in the code
package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

type Instance struct {
	FilePath string "the path"
	Line     int    "the line"
}

// This function recives the string that you're looking for, the path
// that will be used as root to look into the project recursivly and
// an array with the exceptions, that must be commonly be comments.
// The function will return an array with an struct that defines the
// file's path and the line number of theses.
func find(x string, directoryPath string, expect []string) ([]Instance, error) {
	fmt.Println("Looking for", x, "instances")
	// TODO: Implement this logic
	var result = []Instance{}

	recursiveSearch(directoryPath)

	return result, nil
}

func recursiveSearch(directory string) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		fmt.Println(f.Name())
	}
}

func buildMsg(x string) (string, error) {
	return "", nil
}

func readFile(file string) ([]byte, error) {
	content, err := ioutil.ReadFile(file)
	return content, err
}
