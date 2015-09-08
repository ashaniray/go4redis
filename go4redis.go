package go4redis

import (
	"fmt"
	"log"
)

func main() {

	c, err := Dial("localhost:6379")

	if err != nil {
		log.Fatal(err)
	}

	n, _ := c.lpush("foo", "1", "2", "3", "4", "5")

	fmt.Println("inserted", n)

	l, err := c.llen("foo")
	fmt.Println(l)

}
