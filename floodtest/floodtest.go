package flood

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type FloodCounter struct {
	requests map[int64][]time.Time
}

func NewFloodCounter() *FloodCounter {
	return &FloodCounter{
		requests: make(map[int64][]time.Time),
	}
}

type conf struct {
	Second int `json:"seconds"`
	Req    int `json:"trys"`
}

func Config() conf {
	confData, err := os.Open("floodtest/config.json")
	if err != nil {
		fmt.Println("Ошибка открытия конфига")
	}

	defer confData.Close()

	data, err := ioutil.ReadAll(confData)
	if err != nil {
		fmt.Println("Ошибка чтения конфига")
	}

	var config conf
	jsonErr := json.Unmarshal(data, &config)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return config

}

func (fc *FloodCounter) Check(ctx context.Context, userID int64) (bool, error) { // Check реализует метод FloodControl.Check для проверки флуд-контроля.

	result := Config()

	duration := time.Duration(result.Second) * time.Second // вызовы за последние 10 секунд

	fc.cleanup()

	requests := fc.requests[userID] // список временных меток для пользователя

	count := 0
	for _, t := range requests { // фильтр временных метки за 10 секунд
		if time.Since(t) <= duration {
			count++
		}
	}

	if count > result.Req { // проверка на флуд
		return false, nil
	}

	fc.requests[userID] = append(requests, time.Now()) // обновляем время последнего запроса

	return true, nil // флуд контроль пройден
}

func (fc *FloodCounter) cleanup() { //очистка буфера запросов по истечению 10 секунд
	duration := 10 * time.Second
	currentTime := time.Now()

	for userID, requests := range fc.requests {
		var recentRequests []time.Time
		for _, t := range requests {
			if currentTime.Sub(t) <= duration {
				recentRequests = append(recentRequests, t)
			}
		}
		fc.requests[userID] = recentRequests
	}
}

/*func main() {
	fc := NewFloodCounter()

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
*/
