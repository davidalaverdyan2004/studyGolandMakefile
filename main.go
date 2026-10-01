package main

import (
	"context"
	"fmt"
	"studygoland/bd"
)

func main() {
	ctx := context.Background()

	conn := bd.ConnectBd(ctx)
	defer conn.Close(ctx)

	err := bd.CreateTable(ctx, conn)
	if err != nil {
		fmt.Println(err)
		return
	}
}
