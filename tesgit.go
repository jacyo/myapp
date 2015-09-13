package main

import "os"

func FileExists(file string) bool {
	if _, err := os.Stat(file); err != nil {
		return false
	}
	return true
}

func main() {
    println(FileExists("file_exists.go"))
}