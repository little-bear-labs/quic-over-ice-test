package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func calculateSum(number int) int {
	sum := 0
	for i := 1; i <= number; i++ {
		sum += i
	}
	return sum
}
func hello(w http.ResponseWriter, req *http.Request) {
	myURL := "https://localhost:8090/?n=10"
	parsedURL, err := url.Parse(myURL)
	if err != nil {
		fmt.Println("Can't parse URL", err)
		return
	}
	parts := strings.Split(parsedURL.RawQuery, "=")
	n, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Println("number is not found", err)
		return
	}
	var res int = calculateSum(n)
	fmt.Fprintf(w, "The sum from 1 to %d is: %d", n, res)
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8090", nil)
}
