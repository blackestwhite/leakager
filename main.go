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
			fmt.Println("err", line)
			fmt.Println(err)
			continue
		}

		where := fmt.Sprintf("./chat_id/%v.txt", user.ChatId%100)
		f, err := os.OpenFile(where,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		toWrite := fmt.Sprintf("%s\n", raw)
		if _, err := f.WriteString(toWrite); err != nil {
			log.Println(err)
		}
		f.Close()

		if user.Username != "" {
			sum := 0
			for _, v := range user.Username {
				sum += int(v)
			}
			where := fmt.Sprintf("./username/%v.txt", sum%100)
			f, err := os.OpenFile(where,
				os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Println(err)
			}
			toWrite := fmt.Sprintf("%s\n", raw)
			if _, err := f.WriteString(toWrite); err != nil {
				log.Println(err)
			}
			f.Close()
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
