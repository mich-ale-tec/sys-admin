package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Buscando la carpeta del usuario...")
	user, err := GetNameDirUser()
	if err != nil {
		fmt.Println("No se pudo encontrar el usuario: ", err)
		return
	}

	fmt.Println("Buscando fonts...")
	pathDir := BuildPathDir(user)

	dirEntries, err := GetItemsDir(pathDir)
	if err != nil {
		fmt.Println("No se pudo leer el directorio: ", err)
		return
	}

	files := FilterFiles(dirEntries)

	fmt.Print("filter (f), remove (r) :")
	scanner.Scan()
	command := scanner.Text()

	err = ExecCommand(command, files, pathDir)
	if err != nil {
		fmt.Println("No se pudo ejecutar el comando: ", err)
	}
}

func ExecCommand(command string, files []os.DirEntry, pathDir string) error {
	switch command {
	case "f":
		err := CommandFilter(files, pathDir)
		if err != nil {
			return err
		}
	default:
		fmt.Println("Opción inválida")
	}
	return nil
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

func GetItemsDir(pathDir string) ([]os.DirEntry, error) {
	dirEntries, err := os.ReadDir(pathDir)
	if err != nil {
		return nil, err
	}
	return dirEntries, nil
}

func BuildPathDir(user string) string {
	pathsArray := []string{`C:\`, "Users", user, "AppData", "Local", "Microsoft", "Windows", "Fonts"}
	return filepath.Join(pathsArray...)
}

func GetNameDirUser() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}
	return user.Username, nil
}
