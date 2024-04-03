package main

import (
	"fmt"

	"github.com/hardik-bhasin07/puppy"
)

func main() {
	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())
	fmt.Println(puppy.GrownBark())
	fmt.Println(puppy.GrownBarks())
	puppy.Version11()
}
