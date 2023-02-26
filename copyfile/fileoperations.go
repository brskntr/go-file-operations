package copyfile

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func CopyFile(source string, destination string) {
	input, err := ioutil.ReadFile(source)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(destination, input, 0644)
	if err != nil {
		fmt.Println("Error creating", destination)
		fmt.Println(err)
		return
	}
}

func ListFile(directory string) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
	}
}

func MoveFile(destination string, source string) {
	err := os.Rename(source, destination)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteFile(fileDir string) {
	err := os.Remove(fileDir)
	if err != nil {
		log.Fatal(err)
	}
}
