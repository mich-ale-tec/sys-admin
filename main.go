package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	path_parts := []string{`C:\`, "Users", "72720804", "source", "repos"}
	path := filepath.Join(path_parts...)
	entries, err := os.ReadDir(path)

	if err != nil {
		panic(err)
	}

	fmt.Println("Repositorios:")
	fmt.Println("")
	for index, entry := range entries {
		fmt.Printf("[%d] %s\n", index, entry.Name())
	}
	fmt.Println("")

	var nro_repositorio int
	fmt.Print("Repositorio: ")
	fmt.Scanln(&nro_repositorio)

	repositorio := entries[nro_repositorio].Name()
	fmt.Printf("Elegiste: [%s] ¿Desea continuar? (S/n): ", repositorio)

	var res_continuar string
	fmt.Scanln(&res_continuar)

	if strings.ToLower(strings.TrimSpace(res_continuar)) == "n" {
		os.Exit(0)
	}

	path_repository_parts := append(path_parts, entries[nro_repositorio].Name())
	path_repository := filepath.Join(path_repository_parts...)

	cmd := exec.Command("lazygit")
	cmd.Dir = path_repository
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
