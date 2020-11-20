package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type User struct {
	ChatId   int    `json:"id"`
	Username string `json:"username"`
	PhoneNo  string `json:"phone"`
}

func main() {

	file, err := os.Open("./leak.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	line := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line++
		raw := fmt.Sprintf("%s", scanner.Text())
		data := []byte(raw)
		var user User
		err := json.Unmarshal(data, &user)
		if err != nil {
			fmt.Println("err",line)
			panic(err)
			continue;
		}
		if user.Username != "" {
			sum := 0
			for _, v := range user.Username {
				sum += int(v)
			}
			println(sum % 10)
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
