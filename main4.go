package main

import (
	"context"
	"fmt"
	flood "task/floodtest"
	"time"
)

type FloodControl interface {
	Check(ctx context.Context, userID int64) (bool, error)
}

func main() {
	fc := flood.NewFloodCounter()

	for i := 0; i < 8; i++ {
		result, err := fc.Check(context.Background(), 5)
		if err != nil {
			fmt.Println("Ошибка при выполнении проверки флуд-контроля:", err)
			return
		}
		if result {
			fmt.Println("Флуд-контроль пройден")
		} else {
			fmt.Println("Достигнут лимит флуд-контроля")
		}
	}

	time.Sleep(11 * time.Second)

	for i := 0; i < 8; i++ {
		result, err := fc.Check(context.Background(), 1)
		if err != nil {
			fmt.Println("Ошибка при выполнении проверки флуд-контроля:", err)
			return
		}
		if result {
			fmt.Println("Флуд-контроль пройден")
		} else {
			fmt.Println("Достигнут лимит флуд-контроля")
		}
	}
}
