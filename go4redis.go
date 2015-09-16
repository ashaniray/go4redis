package go4redis

import (
	"log"
)


func main() {

	_, err := Dial("localhost:6379")
	if err != nil {
		log.Fatal(err)
	}

}
