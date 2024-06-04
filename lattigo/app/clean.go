package main

import (
	"log"
	"os"
)

// Cleans the temporary files, which includes the keys, encrypted data and results.
func main() {
	cleanFolder("temps/")
}

func cleanFolder(folderPath string) {
	folder, _ := os.Open(folderPath)
	files, _ := folder.Readdir(0)
	for i := range files {
		if files[i].Name() != "donotremove.txt" {
			if err := os.Remove(folderPath + files[i].Name()); err != nil {
				log.Println(err)
			}
		}
	}
}
