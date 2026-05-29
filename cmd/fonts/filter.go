package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CommandFilter(files []os.DirEntry, pathDir string) error {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Match :")
	scanner.Scan()
	match := scanner.Text()

	fmt.Printf("Buscando archivos con %s\n", match)
	filesFilter := FilterFilesByMatch(files, match)

	fmt.Print("Borrar? (Y/n): ")
	scanner.Scan()
	resRemoveFiles := scanner.Text()
	err := ValidConfirmRemove(resRemoveFiles)
	if err != nil {
		return err
	}
	err = RemoveFiles(filesFilter, pathDir)
	if err != nil {
		return err
	}
	return nil
}

func ValidConfirmRemove(response string) error {
	if response == "n" {
		return fmt.Errorf("Cancelando el comando...")
	}
	return nil
}

func RemoveFiles(files []os.DirEntry, pathDir string) error {
	for _, file := range files {
		path := filepath.Join(pathDir, file.Name())
		err := os.Remove(path)
		if err != nil {
			return fmt.Errorf("Error borrando %s\n", file.Name())
		}
	}
	fmt.Printf("Borrado (%d) archivo(s)", len(files))
	return nil
}

func FilterFilesByMatch(files []os.DirEntry, match string) []os.DirEntry {
	var filesFilter []os.DirEntry
	count := 1
	for _, entry := range files {
		if strings.Contains(entry.Name(), match) {
			filesFilter = append(filesFilter, entry)
			fmt.Printf("%d: %s\n", count, entry.Name())
			count += 1
		}
	}
	return filesFilter
}
