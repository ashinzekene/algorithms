package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	date := "2019-04-18T13:37:00Z"
	now := time.Now()
	later, _ := time.Parse(time.RFC3339, date)

	p(now)
	p(later)
	fmt.Printf("%v", later.After(now))
}
