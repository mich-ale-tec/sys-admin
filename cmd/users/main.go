package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jedib0t/go-pretty/table"
)

type TUser struct {
	User  string
	Uid   int
	Shell string
}

func main() {
	file, err := os.Open("/etc/passwd")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var users = []TUser{}
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ":")
		uid, err := strconv.Atoi(fields[2])
		if err != nil {
			fmt.Println("Error al leer el archivo: ", err)
			return
		}
		if uid < 1000 {
			continue
		}
		users = append(users, TUser{
			User:  fields[0],
			Uid:   uid,
			Shell: fields[6],
		})
	}
	/*
	 * Creación de la tabla
	 */
	t := table.NewWriter()

	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{
		"Usuario",
		"UID",
		"Shell",
	})
	for _, user := range users {
		t.AppendRow(table.Row{
			user.User,
			user.Uid,
			user.Shell,
		})
	}
	fmt.Println("Usuarios")
	t.Render()

	if err := scanner.Err(); err != nil {
		fmt.Println("Error leyendo archivo: ", err)
	}
}
