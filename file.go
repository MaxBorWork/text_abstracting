package main

import (
	"io/ioutil"
	"log"
	"os"
)

func getFilesFromDirectory(name string) []os.FileInfo {
	files, err := ioutil.ReadDir(name)
	if err != nil {
		log.Fatal(err)
	}

	return files
}

func readFile(path string) (string, error) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(dat), nil
}

