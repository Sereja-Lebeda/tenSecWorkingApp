package main

import (
	"fmt"
	"os"
	"time"
)

func StartWork() {
	for i := 0; i < 10; i++ {
		fmt.Println(i, " - прошло времени")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	fmt.Println("Программа завершится с ошибкой 1")
	StartWork()
	os.Exit(1) // Устанавливаем код завершения 1
}
