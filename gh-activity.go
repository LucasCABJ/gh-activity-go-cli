package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	username := os.Args[1]
	fetchUri := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	resp, err := http.Get(fetchUri)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(resp.Body)
}
