package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

// post will post all payloads from a file to the url specified
func post(filename string, url string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		payload := scanner.Text()
		strings.TrimSpace(payload)
		client := &http.Client{}
		url := os.Args[2]
		body := bytes.NewBufferString(payload)
		fmt.Println(payload)
		req, err := http.NewRequest("POST", url, body)
		if err != nil {
			fmt.Println("error:", err)
		}
		_, err2 := client.Do(req)
		if err2 != nil {
			fmt.Println("error:", err2)
		}
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: seeder filename url")
	} else {
		post(os.Args[1], os.Args[2])
	}
}
