package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)
//
func StartWork(duration int) {
	for i := 0; i < duration; i++ {
		fmt.Println(i+1, "- прошло времени")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Использование: program <количество_секунд>")
		os.Exit(1)
	}

	duration, err := strconv.Atoi(os.Args[1])
	if err != nil || duration <= 0 {
		fmt.Println("Ошибка: укажите положительное число секунд")
		os.Exit(1)
	}

	StartWork(duration)
	fmt.Println("Программа завершится с ошибкой 1")
	os.Exit(1) // Устанавливаем код завершения 1
}
