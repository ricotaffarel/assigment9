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

type Data struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	var statusWater string
	var statusWind string
	for {
		valueWater := rand.Intn(100) + 1
		valueWind := rand.Intn(100) + 1

		if valueWater < 5 {
			statusWater = "Aman"
		} else if valueWater >= 5 && valueWater <= 8 {
			statusWater = "Siaga"
		} else {
			statusWater = "Bahaya"
		}

		if valueWind < 6 {
			statusWind = "Aman"
		} else if valueWind >= 6 && valueWind <= 15 {
			statusWind = "Siaga"
		} else {
			statusWind = "Bahaya"
		}

		Post(valueWater, valueWind, statusWater, statusWind)
		time.Sleep(15 * time.Second)
	}

}

func Post(water int, wind int, status_water string, status_wind string) {
	var data Data
	data.Water = water
	data.Wind = wind

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

	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		log.Fatalln(err)
		return
	}

	toJson, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Println(string(toJson))
	fmt.Println("status water : ", status_water)
	fmt.Println("status wind : ", status_wind)
	fmt.Println("===============================")
}
