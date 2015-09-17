package main

import (
  "fmt"
  "go4redis"
)

func main() {
	c, _ := go4redis.Dial("localhost:6379")
  fmt.Println(c)
}
