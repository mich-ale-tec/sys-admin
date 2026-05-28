package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CommandFilter(files []os.DirEntry, pathDir string) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Match :")
	scanner.Scan()
	match := scanner.Text()

	fmt.Printf("Buscando archivos con %s\n", match)
	filesFilter := FilterFilesByMatch(files, match)

	fmt.Print("Borrar? (Y/n): ")
	scanner.Scan()
	resRemoveFiles := scanner.Text()
	ValidConfirmRemove(resRemoveFiles)

	RemoveFiles(filesFilter, pathDir)
}

func ValidConfirmRemove(response string) {
	if response == "n" {
		os.Exit(0)
	}
}

func RemoveFiles(files []os.DirEntry, pathDir string) {
	for _, file := range files {
		path := filepath.Join(pathDir, file.Name())
		err := os.Remove(path)
		if err != nil {
			fmt.Printf("Error borrando %s\n", file.Name())
		}
	}
	fmt.Printf("Borrado (%d) archivo(s)", len(files))
}

func FilterFilesByMatch(files []os.DirEntry, match string) []os.DirEntry {
	var filesFilter []os.DirEntry
	for i, entry := range files {
		if strings.Contains(entry.Name(), match) {
			filesFilter = append(filesFilter, entry)
			fmt.Printf("%d: %s\n", i+1, entry.Name())
		}
	}
	return filesFilter
}
