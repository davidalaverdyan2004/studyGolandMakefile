package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	count := os.Getenv("OUTPUT_COUNT")
	member, err := strconv.Atoi(count)
	if err != nil {
		fmt.Println("Ошибка преобразования OUTPUT_COUNT в целое число")
		return
	}

	for i := 1; i <= member; i++ {
		fmt.Println("Работа цикла:", i)
	}
}
