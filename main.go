package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	db "studygoland/Db"
)

func main() {
	ctx := context.Background()

	conn := db.ConnectBd(ctx)
	defer conn.Close(ctx)

	err := db.CreateTable(ctx, conn)
	if err != nil {
		fmt.Println(err)
		return
	}

	newUser := os.Getenv("NEW_USER")

	switch newUser {
	case "YES":
		fmt.Println("add user")
		fmt.Println("введите имя:")
		reader := bufio.NewReader(os.Stdin)
		name, err := reader.ReadString('\n')
		name = strings.TrimSpace(name)
		if err != nil {
			fmt.Println("err")
			return
		}

		var phonePtr *string // по умолчанию nil
		fmt.Println("введите nomer:")
		phone, err := reader.ReadString('\n')

		phone = strings.TrimSpace(phone) // ← вот эта строка

		if phone != "" {
			phonePtr = &phone // указываем на строку, только если она не пустая
		}

		if err != nil {
			fmt.Println("err")
			return
		}

		err = db.AddUser(ctx, conn, name, phonePtr)
		if err != nil {
			return
		}
		fmt.Println("Пользователь добавлен")

	case "NO":
		fmt.Println("all table")
		tables, err := db.GetUsers(ctx, conn)
		if err != nil {
			fmt.Println("err")
			return
		}
		for _, table := range tables {
			if table.PhoneNumber != nil {
				fmt.Printf("%d. %s, телефон: %s\n", table.ID, table.FullName, *table.PhoneNumber)

			} else {
				fmt.Printf("%d. %s, нет телефона\n", table.ID, table.FullName)
			}
		}

	default:
		fmt.Println("err, YES or NO")
		os.Exit(1)
	}

}
