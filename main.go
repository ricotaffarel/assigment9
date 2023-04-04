package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	var statusWater string
	var statusWind string
	for {
		valueWater := rand.Intn(100) + 1
		valueWind := rand.Intn(100) + 1
		if valueWater <= 5 && valueWind <= 6 {
			statusWater = "aman"
			statusWind = "aman"
			Post(valueWater, valueWind, statusWater, statusWind)
		} else if valueWater >= 6 && valueWater <= 8 && valueWind >= 7 && valueWind <= 15 {
			statusWater = "siaga"
			statusWind = "siaga"
			Post(valueWater, valueWind, statusWater, statusWind)
		} else if valueWater >= 8 && valueWind >= 15 {
			statusWater = "bahaya"
			statusWind = "bahaya"
			Post(valueWater, valueWind, statusWater, statusWind)
		}
		time.Sleep(15 * time.Second)
	}

}

func Post(water int, wind int, status_water string, status_wind string) {
	data := map[string]interface{}{
		"water": water,
		"wind":  wind,
	}

	reqJson, err := json.Marshal(data)
	Client := http.Client{}
	if err != nil {
		log.Fatalln(err)
		return
	}

	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(reqJson))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatalln(err)
		return
	}

	res, err := Client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Println(string(body))
	fmt.Println("status water : ", status_water)
	fmt.Println("status wind : ", status_wind)
	fmt.Println("===============================")
}
