package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"charm.land/huh/v2"
	"charm.land/huh/v2/spinner"
	"charm.land/lipgloss/v2"
)

var (
	nroRepository int
)

func main() {
	//pathParts := []string{`C:\`, "Users", "72720804", "source", "repos"}
	pathParts := []string{`C:\`, "Users", "Usuario", "source", "repos"}
	entries := readRepositories(pathParts)

	showTitle()
	showRepositories(entries)

	pathRepositoryParts := append(pathParts, entries[nroRepository].Name())
	pathRepository := filepath.Join(pathRepositoryParts...)
	execCommand(pathRepository)
}

func readRepositories(pathParts []string) []os.DirEntry {
	var entries []os.DirEntry
	err := spinner.New().
		Title("Leyendo repositorios...").
		Action(func() {
			time.Sleep(1 * time.Second)
			path := filepath.Join(pathParts...)
			var err error
			entries, err = os.ReadDir(path)
			if err != nil {
				panic(err)
			}
		}).
		Run()

	if err != nil {
		panic(err)
	}
	return entries
}

func showTitle() {
	title := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("205")).
		Border(lipgloss.RoundedBorder()).
		Padding(0, 3)
	fmt.Println(title.Render("🐱 LG - Repositorios"))
	fmt.Println("")
}

func showRepositories(entries []os.DirEntry) {
	options := make([]huh.Option[int], 0, len(entries))
	for i, entry := range entries {
		if entry.IsDir() {
			options = append(
				options,
				huh.NewOption(entry.Name(), i),
			)
		}
	}
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int]().
				Title("Escoge tu repositorio").
				Options(options...).
				Value(&nroRepository),
		))

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func execCommand(path string) {
	cmd := exec.Command("lazygit")
	cmd.Dir = path
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
