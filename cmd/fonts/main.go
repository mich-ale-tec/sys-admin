package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Buscando la carpeta del usuario...")
	user := GetNameDirUser()

	fmt.Println("Buscando fonts...")
	pathDir := BuildPathDir(user)

	dirEntries := GetItemsDir(pathDir)

	files := FilterFiles(dirEntries)

	fmt.Print("filter (f), remove (r) :")
	scanner.Scan()
	command := scanner.Text()

	ExecCommand(command, files, pathDir)

}

func ExecCommand(command string, files []os.DirEntry, pathDir string) {
	switch command {
	case "f":
		CommandFilter(files, pathDir)
	default:
		fmt.Println("Opción inválida")
	}
}

func FilterFiles(itemsDir []os.DirEntry) []os.DirEntry {
	var files []os.DirEntry
	for i, file := range itemsDir {
		if file.IsDir() {
			continue
		}
		files = append(files, file)
		fmt.Printf("%d. %s\n", i+1, file.Name())
	}
	return files
}

func GetItemsDir(pathDir string) []os.DirEntry {
	dirEntries, err := os.ReadDir(pathDir)
	if err != nil {
		fmt.Println("No se pudo leer el directorio")
		panic(err)
	}
	return dirEntries
}

func BuildPathDir(user string) string {
	pathsArray := []string{`C:\`, "Users", user, "AppData", "Local", "Microsoft", "Windows", "Fonts"}
	return filepath.Join(pathsArray...)
}

func GetNameDirUser() string {
	user := os.Getenv("USER")
	if user == "" {
		user = "72720804"
	}
	return user
}
