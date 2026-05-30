package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	pathDir, err := BootDirUser()
	if err != nil {
		fmt.Println(err)
		return
	}

	files, err := BootFiles(pathDir)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = BootCommand(files, pathDir)
	if err != nil {
		fmt.Println(err)
	}

}

func BootCommand(files []os.DirEntry, path string) error {
	command := ReadCommand()
	err := ExecCommand(command, files, path)
	if err != nil {
		return fmt.Errorf("No se pudo ejecutar el comando: %w", err)
	}
	return nil
}

func BootFiles(pathDir string) ([]os.DirEntry, error) {
	dirEntries, err := GetItemsDir(pathDir)
	if err != nil {
		return nil, fmt.Errorf("No se pudo leer el directorio: %w", err)
	}

	files := FilterFiles(dirEntries)
	return files, nil
}

func ReadCommand() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("filter (f), remove (r) :")
	scanner.Scan()
	command := scanner.Text()
	return command
}

func BootDirUser() (string, error) {
	fmt.Println("Buscando la carpeta del usuario...")
	user, err := GetNameDirUser()
	if err != nil {
		return "", fmt.Errorf("No se pudo obtener el nombre del usuario: %w", err)
	}
	pathDir := BuildPathDir(user)
	return pathDir, nil
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

func BuildPathDir(userDir string) string {
	pathsArray := []string{userDir, "AppData", "Local", "Microsoft", "Windows", "Fonts"}
	return filepath.Join(pathsArray...)
}

func GetNameDirUser() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}
	return u.HomeDir, nil
}
