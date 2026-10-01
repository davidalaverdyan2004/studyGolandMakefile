package main

import (
	"context"
	"fmt"
	"os"
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
