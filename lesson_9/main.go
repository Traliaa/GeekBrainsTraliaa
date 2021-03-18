package main

import (
	"fmt"

	"GeekBrains/lesson_9/config"
)

func main() {

	configEnv, err := config.NewConfigForENV()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", configEnv)

	configYaml, err := config.NewConfigForYaml("config.yaml")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", configYaml)

}
