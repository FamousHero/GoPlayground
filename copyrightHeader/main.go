package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// util script to add a header and title to a new go file.

type FileCreator struct {
	fileName    string
	creatorName string
}

func createFileCreator(filename, creator string) FileCreator {
	return FileCreator{
		fileName:    filename,
		creatorName: creator,
	}
}

func (fc FileCreator) createFile() string {
	date := time.Now().Year()
	copyrightString := new(strings.Builder)
	fmt.Fprintf(copyrightString, "// Copyright %d %s. All rights reserved. ", date, fc.creatorName)

	return copyrightString.String()
}

func main() {
	args := os.Args[1:]
	fc := FileCreator{
		fileName:    "test.go",
		creatorName: strings.Join([]string{"Daniel", "Franco"}, " "),
	}
	fmt.Println(fc.createFile())

	license := flag.String("license", "", "states how the source code is governed")
	flag.Parse()
	// No license provided, check if one in current dir
	if *license == "" {
		dir, err := os.Getwd() //Get the directory you are creating the file in
		if err != nil {
			fmt.Println("Error: ", err)
		}
		files, err := os.ReadDir(dir)
		if err != nil {
			log.Fatal(err)
		}
		// YOU ARE HERE. Ending for today
	}
	// Check if there is a license in the directory

	if len(args) > 3 { // only argument to program should be the filename
		panic("Too many arguments. Max 3: filename, first name, last name")

	}
	switch len(args) {
	case 3:
		break
	case 2:
		break
	case 1:
		break
	}
}
