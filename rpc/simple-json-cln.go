package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Connecting...")
	req, _ := http.NewRequest("POST", "http://localhost:7000",
		strings.NewReader(`{"code": 5,"message": "Anybody there?"}`))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Printf("wrote request")
}
