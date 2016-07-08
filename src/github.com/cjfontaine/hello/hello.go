package main

import (
    "fmt"
    "sync"
    "time"
)

func main(){
	cities := []string{"Orlando", "Miami", "Tampa", "Jacksonville"}

	var wg sync.WaitGroup

	wg.Add(len(cities))

	for _, city := range cities {
		go processCity(city, 1000, &wg)
	}

	wg.Wait()
}

func processCity(city string, milliseconds time.Duration, wg *sync.WaitGroup){
	defer wg.Done()

    duration := milliseconds * time.Millisecond
    time.Sleep(duration)    

    t := time.Now()
    
    fmt.Println("Function in background, time:", t.Format(time.RFC3339), " city: ", city)
}