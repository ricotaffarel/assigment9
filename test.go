package main

// func main() {
// 	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
// 	if err != nil {
// 		log.Fatalln(err)
// 		return
// 	}

// 	fmt.Println(res.Body)
// 	body, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		log.Fatalln(err)
// 		return
// 	}

// 	defer res.Body.Close()
// 	ab := string(body)
// 	log.Println(ab)
// }

// func main() {
// 	data := map[string]interface{}{
// 		"title":  "dono",
// 		"body":   "jordan",
// 		"userId": 10,
// 	}

// 	reqJson, err := json.Marshal(data)
// 	Client := http.Client{}
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	fmt.Println(string(reqJson))

// 	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(reqJson))
// 	req.Header.Set("Content-Type", "application/json")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	res, err := Client.Do(req)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	defer res.Body.Close()

// 	body, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		log.Fatalln(err)
// 		// return
// 	}

// 	log.Println(string(body))
// }
