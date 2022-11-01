package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kardianos/osext"
)

func main() {
	exeFold, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatalln(err)
	}
	os.Chdir(exeFold)

	foldName := strings.Split(exeFold, "/")[len(strings.Split(exeFold, "/"))-1]
	fmt.Println(foldName + "/")

	exePath, err := osext.Executable()
	if err != nil {
		log.Fatalln()
	}
	exeName := strings.Split(exePath, "/")[len(strings.Split(exePath, "/"))-1]

	ReadDir(".", exeName, 0)
}

func ReadDir(dir string, exeName string, depth int) {
	var depthString string
	for i := 0; depth > i; i++ {
		depthString += "│  "
	}
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalln(err)
	}

	for _, file := range files {
		if file.IsDir() && !strings.HasPrefix(file.Name(), ".") {
			fmt.Println(depthString + "├─ " + file.Name() + "/")
			ReadDir(dir+"/"+file.Name(), exeName, depth+1)
		} else if depth == 0 && file.Name() == exeName {
			//skip this because we dont list this executable
		} else {
			fmt.Println(depthString + "├─ " + file.Name())
		}
	}
}
