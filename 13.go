package main

import (
	"fmt"
	"sync"
)

func square(num int, wg *sync.WaitGroup, resultChan chan int) {
	defer wg.Done()
	resultChan <- num * num
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	resultChan := make(chan int, len(numbers))

	for _, num := range numbers {
		wg.Add(1)
		go square(num, &wg, resultChan)
	}

	// Закрываем канал после завершения всех горутин
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Суммируем результаты из канала
	sum := 0
	for res := range resultChan {
		sum += res
	}

	fmt.Printf("Сумма квадратов: %d\n", sum)
}
