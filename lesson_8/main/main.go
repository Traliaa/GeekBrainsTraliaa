package main

import (
	"fmt"

	"GeekBrains/lesson_8/main/config"
)

func main() {

	config, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", config)

}
