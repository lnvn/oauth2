package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("Protocol: %s\n", resp.Proto)
	fmt.Print("Header:")
	for key, values := range resp.Header {
		fmt.Printf("%s: %s\n", key, strings.Join(values, ", "))
	}
	defer resp.Body.Close()
}