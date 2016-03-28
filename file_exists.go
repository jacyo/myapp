package main

import "os"

// 判断文件是否存在123
func FileExists(file string) bool {
	if _, err := os.Stat(file); err != nil {
		return false
	}
	return true
}


func main() {
    println(FileExists("file_exists.go"))
}